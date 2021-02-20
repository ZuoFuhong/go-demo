package tls

import (
	"bufio"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
)

var certDir = "%s/cert/"

func init() {
	path, _ := os.Getwd()
	certDir = fmt.Sprintf(certDir, path)
}

// Go Package tls部分实现了 tls 1.2的功能，可以满足我们日常的应用。
// Package crypto/x509提供了证书管理的相关操作。
func RunServer() {
	cert, err := tls.LoadX509KeyPair(certDir+"server.pem", certDir+"server.key")
	if err != nil {
		log.Panic(err)
	}
	config := &tls.Config{Certificates: []tls.Certificate{cert}}
	ln, err := tls.Listen("tcp", ":443", config)
	if err != nil {
		log.Panic(err)
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Panic(err)
		}
		handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	for {
		msg, err := r.ReadString('\n')
		if err != nil {
			log.Println("ReadString err:", err)
			return
		}
		fmt.Println(msg)
		n, err := conn.Write([]byte("world\n"))
		if err != nil {
			log.Println(n, err)
			return
		}
	}
}

func RunTlsServer() {
	cert, err := tls.LoadX509KeyPair(certDir+"server.pem", certDir+"server.key")
	if err != nil {
		log.Panic(err)
	}
	certBytes, err := ioutil.ReadFile(certDir + "client.pem")
	if err != nil {
		log.Panic(err)
	}
	// 定义一组证书颁发机构
	clientCertPool := x509.NewCertPool()
	ok := clientCertPool.AppendCertsFromPEM(certBytes)
	if !ok {
		log.Panic("failed to parse root certificate")
	}
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    clientCertPool,
	}
	ln, err := tls.Listen("tcp", ":443", config)
	if err != nil {
		log.Panic(err)
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Panic(err)
		}
		go handleConn(conn)
	}
}
