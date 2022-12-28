package main

import (
	"fmt"

	"github.com/123niel/advent-of-code/misc"
)

var trees [][]int
var transposed [][]int

func main() {
	lines := misc.ReadFile("input.txt")

	for _, line := range lines {
		var row []int

		for _, char := range line {
			row = append(row, int(char)-'0')
		}
		trees = append(trees, row)
	}

	transposed = make([][]int, len(trees[0]))

	for i := range trees[0] {
		transposed[i] = make([]int, len(trees))
	}

	for i, row := range trees {
		for j, tree := range row {
			transposed[j][i] = tree
		}
	}

	fmt.Println(part1())
	fmt.Println(part2())
}

func part1() string {

	count := (len(trees)) * (len(trees))
	for i := 1; i < len(trees)-1; i++ {
		row := trees[i]
		for j := 1; j < len(row)-1; j++ {
			tree := row[j]
			if !visible(row[:j], tree) &&
				!visible(row[j+1:], tree) &&
				!visible(transposed[j][:i], tree) &&
				!visible(transposed[j][i+1:], tree) {
				count--
			}
		}
	}

	return fmt.Sprint(count)
}

func part2() string {
	bestScore := 0

	for i := 1; i < len(trees)-1; i++ {
		row := trees[i]
		for j := 1; j < len(row)-1; j++ {
			tree := row[j]
			coloumn := transposed[j]

			distanceLeft := 0
			for k := j - 1; k >= 0; k-- {
				distanceLeft++
				if row[k] >= tree {
					break
				}
			}

			distanceRight := 0
			for k := j + 1; k < len(row); k++ {
				distanceRight++
				if row[k] >= tree {
					break
				}
			}

			distanceUp := 0
			for k := i - 1; k >= 0; k-- {
				distanceUp++
				if coloumn[k] >= tree {
					break
				}
			}

			distanceDown := 0
			for k := i + 1; k < len(coloumn); k++ {
				distanceDown++
				if coloumn[k] >= tree {
					break
				}
			}

			score := distanceDown * distanceLeft * distanceUp * distanceRight
			if score > bestScore {
				bestScore = score
			}
		}
	}
	return fmt.Sprint(bestScore)
}

func visible(row []int, tree int) bool {
	for _, t := range row {
		if t >= tree {
			return false
		}
	}
	return true
}
