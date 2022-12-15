package main

import (
	"fmt"
	"strings"

	"github.com/123niel/advent-of-code/misc"
)

func main() {
	packs := misc.ReadFile("input.txt")

	prios := make(map[rune]int)

	for r, i := 'a', 1; r <= 'z'; r, i = r+1, i+1 {
		prios[r] = i
	}
	for r, i := 'A', 27; r <= 'Z'; r, i = r+1, i+1 {
		prios[r] = i
	}

	fmt.Println(part1(packs, prios))
	fmt.Println(part2(packs, prios))
}

func part1(packs []string, prios map[rune]int) int {

	var sum int

	for _, pack := range packs {
		first := pack[:len(pack)/2]
		second := pack[len(pack)/2:]

		for _, item := range first {
			if strings.Contains(second, string(item)) {
				sum += prios[item]
				break
			}
		}

	}

	return sum
}

func part2(packs []string, prios map[rune]int) int {
	var sum int

	for i := 0; i < len(packs); i += 3 {
		pack1 := packs[i]
		pack2 := packs[i+1]
		pack3 := packs[i+2]

		for _, item := range pack1 {
			if strings.Contains(pack2, string(item)) && strings.Contains(pack3, string(item)) {
				sum += prios[item]
				break
			}
		}
	}

	return sum
}
