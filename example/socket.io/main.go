package main

import (
	"encoding/json"
	"fmt"
	engineio "github.com/googollee/go-engine.io"
	"github.com/googollee/go-engine.io/transport"
	"github.com/googollee/go-engine.io/transport/polling"
	"github.com/googollee/go-engine.io/transport/websocket"
	socketio "github.com/googollee/go-socket.io"
	"github.com/rs/cors"
	"log"
	"net/http"
)

/*
简易的WebRTC信令服务器
*/

type RoomArgs struct {
	UserId   string `json:"userId"`
	RoomName string `json:"roomName"`
}

func main() {
	pt := polling.Default
	wt := websocket.Default
	wt.CheckOrigin = func(req *http.Request) bool {
		return true
	}
	server, err := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			pt,
			wt,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})
	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})
	server.OnEvent("/", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		fmt.Println(msg)
		return "recv " + msg
	})

	// 加入房间
	server.OnEvent("/", "join-room", func(s socketio.Conn, msg string) string {
		var room RoomArgs
		err := json.Unmarshal([]byte(msg), &room)
		if err != nil {
			fmt.Println(err)
			return "ok"
		}
		fmt.Printf("join-room, user: %s, room: %s\n", room.UserId, room.RoomName)
		server.JoinRoom("/", room.RoomName, s)
		broadcastTo(server, s.Rooms(), "user-joined", room.UserId)
		return "ok"
	})

	// 离开房间
	server.OnEvent("/", "leave-room", func(s socketio.Conn, msg string) string {
		var room RoomArgs
		err := json.Unmarshal([]byte(msg), &room)
		if err != nil {
			fmt.Println(err)
			return "ok"
		}
		fmt.Printf("leave-room, user: %s, room: %s\n", room.UserId, room.RoomName)
		broadcastTo(server, s.Rooms(), "user-left", room.UserId)
		s.Leave(room.RoomName)
		return "ok"
	})

	// 广播
	server.OnEvent("/", "broadcast", func(s socketio.Conn, msg string) string {
		broadcastTo(server, s.Rooms(), "broadcast", msg)
		return "ok"
	})

	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		_ = s.Close()
		return last
	})
	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})
	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})
	go server.Serve()
	defer server.Close()

	mux := http.NewServeMux()
	mux.Handle("/socket.io/", server)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost"},
		AllowedMethods:   []string{"GET", "PUT", "OPTIONS", "POST", "DELETE"},
		AllowCredentials: true,
	})
	handler := c.Handler(mux)
	log.Println("Serving at localhost:8888...")
	log.Fatal(http.ListenAndServe(":8888", handler))
}

func broadcastTo(server *socketio.Server, rooms []string, event string, msg interface{}) {
	fmt.Printf("broadcastTo: \n\n%s\n\n", msg)
	if len(rooms) == 0 {
		fmt.Printf("broadcastTo error: not join room !\n")
		return
	}
	for _, room := range rooms {
		server.BroadcastToRoom("/", room, event, msg)
	}
}
