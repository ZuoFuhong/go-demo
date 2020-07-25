// Copyright (c) 2020 ZuoFuhong. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/gorilla/websocket"
	"html/template"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade: ", err)
		return
	}
	defer c.Close()

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Print("read: ", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write: ", err)
			break
		}
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	homeTemplate := template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<body>
    <div>
        <button id="send">发送</button>
    </div>
    <script>
        let ws = new WebSocket("{{.}}")
        ws.onopen = function (evt) {
            console.log("Connect open ...")
        }
        ws.onmessage = function (evt) {
            console.log("Received message: ", evt.data)
        }

        document.getElementById('send').addEventListener('click', function () {
            ws.send("hello server")
        })
    </script>
</body>
</html>
`))
	err := homeTemplate.Execute(w, "ws://"+r.Host+"/echo")
	if err != nil {
		panic(err)
	}
}

func main() {
	log.SetFlags(0)
	http.HandleFunc("/", home)
	http.HandleFunc("/echo", echo)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
