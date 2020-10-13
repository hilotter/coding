package main

import "fmt"

func urlify(s string, length int) string {
	replacedLength := 0

	r := []rune(s)
	for i := 0; i < length; i++ {
		char := r[i]
		if char == ' ' {
			replacedLength += 3
		} else {
			replacedLength++
		}
	}

	replacedR := make([]rune, replacedLength)
	count := 0
	for i := 0; i < length; i++ {
		char := r[i]
		if char == ' ' {
			replacedR[count] = '%'
			replacedR[count+1] = '2'
			replacedR[count+2] = '0'
			count += 3
		} else {
			replacedR[count] = char
			count++
		}
	}

	return string(replacedR)
}

func formatAnswer(s string, length int) {
	fmt.Printf("%s\n", urlify(s, length))
}

func main() {
	formatAnswer("hello world", 11)
	formatAnswer("this is test ", 13)
}
