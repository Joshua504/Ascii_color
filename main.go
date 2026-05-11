package main

import (
	"fmt"
	"os"
)

func main() {

	flag := os.Args[1]

	switch len(os.Args) {
	case 2:
	case 3:
		color, ok := checkValidate(flag)
		if !ok {
			fmt.Println("Usage: go run . [OPTION] [STRING]")
			return
		}
	case 4:
		color, ok := checkValidate(flag)
		if !ok {
			fmt.Println("Usage: go run . [OPTION] [STRING]")
			return
		}

	default:
		fmt.Println("Error: not enough arguement")
		return
	}
}
