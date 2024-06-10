package color

import (
	"fmt"
	"os"
	"strings"
)

// Define ANSI escape codes for colors
var colorMap = map[string]string{
	"red":    "\033[31m",
	"green":  "\033[32m",
	"yellow": "\033[33m",
	"blue":   "\033[34m",
	"purple": "\033[35m",
	"cyan":   "\033[36m",
	"white":  "\033[37m",
	"reset":  "\033[0m", // Reset color
}

// DisplayText displays the provided text along with content lines
func DisplayText(input string, contentLines []string, color string, letters string) {
	if input == "" {
		return
	}

	// Make newline and tab printable in the terminal output
	input = strings.ReplaceAll(input, "\n", "\\n")
	input = strings.ReplaceAll(input, "\t", "\\t")

	wordSlice := strings.Split(input, "\\n")

	for _, word := range wordSlice {
		if word == "" {
			fmt.Println()
		} else {
			if IsEnglish(word) {
				PrintWord(word, contentLines, color, letters)
			} else {
				fmt.Print("Invalid input: not accepted")
				os.Exit(0)
			}
		}
	}
}

// IsEnglish checks if a word contains only printable ASCII characters
func IsEnglish(word string) bool {
	for _, char := range word {
		if char < 32 || char > 126 {
			return false
		}
	}
	return true
}

// PrintWord prints a word if it exists in the content lines
func PrintWord(word string, contentLines []string, color string, letters string) {
	linesOfSlice := make([]string, 9)

	for _, v := range word {
		for i := 1; i < 9; i++ {
			charLine := contentLines[int(v-32)*9+i]
			if letters == "" || strings.ContainsRune(letters, v) {
				linesOfSlice[i-1] += ApplyColor(charLine, letters, color)
			} else {
				linesOfSlice[i-1] += charLine
			}
		}
	}

	fmt.Print(strings.Join(linesOfSlice, "\n"))
}

// ApplyColor applies the specified color to the text
func ApplyColor(text, letters, color string) string {
	colorCode, exists := colorMap[color]
	if !exists {
		fmt.Println("Invalid color specified.")
		os.Exit(0)
	}
	return fmt.Sprintf("%s%s%s", colorCode, text, colorMap["reset"])
}
