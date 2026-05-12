package main

import (
	"fmt"
	"os"
	"strings"
)

func ascii() {
	//! checks the LENGTH of the argument
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <string>")
		return
	}

	//^ store the argument in a variable
	args := os.Args[1]

	//! checks if the argument is EMPTY
	if args == "" {
		return
	}

	//^ reading the standard.txt file that returns a byte and error if there is any
	readingByte, err := (os.ReadFile("templates/standard.txt"))

	errorHandling(err)

	//^ converting byte to string
	readingStr := string(readingByte)
	readingStr = strings.ReplaceAll(readingStr, "\r\n", "\n")

	//! split the strings using NEWLINE
	splitStr := strings.Split(readingStr, "\n")

	//^ create a map for the LETTERS or ASCII
	charMap := map[rune][]string{}

	for i := 32; i <= 126; i++ {
		startLine := (i - 32) * 9
		charLines := splitStr[startLine+1 : startLine+9]
		charMap[rune(i)] = charLines
	}

	lines := strings.Split(args, "\\n")

	for i, line := range lines {
		if line == "" {
			if i < len(lines)-1 {
				fmt.Println()
			}
			continue
		}
		row := make([]string, 8)

		for _, char := range line {
			for i := 0; i < 8; i++ {
				position := coloredPositions(i)
				row[i] += charMap[char][i]
			}
		}
		for _, r := range row {
			fmt.Println(r)
		}
	}
}

func errorHandling(err error) {
	if err != nil {
		panic(err)
	}
}