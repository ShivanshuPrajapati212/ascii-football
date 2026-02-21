package gameloop

import (
	"fmt"

	"github.com/ShivanshuPrajapati212/ascii-football/ascii"
)

func printField() {
	fmt.Print("\033[2J\033[0;0H")
	fmt.Print(ascii.FootballFieldASCII)
}
