package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	parseArgs := parseFlags(os.Args)
	color := parseArgs.color
	subStr := parseArgs.subStr
	sentence := parseArgs.sentence

	fmt.Println(color, subStr, sentence)
	fmt.Println(coloredIndices("a king kitten have kit", "kit"))
}

type FlagConfig struct {
	color    string
	subStr   string
	sentence string
}

func parseFlags(s []string) FlagConfig {
	switch {
	case len(s) == 2:
		return FlagConfig{
			sentence: s[1],
		}
	case len(s) == 4:
		arg := s[1]
		if strings.HasPrefix(arg, "--color=") {
			split := strings.Split(arg, "=")
			return FlagConfig{
				color:    split[1],
				subStr:   s[2],
				sentence: s[3],
			}
		}
		fmt.Println("Usage: go run . [OPTION] [STRING]\nEX: go run . --color=<color> <substring to be colored> \"something\"")
		os.Exit(1)
	default:
		fmt.Println("Usage: go run . [OPTION] [STRING]\nEX: go run . --color=<color> <substring to be colored> \"something\"")
		os.Exit(1)
	}
	return FlagConfig{}
}

func coloredIndices(sentence, subStr string) map[int]bool {
	found := make(map[int]bool)
	remaining := sentence
	offset := 0

	for {
		pos := strings.Index(remaining, subStr)
		if pos == -1 {
			break
		}
		for i := 0; i < len(subStr); i++ {
			found[pos+offset+i] = true
		}
		offset += pos + len(subStr)
		remaining = remaining[pos+len(subStr):]
	}
	return found
}
