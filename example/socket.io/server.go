package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	engineio "github.com/googollee/go-engine.io"
	"github.com/googollee/go-engine.io/transport"
	"github.com/googollee/go-engine.io/transport/polling"
	"github.com/googollee/go-engine.io/transport/websocket"
	socketio "github.com/googollee/go-socket.io"
)

type RoomArgs struct {
	UserId   string `json:"userId"`
	RoomName string `json:"roomName"`
}

type Server struct {
	*socketio.Server
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	origin := r.Header.Get("Origin")
	header.Set("Access-Control-Allow-Origin", origin)
	header.Set("Access-Control-Allow-Headers", "*")
	header.Set("Access-Control-Allow-Credentials", "true")
	header.Set("Access-Control-Allow-Methods", "POST,GET,DELETE,PUT,OPTIONS")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	s.Server.ServeHTTP(w, r)
}

func NewServer() *Server {
	pt := polling.Default
	wt := websocket.Default
	wt.CheckOrigin = func(req *http.Request) bool {
		return true
	}
	socketServer, err := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			pt,
			wt,
		},
	})
	if err != nil {
		panic(err)
	}
	return &Server{
		socketServer,
	}
}

func AddMonitorService() {
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Status: OK"))
	})
}

func AddSignalService() {
	server := NewServer()
	server.OnConnect("/", func(s socketio.Conn) error {
		fmt.Println("新用户 " + s.ID() + " 上线")
		return nil
	})
	server.OnDisconnect("/", func(conn socketio.Conn, s string) {
		fmt.Println("用户 " + conn.ID() + " 离线")
	})
	server.OnEvent("/", "msg", func(s socketio.Conn, msg string) {
		fmt.Println("服务器收到 用户"+s.ID()+" 发来的消息: ", msg)
	})
	server.OnEvent("/", "join-room", func(s socketio.Conn, data string) {
		room := &RoomArgs{}
		err := json.Unmarshal([]byte(data), &room)
		if err != nil {
			panic(err)
		}
		fmt.Println("用户 " + room.UserId + " 加入房间 " + room.RoomName)
		s.Join(room.RoomName)
		broadcastTo(server, s.Rooms(), "user-joined", room.UserId)
	})
	server.OnEvent("/", "leave-room", func(s socketio.Conn, data string) {
		room := &RoomArgs{}
		err := json.Unmarshal([]byte(data), &room)
		if err != nil {
			panic(err)
		}
		fmt.Println("用户 " + room.UserId + " 离开房间 " + room.RoomName)
		s.Leave(room.RoomName)
		broadcastTo(server, s.Rooms(), "user-left", room.UserId)
	})
	server.OnEvent("/", "broadcast", func(s socketio.Conn, msg interface{}) {
		fmt.Println("用户 "+s.ID()+" 发送了广播消息：", msg)
		broadcastTo(server, s.Rooms(), "broadcast", msg)
	})
	server.OnEvent("/", "bye", func(s socketio.Conn) {
		last := s.Context().(string)
		s.Emit("bye", last)
		_ = s.Close()
	})
	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})
	go func() {
		_ = server.Serve()
		//defer server.Close()
	}()
	http.Handle("/socket.io/", server)
}

func broadcastTo(server *Server, rooms []string, event string, msg interface{}) {
	for _, room := range rooms {
		server.BroadcastToRoom("/", room, event, msg)
	}
}

func main() {
	AddMonitorService()
	AddSignalService()
	err := http.ListenAndServe("127.0.0.1:8888", nil)
	if err != nil {
		panic(err)
	}
}
