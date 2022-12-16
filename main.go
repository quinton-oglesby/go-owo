package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/pterm/pterm"
	"github.com/varsapphire/OwO/generate"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Command usage: `OwO some string to convert`")
	}

	toShow := pterm.FgMagenta.Sprint(generate.WhatsThis(strings.Join(os.Args[1:], " ")))
	fmt.Println(toShow)
}
