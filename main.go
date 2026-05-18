package main

import (
	"ascii-art-color/handlers"
	"fmt"
	"os"
	"strings"
)

func main() {
	config := handlers.ParseFlags(os.Args)

	banner := map[string]string{
		"standard":   "templates/standard.txt",
		"shadow":     "templates/shadow.txt",
		"thinkertoy": "templates/thinkertoy.txt",
	}

	bannerFile := banner[config.Banner]
	if bannerFile == "" {
		bannerFile = banner["standard"]
	}
	//^ reading the standard.txt file that returns a byte and error if there is any
	readingByte, err := (os.ReadFile(bannerFile))

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

	lines := strings.Split(config.Sentence, "\\n")
	subStr := config.SubStr
	if subStr == "" {
		subStr = config.Sentence
	}
	colored := handlers.ColoredIndices(config.Sentence, subStr)
	color := ""
	if config.Color != "" {
		color = handlers.GetColor(config.Color)
	}

	for i, line := range lines {
		if line == "" {
			if i < len(lines)-1 {
				fmt.Println()
			}
			continue
		}
		row := make([]string, 8)

		for index, char := range line {
			for i := 0; i < 8; i++ {
				if colored[index] {
					row[i] += color + charMap[char][i] + "\033[0m"
				} else {
					row[i] += charMap[char][i]
				}
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
