package websockets

import (
	"fmt"

	"github.com/ShivanshuPrajapati212/ascii-football-server/internal/models"
	"github.com/gorilla/websocket"
)

func initGameWSHandler(conn *websocket.Conn, message GameMessage) {
	if watingGame == nil {
		player := models.Player{
			Socket: conn,
			Name:   message.Payload,
		}

		watingGame = &models.Game{
			TeamA:  []models.Player{player},
			Target: 7,
		}
	} else {
		player := models.Player{Socket: conn, Name: message.Payload}
		if len(watingGame.TeamA) == 11 && len(watingGame.TeamB) == 10 {
			watingGame.TeamB = append(watingGame.TeamB, player)
			onGoingGames = append(onGoingGames, watingGame)
			fmt.Println("New Game created: ", watingGame)
			watingGame = nil
		} else if len(watingGame.TeamA) > len(watingGame.TeamB) {
			watingGame.TeamB = append(watingGame.TeamB, player)
		} else if len(watingGame.TeamA) < len(watingGame.TeamB) {
			watingGame.TeamA = append(watingGame.TeamA, player)
		} else {
			watingGame.TeamA = append(watingGame.TeamA, player)
		}
	}
}
