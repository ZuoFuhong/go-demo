package http_client

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"net/url"
	"strings"
)

// runProxyServer 代理服务器
func runProxyServer() {
	l, err := net.Listen("tcp", "127.0.0.1:9999")
	if err != nil {
		log.Panic(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Panic(err)
		}
		go handleClientRequest(conn)
	}
}

func handleClientRequest(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 0, 4096)
	n, err := conn.Read(buf[len(buf):cap(buf)])
	if err != nil {
		log.Panic(err)
	}
	var method, host, address string
	fmt.Sscanf(string(buf[:bytes.IndexByte(buf[0:n], '\n')]), "%s%s", &method, &host)
	hostPortUrl, err := url.Parse(host)
	if err != nil {
		log.Panic(err)
	}
	if hostPortUrl.Opaque == "443" {
		address = hostPortUrl.Scheme + ":443"
	} else {
		if strings.Index(hostPortUrl.Host, ":") == -1 {
			address = hostPortUrl.Host + ":80"
		} else {
			address = hostPortUrl.Host
		}
	}
	log.Printf("new request method = %s, host = %s, address = %s\n", method, host, address)
	// 开始拨号
	server, err := net.Dial("tcp", address)
	if err != nil {
		log.Panic(err)
	}
	if method == "CONNECT" {
		fmt.Fprint(conn, "HTTP/1.1 200 Connection established\\r\\n\\r\\n")
	} else {
		server.Write(buf[:n])
	}
	// 进行转发
	io.Copy(conn, server)
	io.Copy(server, conn)
}
