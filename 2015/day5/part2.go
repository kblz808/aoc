package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("error opening file")
		return
	}

	niceCount := 0

	for _, line := range strings.Split(string(file), "\n") {
		if validPair(line) && validLetter(line) {
			niceCount += 1
		}
	}

	fmt.Println(niceCount)
}

func validPair(line string) bool {
	for i := 0; i < len(line)-1; i += 1 {
		chunk := string(line[i : i+2])
		if strings.Count(line, chunk) > 1 {
			return true
		}
	}
	return false
}

func validLetter(line string) bool {
	for i := 0; i < len(line)-2; i += 1 {
		chunk := string(line[i : i+3])

		if chunk[0] == chunk[2] {
			return true
		}
	}
	return false
}
