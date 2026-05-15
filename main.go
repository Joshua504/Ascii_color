package main

import (
	"ascii-art-color/handlers"
	"fmt"
	"os"
)

func main() {
	parseArgs := handlers.ParseFlags(os.Args)
	color := parseArgs.Color
	subStr := parseArgs.SubStr
	sentence := parseArgs.Sentence

	fmt.Println(color, subStr, sentence)
	fmt.Println(handlers.ColoredIndices("a king kitten have kit", "kit"))
}
