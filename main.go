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
	if len(args) > 1 && strings.HasPrefix(args[0], "-") {
		if !IsColorFlagValid(args[0]) {
			os.Stdout.WriteString("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
		}
	}
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

		if *colorFlag == "" {
			fmt.Println("Usage: go run . --color=<color> [letters to be colored] [STRING]")
			return
		}

	} else if len(args) == 3 {
		if args[2] == "thinkertoy.txt" || args[2] == "standard.txt" || args[2] == "shadow.txt" || args[2] == "thinkertoy" || args[2] == "standard" || args[2] == "shadow" {
			inputfile := Filenamevalidate(args[2])
		}

		data, err := os.ReadFile(os.Args[2])
		if err != nil {
			println("INVALID FILE")
			os.Exit(0)
		}
		// exclusion for the thinkertoy.txt to remove the courage return
		if os.Args[2] == "thinkertoy.txt" {
			lines := strings.Split(string(data), "\r\n")

			if len(lines) != 856 {
				println("THE TEXT FILE IS INCORRECT")
				os.Exit(1)
			}
			if len(os.Args) < 2 {
				println("Please provide text to display.")
				return
			}
			color.Display(strings.Join(os.Args[1:2], ""), lines)
			// in the case of standard and shadoe text files
		} else {
			lines := strings.Split(string(data), "\n")
			color.Display(strings.Join(os.Args[1:2], ""), lines)
		}
	}
}
