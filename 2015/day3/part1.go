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

	x, y := 0, 0

	houses := map[string]int{
		"00": 1,
	}

	for _, val := range string(file) {
		switch val {
		case '^':
			y += 1
		case 'v':
			y -= 1
		case '>':
			x += 1
		case '<':
			x -= 1
		}

		house := fmt.Sprintf("%d%d", x, y)
		houses[house] += 1
	}

	fmt.Println(len(houses))
}
