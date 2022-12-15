package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/123niel/advent-of-code/misc"
)

type move struct {
	Quant int
	From  int
	To    int
}

var lines []string
var moves []move

func readStacks(lines []string) []misc.Stack[rune] {
	stacks := make([]misc.Stack[rune], 9)

	for i := 7; i >= 0; i-- {
		line := lines[i]

		for j := 0; j < 9; j++ {
			char := rune(line[1+4*j])

			if char != ' ' {
				stacks[j].Push(char)
			}
		}
	}

	return stacks
}

func main() {
	lines = misc.ReadFile("inputs.txt")

	re := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

	for _, line := range lines[10:] {

		match := re.FindStringSubmatch(line)

		q, _ := strconv.Atoi(match[1])
		f, _ := strconv.Atoi(match[2])
		t, _ := strconv.Atoi(match[3])

		moves = append(moves, move{
			Quant: q,
			From:  f,
			To:    t,
		})

	}

	fmt.Println(part1())
	fmt.Println(part2())
}

func part1() string {

	stacks := readStacks(lines)

	for _, m := range moves {
		for i := 0; i < m.Quant; i++ {
			val, _ := stacks[m.From-1].Pop()

			stacks[m.To-1].Push(val)
		}
	}

	var tops string

	for _, s := range stacks {
		top, _ := s.Peek()
		tops += string(top)
	}

	return tops
}

func part2() string {
	stacks := readStacks(lines)

	for _, m := range moves {
		vals, _ := stacks[m.From-1].PopMutliple(m.Quant)
		stacks[m.To-1].PushMultiple(vals)
	}

	var tops string

	for _, s := range stacks {
		top, _ := s.Peek()
		tops += string(top)
	}

	return tops
}
