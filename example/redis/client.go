package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"net"
	"strconv"
)

// Simple Redis client
// [Redis Protocol specification](https://redis.io/topics/protocol)
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatal(err)
	}
	// auth 123456
	_, _ = conn.Write([]byte("*2\r\n$4\r\nauth\r\n$6\r\n123456\r\n"))
	readReply(&conn)

	// set name mars
	_, _ = conn.Write([]byte("*3\r\n$3\r\nset\r\n$4\r\nname\r\n$4\r\nmars\r\n"))
	readReply(&conn)

	// get name
	_, _ = conn.Write([]byte("*2\r\n$3\r\nget\r\n$4\r\nname\r\n"))
	readReply(&conn)
}

func readReply(conn *net.Conn) {
	rd := bufio.NewReader(*conn)
	line, isPrefix, err := rd.ReadLine()
	if err != nil {
		log.Fatal(err)
	}
	if isPrefix {
		log.Fatal(bufio.ErrBufferFull)
	}
	if len(line) == 0 {
		log.Fatal(errors.New("reply is Empty"))
	}
	switch line[0] {
	case '+':
		log.Println(string(line[1:]))
	case '-':
		log.Println("error: ", string(line[1:]))
	case '$':
		bodyLen, err := strconv.Atoi(string(line[1:]))
		if err != nil {
			log.Fatal(err)
		}
		b := make([]byte, bodyLen+2)
		_, _ = io.ReadFull(rd, b)
		log.Println(string(b[:bodyLen]))
	}
}
