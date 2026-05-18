package handlers

import (
	"fmt"
	"os"
)

func GetColor(ansi string) string {
	color := map[string]string{
		"red":     "\033[31m",
		"green":   "\033[32m",
		"yellow":  "\033[33m",
		"blue":    "\033[34m",
		"magenta": "\033[35m",
		"cyan":    "\033[36m",
		"white":   "\033[37m",
	}

	code, ok := color[ansi]
	if !ok {
		fmt.Println("Error: color not found")
		os.Exit(2)
	}

	return code
}
