package main

import (
	"fmt"
	"strings"

	"github.com/123niel/advent-of-code/misc"
)

func main() {
	lines := misc.ReadFile("input.txt")
	var sum int

	for _, line := range lines {
		splitted := strings.Split(line, " ")
		opp := splitted[0]
		self := splitted[1]

		sum += outcomeScore(opp, self) + shapeScore(self)
	}

	fmt.Println(sum)

	sum = 0

	for _, line := range lines {
		splitted := strings.Split(line, " ")
		opp := splitted[0]
		res := splitted[1]

		sum += calculateScore2(res) + caluclateSignScore2(calculateSign(opp, res))
	}

	fmt.Println(sum)
}

// A, X: Rock
// B, Y: Paper
// C, Z: Scissors
func outcomeScore(opp, self string) int {
	if opp == "A" {
		switch self {
		case "X":
			return 3
		case "Y":
			return 6
		default:
			return 0
		}
	} else if opp == "B" {
		switch self {
		case "X":
			return 0
		case "Y":
			return 3
		default:
			return 6
		}
	} else { // opp == C
		switch self {
		case "X":
			return 6
		case "Y":
			return 0
		default:
			return 3
		}
	}
}

func shapeScore(shape string) int {
	switch shape {
	case "X":
		return 1
	case "Y":
		return 2
	default:
		return 3
	}
}

func calculateSign(opp, res string) string {
	if res == "X" {
		switch opp {
		case "A":
			return "C"
		case "B":
			return "A"
		default:
			return "B"
		}
	} else if res == "Y" {
		return opp
	} else { //res == "Z"
		switch opp {
		case "A":
			return "B"
		case "B":
			return "C"
		default:
			return "A"
		}
	}
}

func caluclateSignScore2(sign string) int {
	switch sign {
	case "A":
		return 1
	case "B":
		return 2
	default:
		return 3
	}
}

func calculateScore2(res string) int {
	switch res {
	case "X":
		return 0
	case "Y":
		return 3
	default:
		return 6
	}
}
