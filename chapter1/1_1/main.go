package main

import "fmt"

func isUniqueString(s string) bool {
	counts := make(map[rune]int)
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		char := r[i]
		counts[char]++

		if counts[char] > 1 {
			return false
		}
	}
	return true
}

func formatAnswer(s string) {
	res := "not unique"
	if isUniqueString(s) {
		res = "unique"
	}
	fmt.Printf("%s is %s\n", s, res)
}

func main() {
	formatAnswer("hi")
	formatAnswer("uniq")
	formatAnswer("hello")
	formatAnswer("test")
}
