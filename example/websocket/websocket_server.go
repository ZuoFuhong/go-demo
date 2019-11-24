package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// 协议升级
func wsHandle(resp http.ResponseWriter, request *http.Request) {
	var (
		conn *websocket.Conn
		data []byte
		err  error
	)
	conn, err = upgrader.Upgrade(resp, request, nil)
	if err != nil {
		return
	}
	for {
		if _, data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		fmt.Println("服务端收到消息：", string(data))
		if err = conn.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}
ERR:
	conn.Close()
}

func main() {
	http.HandleFunc("/ws", wsHandle)
	fmt.Println("websocket server runs at 127.0.0.1:8888")
	_ = http.ListenAndServe("127.0.0.1:8888", nil)
}
