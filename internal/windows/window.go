package windows

import "fmt"

func CreateAndShowWindow(fieldHeight, fieldWidth int, title, body, cta string) {
	var i int
	fmt.Printf("\033[%d;%dH%s", fieldHeight/2+i, fieldWidth/2-len(title)/2, title)
	i++
	fmt.Printf("\033[%d;%dH%s", fieldHeight/2+i, fieldWidth/2-len(body)/2, body)
	i++
	fmt.Printf("\033[%d;%dH%s", fieldHeight/2+i, fieldWidth/2-len(cta)/2, cta)
}
