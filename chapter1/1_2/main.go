package main

import "fmt"

func runesCountMap(s string) map[rune]int {
	counts := make(map[rune]int)
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		char := r[i]
		counts[char]++
	}
	return counts
}

func isRearrangeString(source string, target string) bool {
	if len(source) != len(target) {
		return false
	}

	sourceCounts := runesCountMap(source)
	targetCounts := runesCountMap(target)

	for sourceKey, sourceCount := range sourceCounts {
		targetCount := targetCounts[sourceKey]
		if sourceCount != targetCount {
			return false
		}
	}
	return true
}

func formatAnswer(source string, target string) {
	res := "not rearrange string"
	if isRearrangeString(source, target) {
		res = "rearrange string"
	}
	fmt.Printf("%s and %s is %s\n", source, target, res)
}

func main() {
	formatAnswer("hi", "he")
	formatAnswer("uniq", "quin")
	formatAnswer("hello", "leloh")
	formatAnswer("abc", "bcad")
}
