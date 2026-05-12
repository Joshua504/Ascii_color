package main

import "strings"

func findCharacter(s, subStr string) []int {
	indices := []int{}
	offset := 0
	for {
		idx := strings.Index(s[offset:], subStr)

		if idx == -1{
			break
		}

		indices = append(indices, idx+offset)

		offset += idx + 1
	}
	return indices
}