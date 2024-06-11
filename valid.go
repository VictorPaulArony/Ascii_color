package main

import (
	"fmt"
	"regexp"
)

// check if the color flag format is a valid format
func IsColorFlagValid(mycolorflag string) bool {
	match, err := regexp.MatchString(`^--color=([a-z]|[A-Z])+$`, mycolorflag)
	if err != nil {
		fmt.Printf("WE HAVE THIS ERROR %v", err)
		return false
	}
	return match
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
