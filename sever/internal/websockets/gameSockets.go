package websockets

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ShivanshuPrajapati212/ascii-football-server/internal/models"
)

var (
	onGoingGames []*models.Game
	watingGame   *models.Game
)

func GameWSHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		var message GameMessage

		if err := json.Unmarshal(p, &message); err != nil {
			fmt.Println("Error parsing json: ", err)
			continue
		}

		if message.Type == InitGameType {
			if watingGame == nil {
				player := models.Player{
					Name: message.Payload,
				}

				watingGame = &models.Game{
					TeamA:  []models.Player{player},
					Target: 7,
				}
			} else {
				player := models.Player{Name: message.Payload}
				if len(watingGame.TeamA) == 11 && len(watingGame.TeamB) == 10 {
					watingGame.TeamB = append(watingGame.TeamB, player)
					onGoingGames = append(onGoingGames, watingGame)
					watingGame = nil
				}
				if len(watingGame.TeamA) > len(watingGame.TeamB) {
					watingGame.TeamB = append(watingGame.TeamB, player)
				} else if len(watingGame.TeamA) < len(watingGame.TeamB) {
					watingGame.TeamA = append(watingGame.TeamA, player)
				}
			}
			fmt.Println(onGoingGames)
			fmt.Println(watingGame)
		}
	}
}
