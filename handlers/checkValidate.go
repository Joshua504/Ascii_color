package main

import (
	"fmt"
	"strings"
)

func checkValidate(s string) (string, bool) {
	color := ""
	var valid bool

	if strings.HasPrefix(s, "--color=") {
		splitFlag := strings.Split(s, "=")
		color = splitFlag[1]
		valid = true
	} else {
		valid = false
		fmt.Println("Error: not a color")
		
	}

	return color, valid
}