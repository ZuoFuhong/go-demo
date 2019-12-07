package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// 简单的负载均衡

// 后端服务的信息
type Backend struct {
	URL          *url.URL
	Alive        bool
	mux          sync.RWMutex
	ReverseProxy *httputil.ReverseProxy
}

// 设置服务的状态
func (b *Backend) setAlive(alive bool) {
	b.mux.Lock()
	b.Alive = alive
	b.mux.Unlock()
}

// 判断服务状态
func (b *Backend) isAlive() bool {
	var alive bool
	b.mux.RLock()
	alive = b.Alive
	b.mux.RUnlock()
	return alive
}

// 后端服务映射池
type ServerPool struct {
	backends []*Backend
	current  uint64
}

// 增加后端服务到映射池
func (s *ServerPool) addBackend(backend *Backend) {
	s.backends = append(s.backends, backend)
}

// 自动增加计数器，取模计算索引值
func (s *ServerPool) nextIndex() int {
	index := atomic.AddUint64(&s.current, uint64(1)) % uint64(len(s.backends))
	return int(index)
}

// 标记后端服务状态
func (s *ServerPool) markBackendStatus(backend *url.URL, alive bool) {
	for _, b := range s.backends {
		// 地址比较
		if b.URL.String() == backend.String() {
			b.setAlive(alive)
			return
		}
	}
}

// 获取后端服务的连接（轮询调度算法Round-Robin）
func (s *ServerPool) getNextPeer() *Backend {
	next := s.nextIndex()
	l := len(s.backends) + next
	for i := next; i < l; i++ {
		idx := i % len(s.backends)
		if s.backends[idx].isAlive() {
			if i != next {
				atomic.StoreUint64(&s.current, uint64(idx))
			}
			return s.backends[idx]
		}
	}
	return nil
}

// 检查后端服务是否存活
func isBackendAlive(u *url.URL) bool {
	timeout := 2 * time.Second
	conn, e := net.DialTimeout("tcp", u.Host, timeout)
	if e != nil {
		log.Println("service unreachble, error", e)
		return false
	}
	_ = conn.Close()
	return true
}

// 后端服务状态检查及状态更新
func (s *ServerPool) HealthCheck() {
	for _, b := range s.backends {
		status := "up"
		alive := isBackendAlive(b.URL)
		b.setAlive(alive)
		if !alive {
			status = "down"
		}
		log.Printf("%s [%s]\n", b.URL, status)
	}
}

const (
	Attempts = iota
	Retry
)

// 提取请求上下文参数
func getAttemptsFromContext(r *http.Request) int {
	if attempts, ok := r.Context().Value(Attempts).(int); ok {
		return attempts
	}
	return 1
}

func getRetryFromContext(r *http.Request) int {
	if retry, ok := r.Context().Value(Retry).(int); ok {
		return retry
	}
	return 0
}

// 负载均衡
func lb(w http.ResponseWriter, r *http.Request) {
	attempts := getAttemptsFromContext(r)
	if attempts > 3 {
		log.Printf("%s(%s) Max attempts reached, terminating\n", r.RemoteAddr, r.URL.Path)
		http.Error(w, "Service not available", http.StatusServiceUnavailable)
		return
	}
	// 反向代理
	peer := serverPool.getNextPeer()
	if peer != nil {
		peer.ReverseProxy.ServeHTTP(w, r)
		return
	}
	http.Error(w, "Service not available", http.StatusServiceUnavailable)
}

// 每10s执行一次服务检查
func healthCheck() {
	ticker := time.NewTicker(time.Second * 10)
	for {
		select {
		case <-ticker.C:
			log.Println("Starting health check...")
			serverPool.HealthCheck()
			log.Println("Health check completed")
		}
	}
}

var serverPool ServerPool

// 负载均衡
// 启动：./simple_lb -backends=http://127.0.0.1:3031,http://127.0.0.1:3032,http://127.0.0.1:3033,http://127.0.0.1:3034
func main() {
	// 解析命令行参数
	var serverList string
	var port int
	flag.StringVar(&serverList, "backends", "", "Load balanced backends, use commas to separate")
	flag.IntVar(&port, "port", 3030, "Port to serve")
	flag.Parse()

	if len(serverList) == 0 {
		log.Fatal("Please provide one or more backends to load balance")
	}

	// 后端服务注册
	tokens := strings.Split(serverList, ",")
	for _, tok := range tokens {
		serverUrl, e := url.Parse(tok)
		if e != nil {
			log.Fatal(e)
		}
		proxy := httputil.NewSingleHostReverseProxy(serverUrl)
		proxy.ErrorHandler = func(writer http.ResponseWriter, request *http.Request, e error) {
			log.Printf("[%s] %s\n", serverUrl.Host, e.Error())
			retries := getRetryFromContext(request)
			if retries < 3 {
				select {
				case <-time.After(10 * time.Millisecond):
					ctx := context.WithValue(request.Context(), Retry, retries+1)
					proxy.ServeHTTP(writer, request.WithContext(ctx))
				}
				return
			}
			serverPool.markBackendStatus(serverUrl, false)
			attempts := getAttemptsFromContext(request)
			log.Printf("%s(%s) Attempting retry %d\n", request.RemoteAddr, request.URL.Path, attempts)
			ctx := context.WithValue(request.Context(), Attempts, attempts+1)
			lb(writer, request.WithContext(ctx))
		}

		// 添加后端服务
		serverPool.addBackend(&Backend{URL: serverUrl, Alive: true, ReverseProxy: proxy})
		log.Printf("Configured server: %s\n", serverUrl)
	}

	// 启动负载均衡服务
	server := http.Server{
		Addr:    fmt.Sprintf("127.0.0.1:%d", port),
		Handler: http.HandlerFunc(lb),
	}

	// 监控检查
	go healthCheck()

	log.Printf("Load Balancer started at :%d\n", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
