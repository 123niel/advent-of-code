package main

import (
	"fmt"

	"github.com/123niel/advent-of-code/misc"
)

var line string

func main() {
	line = misc.ReadFile("input.txt")[0]

	fmt.Println(substringWithoutDuplicates(line, 4))
	fmt.Println(substringWithoutDuplicates(line, 14))
}

func substringWithoutDuplicates(str string, n int) int {
	for i := n; i < len(line); i++ {
		substr := line[i-n : i]

		if !hasDuplicates(substr) {
			return i
		}
	}

	return 0
}

func hasDuplicates(str string) bool {
	visited := make(map[rune]bool, 0)

	for _, c := range str {
		if !visited[c] {
			visited[c] = true
		} else {
			return true
		}
	}

	return false
}
