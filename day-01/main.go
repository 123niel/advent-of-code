package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/123niel/advent-of-code/misc"
)

func main() {
	lines := misc.ReadFile("inputs/01.txt")

	elves := make([]int, 10)
	currentElve := 0

	for _, line := range lines {

		if line != "" {
			num, _ := strconv.Atoi(line)
			currentElve = currentElve + num
		} else {
			elves = append(elves, currentElve)
			currentElve = 0
		}

	}

	sort.Ints(elves)

	len := len(elves)

	fmt.Println(elves[len-1])
	fmt.Println(elves[len-1] + elves[len-2] + elves[len-3])

}
