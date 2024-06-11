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

	args := os.Args[1:]
	if len(args) > 1 && strings.HasPrefix(args[0], "-") {
		if !IsColorFlagValid(args[0]) {
			fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
			return
		}
	}

	if strings.HasPrefix(args[0], "-") {
		colorFlag := flag.String("color", "", "Color to apply to the text")
		flag.Parse()

		if *colorFlag == "" {
			fmt.Println("Usage: go run . --color=<color> [letters to be colored] [STRING]")
			return
		}

		letters := ""
		text := ""

		if len(args) == 3 {
			letters = args[1]
			text = args[2]
		} else if len(args) == 2 {
			text = args[1]
		} else {
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

	} else if len(args) >= 2 {
		fileName := ""
		if len(args) == 3 {
			fileName = args[2]
		} else {
			fileName = args[1]
		}

		if !Filenamevalidate(fileName) {
			fmt.Println("INVALID FILE")
			os.Exit(0)
		}

		normalizedFileName := NormalizeFilename(fileName)
		data, err := os.ReadFile(normalizedFileName)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		var lines []string
		if normalizedFileName == "thinkertoy.txt" {
			lines = strings.Split(string(data), "\r\n")
			if len(lines) != 856 {
				fmt.Println("THE TEXT FILE IS INCORRECT")
				os.Exit(1)
			}
		} else {
			lines = strings.Split(string(data), "\n")
		}

		if len(args) < 1 {
			fmt.Println("Please provide text to display.")
			return
		}

		color.Display(strings.Join(args[:1], ""), lines)
	} else {
		fmt.Println("Usage: go run . [OPTION] [STRING]")
	}
}

