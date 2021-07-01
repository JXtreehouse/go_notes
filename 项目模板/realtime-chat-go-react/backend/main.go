/*
 * @Author: your name
 * @Date: 2021-07-01 17:09:44
 * @LastEditTime: 2021-07-01 21:11:06
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /go_notes/项目模板/realtime-chat-go-react/backend/main.go
 */
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)


// 我们需要定义一个Upgrader
// 我们需要定义读写缓存大小
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool { return true },
}

func reader(conn *websocket.Conn) {
	for {
		messageType, p , err := conn.ReadMessage()
		if err !=nil {
			log.Println(err)
			return 
		}

		fmt.Println(string(p))

		if err:= conn.WriteMessage(messageType, p); err!=nil {
			log.Println(err)
			return 
		}
	}
}

// define our WebSocket endpoint
func serveWs(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.Host)

  // upgrade this connection to a WebSocket
  // connection
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
  }
  // listen indefinitely for new messages coming
  // through on our WebSocket connection
    reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Simple Server")
  })
  // mape our `/ws` endpoint to the `serveWs` function
    http.HandleFunc("/ws", serveWs)
}
func main() {
	fmt.Println("Chat App v0.01")
    setupRoutes()
    http.ListenAndServe(":8080", nil)
}