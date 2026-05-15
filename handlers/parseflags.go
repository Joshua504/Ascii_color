package handlers

import (
	"fmt"
	"os"
	"strings"
)

type FlagConfig struct {
	Color    string
	SubStr   string
	Sentence string
}

func ParseFlags(s []string) FlagConfig {
	switch {
	case len(s) == 2:
		return FlagConfig{
			Sentence: s[1],
		}
	case len(s) == 4:
		arg := s[1]
		if strings.HasPrefix(arg, "--color=") {
			split := strings.Split(arg, "=")
			return FlagConfig{
				Color:    split[1],
				SubStr:   s[2],
				Sentence: s[3],
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
