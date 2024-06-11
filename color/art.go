package color

import (
	"fmt"
	"strings"
)

// DisplayText displays the provided text along with content lines
func Display(input string, contentLines []string) {
	if input == "" {
		return
	}

	if input == "\\n" || input == "\n" {
		fmt.Println()
		return
	}
	// make newline and tab printable inthe terminal output
	input = strings.ReplaceAll(input, "\n", "\\n")
	input = strings.ReplaceAll(input, "\\t", "\t")

	wordslice := strings.Split(input, "\\n")

	for _, word := range wordslice {
		if word == "" {
			fmt.Println()
		} else {
			if English(word) {
				PrintAsciiWord(word, contentLines)
			} else {
				fmt.Println("Invalid input:  not accepted")
				// Optionally continue processing other words
			}
		}
	}
}

// IsEnglish checks if a word contains only English alphabets
func English(words string) bool {
	for _, word := range words {
		if word < 32 || word > 126 {
			return false
		}
	}
	return true
}

// PrintWord prints a word if it exists in the content lines
func PrintAsciiWord(word string, contentLines []string) {
	linesOfSlice := make([]string, 9)

	for _, v := range word {
		for i := 1; i <= 9; i++ {
			linesOfSlice[i-1] +=  contentLines[int(v-32)*9+i]
		}
	}

	fmt.Print(strings.Join(linesOfSlice, "\n"))
}
