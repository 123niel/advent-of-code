package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/123niel/advent-of-code/misc"
)

var lines []string

func main() {
	lines = misc.ReadFile("input.txt")

	fmt.Println(part1())
	fmt.Println(part2())
}

func part1() string {

	cycles := []int{20, 60, 100, 140, 180, 220}
	cycle := 0
	x := 1
	sum := 0

	for _, line := range lines {
		cycle++
		if contains(cycles, cycle) {
			sum += cycle * x
		}

		if strings.HasPrefix(line, "addx") {
			cycle++
			if contains(cycles, cycle) {
				sum += cycle * x
			}
			args := strings.Split(line, " ")
			v, _ := strconv.Atoi(args[1])

			x += v
		}
	}

	return fmt.Sprint(sum)
}

func part2() string {

	return fmt.Sprint()
}

func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
