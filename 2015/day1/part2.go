package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("error opening file")
		return
	}

	level := 0
	position := 1

	for _, num := range string(file) {
		switch num {
		case '(':
			level += 1
		case ')':
			level -= 1
		}

		if level == -1 {
			break
		}

		position += 1
	}

	fmt.Println(position)
}
