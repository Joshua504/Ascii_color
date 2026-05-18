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
	Banner   string
}

func ParseFlags(s []string) FlagConfig {
	switch {
	case len(s) == 2:
		return FlagConfig{
			Sentence: s[1],
		}
	case len(s) == 3:
		if strings.HasPrefix(s[1], "--color=") {
			split := strings.Split(s[1], "=")
			return FlagConfig{
				Color:    split[1],
				Sentence: s[2],
			}
		}
		return FlagConfig{
			Sentence: s[1],
			Banner:   s[2],
		}
	case len(s) == 4:
		if strings.HasPrefix(s[1], "--color=") {
			split := strings.Split(s[1], "=")
			bannerNames := map[string]bool{"standard": true, "shadow": true, "thinkertoy": true}
			if bannerNames[s[3]] {
				return FlagConfig{
					Color:    split[1],
					Sentence: s[2],
					Banner:   s[3],
				}
			}
			return FlagConfig{
				Color:    split[1],
				SubStr:   s[2],
				Sentence: s[3],
			}
		}
		fmt.Println("Usage: go run . [OPTION] [STRING]\nEX: go run . --color=<color> <substring to be colored> \"something\"")
		os.Exit(1)
	case len(s) == 5:
		if strings.HasPrefix(s[1], "--color=") {
			split := strings.Split(s[1], "=")
			return FlagConfig{
				Color:    split[1],
				SubStr:   s[2],
				Sentence: s[3],
				Banner:   s[4],
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
