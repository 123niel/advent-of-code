package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func readFile(filename string) []string {
	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	lines := make([]string, 5)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func main() {
	lines := readFile("01-input.txt")

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
