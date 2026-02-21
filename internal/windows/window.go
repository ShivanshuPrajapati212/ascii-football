package windows

import (
	"fmt"
	"strings"
)

func CreateAndShowWindow(fieldHeight, fieldWidth int, title, body, cta string) {
	var maxLen string

	if len(title) < len(body) {
		maxLen = body
	} else {
		maxLen = title
	}

	for i := -2; i <= 6; i++ {
		fmt.Printf("\033[%d;%dH%s", fieldHeight/2+i, fieldWidth/2-len(maxLen)/2, strings.Repeat(" ", len(maxLen)))
		fmt.Printf("\033[%d;%dH%s", fieldHeight/2+i, fieldWidth/2-len(maxLen)/2-2, "#")
		fmt.Printf("\033[%d;%dH%s", fieldHeight/2+i, fieldWidth/2+len(maxLen)/2+2, "#")
	}

	fmt.Printf("\033[%d;%dH%s", fieldHeight/2-2, fieldWidth/2-len(maxLen)/2-1, strings.Repeat("#", len(maxLen)+3))
	fmt.Printf("\033[%d;%dH%s", fieldHeight/2+6, fieldWidth/2-len(maxLen)/2-1, strings.Repeat("#", len(maxLen)+3))

	fmt.Printf("\033[%d;%dH%s", fieldHeight/2, fieldWidth/2-len(title)/2, title)
	fmt.Printf("\033[%d;%dH%s", fieldHeight/2+2, fieldWidth/2-len(body)/2, body)
	fmt.Printf("\033[%d;%dH%s", fieldHeight/2+4, fieldWidth/2-len(cta)/2, cta)
}
