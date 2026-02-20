package websockets

import (
	"fmt"

	"github.com/ShivanshuPrajapati212/ascii-football-server/internal/models"
	"github.com/gorilla/websocket"
)

func sendBulkMessage(conns []*websocket.Conn, typeOfMessage string, payload string) {
	for _, conn := range conns {
		if err := conn.WriteJSON(GameMessage{Type: typeOfMessage, Payload: payload}); err != nil {
			fmt.Println("error sendiing message: ", err)
		}
	}
}

func getAllSockets(game *models.Game) []*websocket.Conn {
	var sockets []*websocket.Conn
	for _, v := range game.TeamA {
		sockets = append(sockets, v.Socket)
	}
	for _, v := range game.TeamB {
		sockets = append(sockets, v.Socket)
	}

	return sockets
}
