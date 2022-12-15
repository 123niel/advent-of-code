package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/123niel/advent-of-code/misc"
)

type command struct {
	input  []string
	output [][]string
}

type dir struct {
	name    string
	parent  *dir
	subdirs map[string]*dir
	files   []file
}

type file struct {
	name string
	size int
}

func newDir(name string, parent *dir) *dir {
	return &dir{
		name,
		parent,
		make(map[string]*dir),
		make([]file, 0),
	}
}

func newFile(name string, size int) *file {
	return &file{name, size}
}

var lines []string
var root *dir

func main() {
	lines = misc.ReadFile("input.txt")
	root = parse()

	fmt.Println(part1())
	fmt.Println(part2())
}

func part1() string {
	sizes := make([]int, 0)
	calcSubDirSizes(root, &sizes)

	var sum int

	for _, size := range sizes {
		if size <= 100000 {
			sum += size
		}
	}

	return fmt.Sprint(sum)
}

func part2() string {
	total := 70000000
	needed := 30000000
	used := calcDirSize(root)

	sizes := make([]int, 0)
	calcSubDirSizes(root, &sizes)

	sort.Ints(sizes)

	for _, size := range sizes {
		if size > (needed)-(total-used) {
			return fmt.Sprint(size)
		}
	}

	return fmt.Sprint()
}

func parse() *dir {

	var commands []command
	i := -1

	for _, line := range lines {
		fragments := strings.Split(line, " ")

		if fragments[0] == "$" {
			i++
			commands = append(commands, command{
				input: fragments[1:],
			})
		} else {
			commands[i].output = append(commands[i].output, fragments)
		}
	}

	root := newDir("/", nil)
	current := root

	for _, command := range commands[1:] {
		if command.input[0] == "cd" {
			if command.input[1] == ".." {
				current = current.parent
			} else {
				current = current.subdirs[command.input[1]]
			}
		} else {
			for _, output := range command.output {
				if output[0] == "dir" {
					name := output[1]
					d := newDir(name, current)
					current.subdirs[name] = d
				} else {
					size, _ := strconv.Atoi(output[0])
					current.files = append(current.files, *newFile(output[1], size))
				}
			}
		}
	}

	return root
}

func calcSubDirSizes(d *dir, sizes *[]int) {
	*sizes = append(*sizes, calcDirSize(d))

	for _, sub := range d.subdirs {
		calcSubDirSizes(sub, sizes)
	}
}

func calcDirSize(d *dir) int {
	var size int

	for _, file := range d.files {
		size += file.size
	}

	for _, sub := range d.subdirs {
		size += calcDirSize(sub)
	}

	return size
}
