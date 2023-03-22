package utils

import "fmt"

const (
	RESET  = "\x1b[0m"
	GREEN  = "\x1b[32m"
	RED    = "\x1b[31m"
	YELLOW = "\x1b[33m"
	BLUE   = "\x1b[34m"
	PURPLE = "\x1b[35m"
	CYAN   = "\x1b[36m"
	GRAY   = "\x1b[37m"
	WHITE  = "\x1b[97m"
)

func Color(color string, text string) {
	fmt.Println(color + text + RESET)
}
