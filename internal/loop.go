package internal

import (
	"fmt"
	"log"
	"os"

	"github.com/ShivanshuPrajapati212/ascii-football/ascii"
	"golang.org/x/term"
)

func MainLoop() {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		log.Fatal("Error getting termianl size.")
	}
	if width < 120 || height < 30 {
		log.Fatal("Need bigger terminal, 120x30")
	}
	fmt.Print(ascii.FootballFieldASCII)
}
