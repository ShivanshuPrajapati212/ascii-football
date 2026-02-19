package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to the club.")
	})

	err := http.ListenAndServe(":42069", nil)
	if err != nil {
		log.Fatal("error starting server")
	}
}
