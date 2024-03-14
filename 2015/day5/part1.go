package main

import (
	"fmt"
	"os"
	"regexp"
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
		if validRow(line) && validLetters(line) && validVowels(line) {
			niceCount += 1
		}
	}

	fmt.Println(niceCount)
}

func validVowels(line string) bool {
	reg := regexp.MustCompile("(?i)[aeiou]")

	if len(reg.FindAllString(line, -1)) < 3 {
		return false
	}
	return true
}

func validRow(line string) bool {
	for i := 0; i < len(line)-1; i++ {
		if line[i] == line[i+1] {
			return true
		}
	}
	return false
}

func validLetters(line string) bool {
	reg := regexp.MustCompile(`(ab|cd|pq|xy)`)

	if reg.MatchString(line) {
		return false
	}
	return true
}
