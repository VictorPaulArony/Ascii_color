package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"Ascii_color/color"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . [OPTION] [STRING]")
		return
	}

	// Define the color flag
	colorFlag := flag.String("color", "", "Color to apply to the text")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: go run . [OPTION] [STRING]")
		return
	}

	text := args[len(args)-1]
	letters := ""
	if len(args) > 1 {
		letters = args[len(args)-2]
	}

	if *colorFlag == "" {
		fmt.Println("Usage: go run . --color=<color> [letters to be colored] [STRING]")
		return
	}

	data, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	lines := strings.Split(string(data), "\n")
	color.DisplayText(text, lines, *colorFlag, letters)
}
