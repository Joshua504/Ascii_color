package handlers

import "strings"

func ColoredIndices(sentence, subStr string) map[int]bool {
	found := make(map[int]bool)
	if subStr == "" {
		return found
	}
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
