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

	up := 0
	down := 0

	for _, num := range string(file) {
		switch num {
		case '(':
			up += 1
		case ')':
			down += 1
		}
	}

	floor := up - down

	fmt.Println(floor)
}
