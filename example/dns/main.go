package main

import (
	"log"
	"net"
)

func main() {
	// 解析cname
	cname, err := net.LookupCNAME("www.baidu.com")
	if err != nil {
		log.Printf("LookupCNAME err: %v", err)
		return
	}
	log.Printf("cname = %s\n", cname)
	// 解析ip地址
	host, err := net.LookupHost("www.baidu.com")
	if err != nil {
		log.Printf("LookupHost err: %v", err)
		return
	}
	log.Printf("host = %s\n", host)
}
