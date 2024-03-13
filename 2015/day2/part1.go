package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("error opening file")
		return
	}

	total := 0

	for _, line := range strings.Split(string(file), "\n") {
		values := strings.Split(line, "x")

		if len(values) != 3 {
			continue
		}

		l, err := strconv.Atoi(values[0])
		w, err := strconv.Atoi(values[1])
		h, err := strconv.Atoi(values[2])

		if err != nil {
			continue
		}

		result := ((2 * l * w) + (2 * w * h) + (2 * h * l)) + min(l*w, l*h, w*h)

		total += result
	}

	fmt.Println(total)
}
