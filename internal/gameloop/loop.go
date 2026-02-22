package gameloop

import (
	"log"
	"os"
	"time"

	"github.com/ShivanshuPrajapati212/ascii-football/internal/windows"
	"github.com/ShivanshuPrajapati212/ascii-football/internal/ws"
	"github.com/gorilla/websocket"
	"golang.org/x/term"
)

const FPS int = 24

func MainLoop() {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		log.Fatal("Error getting termianl size.")
	}
	if width < 120 || height < 30 {
		log.Fatal("Need bigger terminal, 120x30")
	}

	// Connect to WebSocket (persistent connection)
	ws.ConnectToWS()

	// Optional: Send initial message
	ws.GetManager().SendMessage(websocket.TextMessage, []byte("Hello, server!"))

	for {
		startTime := time.Now()

		// Get latest message from WebSocket
		_, latestMsg := ws.GetMessage()

		printField()
		windows.CreateAndShowWindow(30, 120, "Join A Game", string(latestMsg))

		time.Sleep(time.Duration((1.0/float32(FPS))*float32(time.Second)) - time.Since(startTime))
	}
}
