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
	lines := misc.ReadFile("input.txt")

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

	visited := make(map[vec]bool)

	h := vec{X: 0, Y: 0}
	t := vec{X: 0, Y: 0}

	for _, move := range moves {
		for i := 0; i < move.Count; i++ {
			h.X += move.DX
			h.Y += move.DY

			pull(&h, &t)

			visited[t] = true
		}
	}

	return fmt.Sprint(len(visited))
}

func part2() string {
	rope := make([]*vec, 10)
	for i := 0; i < 10; i++ {
		rope[i] = &vec{X: 0, Y: 0}
	}

	visited := make(map[vec]bool)

	for _, move := range moves {
		// printRope(rope)
		for i := 0; i < move.Count; i++ {
			rope[0].X += move.DX
			rope[0].Y += move.DY

			for i := 0; i < 9; i++ {
				pull(rope[i], rope[i+1])
			}

			visited[*rope[9]] = true
		}
	}

	return fmt.Sprint(len(visited))
}

func pull(h, t *vec) {
	diffX := h.X - t.X
	diffY := h.Y - t.Y

	if diffX == 2 {
		t.X++
		if diffY >= 1 {
			t.Y++
		} else if diffY <= -1 {
			t.Y--
		}
	} else if diffX == -2 {
		t.X--
		if diffY >= 1 {
			t.Y++
		} else if diffY <= -1 {
			t.Y--
		}
	} else if diffY == 2 {
		t.Y++
		if diffX >= 1 {
			t.X++
		} else if diffX <= -1 {
			t.X--
		}
	} else if diffY == -2 {
		t.Y--
		if diffX >= 1 {
			t.X++
		} else if diffX <= -1 {
			t.X--
		}
	}
}

func printRope(rope []*vec) {
	var maxX, minX, maxY, minY int
	for _, knot := range rope {
		if knot.X > maxX {
			maxX = knot.X
		}
		if knot.X < minX {
			minX = knot.X
		}
		if knot.Y > maxY {
			maxY = knot.Y
		}
		if knot.Y < minY {
			minY = knot.Y
		}
	}

	var output string

	for y := maxY; y >= minY; y-- {
		for x := minX; x <= maxX; x++ {
			char := "."
			if y == 0 && x == 0 {
				char = "S"
			}
			for i := len(rope) - 1; i >= 0; i-- {
				knot := *rope[i]
				if x == knot.X && y == knot.Y {
					char = fmt.Sprint(i)
				}
			}
			output += string(char)
		}
		output += "\n"
	}
	fmt.Println(output)
}
