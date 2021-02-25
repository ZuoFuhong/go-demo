package http_protocol

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/url"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// 基于TCP实现http1.1协议
// 1.直接读取流字节
// 2.使用缓冲IO读取流字节
// 3.使用sync.Pool复用大对象
// 4.处理异常
var errTooLarge = errors.New("http: request too large")

type Server struct {
	addr string
}

// conn 网络连接
type conn struct {
	server     *Server
	cancelCtx  context.CancelFunc
	rwc        net.Conn
	remoteAddr string
	bufr       *bufio.Reader
	bufw       *bufio.Writer
}

type Request struct {
	Method     string
	RequestURI string
	Proto      string // "HTTP/1.0"
	URL        *url.URL
}

func NewServer(addr string) Server {
	return Server{
		addr: addr,
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Printf("net.Listen err:%v", err)
		return err
	}
	defer ln.Close()
	for {
		rwc, err := ln.Accept()
		if err != nil {
			log.Printf("Accept conn err:%v", err)
			return err
		}
		conn := s.newConn(rwc)
		go conn.serve()
	}
}

func (s *Server) newConn(rwc net.Conn) *conn {
	c := &conn{
		rwc:    rwc,
		server: s,
	}
	return c
}

func (c *conn) serve() {
	c.remoteAddr = c.rwc.RemoteAddr().String()
	defer func() {
		if err := recover(); err != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			log.Printf("http: panic serving %v: %v\n%s", c.remoteAddr, err, buf)
		}
		c.close()
	}()
	c.bufr = bufio.NewReader(c.rwc)
	c.bufw = bufio.NewWriter(c.rwc)

	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	for {
		err := c.handleRequest(ctx)
		if err != nil {
			const errorHeaders = "\r\nContent-Type: text/plain; charset=utf-8\r\nConnection: close\r\n\r\n"
			switch {
			case err == errTooLarge:
				const publicErr = "431 Request Header Fields Too Large"
				fmt.Fprintf(c.rwc, "HTTP/1.1 "+publicErr+errorHeaders+publicErr)
				c.closeWriteAndWait()
			default:
				publicErr := "400 Bad Request"
				fmt.Fprintf(c.rwc, "HTTP/1.1 "+publicErr+errorHeaders+publicErr)
				return
			}
		}
		c.finalFlush()
	}
}

func (c *conn) handleRequest(ctx context.Context) error {
	// 1.读取request line
	// 忽略一行打满缓存区
	reqLine, _, err := c.bufr.ReadLine()
	if err != nil {
		return err
	}
	req := new(Request)
	var ok bool
	req.Method, req.RequestURI, req.Proto, ok = parseRequestLine(string(reqLine))
	if !ok {
		return badStringError("malformed HTTP request", string(reqLine))
	}
	rawurl := req.RequestURI
	if req.URL, err = url.ParseRequestURI(rawurl); err != nil {
		return err
	}
	log.Printf("request Line: %s\n", string(reqLine))
	// 2.读取request header
	headers := make(map[string][]string, 0)
	for {
		kv, _, err := c.bufr.ReadLine()
		if err != nil {
			return err
		}
		if len(kv) == 0 { // blank line - no continuation
			break
		}
		idx := bytes.IndexByte(kv, ':')
		key := string(kv[:idx])
		idx++
		if idx < len(kv) && (kv[idx] == ' ' || kv[idx] == '\t') {
			idx++
		}
		value := string(kv[idx:])
		if val, ok := headers[key]; ok {
			headers[key] = append(val, value)
		} else {
			headers[key] = []string{value}
		}
	}
	log.Printf("request headers: %v\n", headers)
	// 3.读取request Body
	if req.Method == "POST" {
		// 仅考虑固定长度的body
		cLen, _ := strconv.ParseInt(headers["Content-Length"][0], 10, 32)
		b := make([]byte, 0, 512)
		for {
			if len(b) == cap(b) {
				b = append(b, 0)[:len(b)]
			}
			n, err := c.bufr.Read(b[len(b):cap(b)])
			b = b[:len(b)+n]
			if int64(len(b)) >= cLen {
				break
			}
			if err != nil {
				if err == io.EOF {
					break
				}
			}
		}
		log.Printf("request contentLen:%d, bodyLen: %d, body: %s\n", cLen, len(b), string(b))
	}
	// 4.响应请求
	const rspHeaders = "\r\nContent-Type: application/json; charset=utf-8\r\nConnection: close\r\n\r\n"
	fmt.Fprintf(c.rwc, "HTTP/1.1 200 OK"+rspHeaders+`{"retcode":0}`)
	return nil
}

// parseRequestLine parses "GET /foo HTTP/1.1" into its three parts.
func parseRequestLine(line string) (method, requestURI, proto string, ok bool) {
	s1 := strings.Index(line, " ")
	s2 := strings.Index(line[s1+1:], " ")
	if s1 < 0 || s2 < 0 {
		return
	}
	s2 += s1 + 1
	return line[:s1], line[s1+1 : s2], line[s2+1:], true
}

func (c *conn) close() {
	_ = c.rwc.Close()
}

// closeWriteAndWait 刷新所有未完成的数据并发送FIN数据包。然后，稍等片刻，希望客户端在后续的RST之前对其进行处理。
func (c *conn) closeWriteAndWait() {
	c.finalFlush()
	if tcp, ok := c.rwc.(*net.TCPConn); ok {
		tcp.CloseWrite()
	}
	time.Sleep(500 * time.Millisecond)
}

func (c *conn) finalFlush() {
	_ = c.rwc.Close()
}

func badStringError(what, val string) error {
	return fmt.Errorf("%s %q", what, val)
}
