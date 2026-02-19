package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ShivanshuPrajapati212/ascii-football-server/internal/websockets"
)

func StartServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to the club.")
	})

	http.HandleFunc("/helloHell", websockets.HelloHellConnection)
	http.HandleFunc("/ws", websockets.GameWSHandler)

	err := http.ListenAndServe(":42069", nil)
	if err != nil {
		log.Fatal("error starting server")
	}
}
