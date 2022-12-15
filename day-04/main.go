package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/123niel/advent-of-code/misc"
)

type p struct {
	start1 int
	end1   int
	start2 int
	end2   int
}

var lines []string
var pairs []p

func main() {
	lines = misc.ReadFile("input.txt")

	for _, line := range lines {
		pairs = append(pairs, parse(line))
	}

	fmt.Println(part1())
	fmt.Println(part2())
}

func part1() int {

	var sum int

	for _, pair := range pairs {

		if containing(pair) {
			sum++
		}
	}

	return sum
}

func part2() int {
	var sum int

	for _, pair := range pairs {
		if overlapping(pair) {
			sum++
		}
	}

	return sum
}

func parse(line string) p {
	elves := strings.Split(line, ",")

	elve1 := strings.Split(elves[0], "-")
	elve2 := strings.Split(elves[1], "-")

	start1, _ := strconv.Atoi(elve1[0])
	end1, _ := strconv.Atoi(elve1[1])

	start2, _ := strconv.Atoi(elve2[0])
	end2, _ := strconv.Atoi(elve2[1])

	pair := p{
		start1,
		end1,
		start2,
		end2,
	}

	return pair
}

func containing(pair p) bool {
	return (pair.start1 <= pair.start2 && pair.end1 >= pair.end2) || (pair.start1 >= pair.start2 && pair.end1 <= pair.end2)
}

func overlapping(pair p) bool {
	return (pair.start2 >= pair.start1 && pair.start2 <= pair.end1) ||
		(pair.start1 >= pair.start2 && pair.start1 <= pair.end2) ||
		(pair.end1 >= pair.start2 && pair.end1 <= pair.end2) ||
		(pair.end2 >= pair.start1 && pair.end2 <= pair.end1)
}
