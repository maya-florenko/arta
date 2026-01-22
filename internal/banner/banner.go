package banner

import (
	_ "embed"
	"fmt"
)

//go:embed ascii.txt
var ascii string

func Print() {
	fmt.Print("\033[H\033[2J")
	fmt.Print(ascii)
	fmt.Print("\nProject: Arta\n")
	fmt.Print("Author: Mayusha256\n\n")
}
