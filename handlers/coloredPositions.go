package main

func coloredPositions(s, subStr string) map[int]bool{
	positions := map[int]bool{}
	indices := findCharacter(s, subStr)

	for _, start := range indices{
		for j := 0; j < len(subStr); j++ {
			positions[start+j] = true
		}
	}
	return positions
}