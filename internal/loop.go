package internal

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ShivanshuPrajapati212/ascii-football/ascii"
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
		fmt.Print("\033[2J\033[0;0H")
		fmt.Print(ascii.FootballFieldASCII)
		time.Sleep(time.Duration((1.0/float32(FPS))*float32(time.Second)) - time.Since(startTime))

	}
}
