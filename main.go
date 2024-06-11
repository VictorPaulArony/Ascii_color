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

	// args := flag.Args()
	args := os.Args[1:]
	if strings.HasPrefix(args[0], "-") {

		// Define the color flag
		colorFlag := flag.String("color", "", "Color to apply to the text")
		flag.Parse()
		// make the arguments dynamic to run all the ascii projects
		if len(args) < 1 {
			fmt.Println("Usage: go run . [OPTION] [STRING]")
			return
		}

		letters := ""
		data, err := os.ReadFile("standard.txt")
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		if len(args) == 3 {
			text := args[2]
			letters = args[1]
			lines := strings.Split(string(data), "\n")
			color.DisplayText(text, lines, *colorFlag, letters)
		} else if len(args) == 2 {
			text := args[1]
			letters = ""
			lines := strings.Split(string(data), "\n")
			color.DisplayText(text, lines, *colorFlag, letters)
		}

		// if len(args) > 1 {
		// 	letters = args[1]
		// }

		if *colorFlag == "" {
			fmt.Println("Usage: go run . --color=<color> [letters to be colored] [STRING]")
			return
		}

		
	}
}
