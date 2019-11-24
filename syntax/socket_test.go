package syntax

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"testing"
)

// socket网络通信

// Http Get请求
func TestHttpGet(t *testing.T) {
	resp, err := http.Get("https://www.baidu.com")
	if err == nil {
		fmt.Printf("类型 = %T\n", resp)
		fmt.Printf("resp = %v", *resp)
	}
}

// Http服务器
func TestHttpServer(t *testing.T) {
	http.Handle("/", http.FileServer(http.Dir(".")))
	_ = http.ListenAndServe(":8080", nil)
}

// TCP 服务端
func Test_TcpServer(t *testing.T) {
	service := "127.0.0.1:8888"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		handleClient(conn)
		_ = conn.Close()
	}
}

func handleClient(conn net.Conn) {
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}
		rAddr := conn.RemoteAddr()
		fmt.Println("Receive from client", rAddr.String(), string(buf[0:n]))
		_, err2 := conn.Write([]byte("Welcome client"))
		if err2 != nil {
			return
		}
	}
}

func checkError(err error) {
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error())
		os.Exit(1)
	}
}

// TCP 客户端
func TestTcpClient(t *testing.T) {
	var buf [512]byte
	service := "127.0.0.1:8888"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	rAddr := conn.RemoteAddr()
	n, err := conn.Write([]byte("Hello server"))
	checkError(err)
	n, err = conn.Read(buf[0:])
	checkError(err)
	fmt.Println("Reply form server", rAddr.String(), string(buf[0:n]))
	_ = conn.Close()
	os.Exit(0)
}
