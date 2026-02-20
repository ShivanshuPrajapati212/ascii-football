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
			initGameWSHandler(conn, message)
		}
	}
}
