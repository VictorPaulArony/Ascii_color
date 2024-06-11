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
