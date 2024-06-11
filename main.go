package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"Ascii_color/color"
)

func main() {
	// Check if there are enough command-line arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . [OPTION] [STRING]")
		return
	}

	args := os.Args[1:]

	// Check for color option and validate it
	if len(args) > 1 && strings.HasPrefix(args[0], "-") {
		if !IsColorFlagValid(args[0]) {
			fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
			return
		}
	}

	// If the first argument is a flag (color option)
	if strings.HasPrefix(args[0], "-") {
		colorFlag := flag.String("color", "", "Color to apply to the text")
		flag.Parse()

		// Ensure color flag is provided
		if *colorFlag == "" {
			fmt.Println("Usage: go run . --color=<color> [letters to be colored] [STRING]")
			return
		}

		letters := ""
		text := ""

		// Parse additional arguments for letters and text
		if len(args) == 3 {
			letters = args[1]
			text = args[2]
		} else if len(args) == 2 {
			text = args[1]
		} else {
			fmt.Println("Usage: go run . --color=<color> [letters to be colored] [STRING]")
			return
		}

		// Read the standard text file
		data, err := os.ReadFile("standard.txt")
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		lines := strings.Split(string(data), "\n")

		// Display the text with the specified color
		color.DisplayText(text, lines, *colorFlag, letters)

	} else if len(args) >= 2 {
		// If there are at least two arguments, process the input file and text
		fileName := ""
		if len(args) == 3 {
			fileName = args[2]
		} else {
			fileName = args[1]
		}

		// Validate the filename
		if !Filenamevalidate(fileName) {
			fmt.Println("INVALID FILE")
			os.Exit(0)
		}

		// Normalize the filename to the correct format
		normalizedFileName := NormalizeFilename(fileName)

		// Read the normalized text file
		data, err := os.ReadFile(normalizedFileName)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		var lines []string

		// Handle specific file types differently
		if normalizedFileName == "thinkertoy.txt" {
			lines = strings.Split(string(data), "\r\n")
			if len(lines) != 856 {
				fmt.Println("THE TEXT FILE IS INCORRECT")
				os.Exit(1)
			}
		} else {
			lines = strings.Split(string(data), "\n")
		}

		// Ensure there is text to display
		if len(args) < 1 {
			fmt.Println("Please provide text to display.")
			return
		}

		// Display the text without color
		color.Display(strings.Join(args[:1], ""), lines)
	} else {
		// If arguments do not match the required pattern, show usage message
		fmt.Println("Usage: go run . [OPTION] [STRING]")
	}
}

// Filenamevalidate checks if the provided filename is valid
func Filenamevalidate(m string) bool {
	return filenameExist(m)
}

// NormalizeFilename converts the provided filename to its standard form
func NormalizeFilename(m string) string {
	if m == "shadow" || m == "shadow.txt" {
		return "shadow.txt"
	} else if m == "thinkertoy" || m == "thinkertoy.txt" {
		return "thinkertoy.txt"
	} else {
		return "standard.txt"
	}
}

// filenameExist checks if the filename exists in the allowed set
func filenameExist(m string) bool {
	return m == "shadow.txt" || m == "thinkertoy.txt" || m == "standard.txt" || m == "shadow" || m == "thinkertoy" || m == "standard"
}
