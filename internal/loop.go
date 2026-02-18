package internal

import (
	"log"
	"os"
	"time"

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
	for {
		startTime := time.Now()

		printField()

		time.Sleep(time.Duration((1.0/float32(FPS))*float32(time.Second)) - time.Since(startTime))

	}
}
