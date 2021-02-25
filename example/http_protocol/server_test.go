package http_protocol

import "testing"

func Test_Server(t *testing.T) {
	server := NewServer("127.0.0.1:8080")
	server.Start()
}
