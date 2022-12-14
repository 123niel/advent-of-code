package misc

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	var stack Stack[rune]

	stack.Push('T')
	stack.Push('A')

	fmt.Println(stack)
}
