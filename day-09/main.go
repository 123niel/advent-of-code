package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/123niel/advent-of-code/misc"
)

type vec struct {
	X, Y int
}

type move struct {
	DX, DY int
	Count  int
}

var moves []move

func main() {
	lines := misc.ReadFile("input_.txt")

	for _, line := range lines {
		tmp := strings.Split(line, " ")

		dir := tmp[0]
		count, _ := strconv.Atoi(tmp[1])

		switch dir {
		case "R":
			moves = append(moves, move{DX: 1, DY: 0, Count: count})
		case "U":
			moves = append(moves, move{DX: 0, DY: 1, Count: count})
		case "L":
			moves = append(moves, move{DX: -1, DY: 0, Count: count})
		default:
			moves = append(moves, move{DX: 0, DY: -1, Count: count})
		}
	}

	fmt.Println(part1())
	fmt.Println(part2())
}

func part1() string {
	h := vec{X: 0, Y: 0}
	t := vec{X: 0, Y: 0}

	for _, move := range moves {
		for i := 0; i < move.Count; i++ {
			h.X += move.DX
			h.Y += move.DY

			diffX := h.X - t.X
			diffY := h.Y - t.Y

			if diffX == 2 {
				t.X++
				if diffY == 1 {
					t.Y++
				}
			}

		}
	}

	return fmt.Sprint()
}

func part2() string {

	return fmt.Sprint()
}
