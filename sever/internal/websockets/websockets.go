package websockets

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// CheckOrigin allows us to accept requests from different domains
	CheckOrigin: func(r *http.Request) bool { return true },
}

func HelloHellConnection(w http.ResponseWriter, r *http.Request) {
	// 1. Upgrade initial GET request to a WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	for {
		// 2. Read message from browser
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		// 3. Print and write message back to browser
		fmt.Printf("Received: %s\n", string(p))
		if err := conn.WriteMessage(messageType, []byte("Hello Hell")); err != nil {
			fmt.Println(err)
			return
		}
	}
}
