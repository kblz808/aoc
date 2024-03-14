package main

import (
	"fmt"
	"os"
)

var houses = map[string]int{
	"00": 1,
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("error opening file")
		return
	}

	// [:len(file)-1] is to remove the newline at the end
	content := string(file[:len(file)-1])

	sx, sy := 0, 0
	rsx, rsy := 0, 0

	for i := 0; i < len(content); i += 2 {
		s := rune(content[i])
		calculateMove(s, &sx, &sy)

		if i+1 < len(content) {
			rs := rune(content[i+1])
			calculateMove(rs, &rsx, &rsy)
		}
	}

	fmt.Println(len(houses))
}

func calculateMove(move rune, x, y *int) {
	switch move {
	case '^':
		*y += 1
	case 'v':
		*y -= 1
	case '>':
		*x += 1
	case '<':
		*x -= 1
	}

	house := fmt.Sprintf("%d%d", *x, *y)

	houses[house] += 1
}
