package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var lights = [1000][1000]bool{}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("error opening file")
		return
	}

	for _, line := range strings.Split(string(file), "\n") {
		if strings.Contains(line, "turn on") {
			changeLights(line, "on")
		}
		if strings.Contains(line, "turn off") {
			changeLights(line, "off")
		}
		if strings.Contains(line, "toggle") {
			changeLights(line, "toggle")
		}
	}

	count := 0

	for y := 0; y < len(lights); y++ {
		for x := 0; x < len(lights[y]); x++ {
			if lights[y][x] {
				count++
			}
		}
	}

	fmt.Println(count)
}

func changeLights(line string, mode string) {
	nums := getNums(line)

	for y := nums[0]; y <= nums[2]; y++ {
		for x := nums[1]; x <= nums[3]; x++ {
			switch mode {
			case "on":
				lights[y][x] = true
			case "off":
				lights[y][x] = false
			case "toggle":
				lights[y][x] = !lights[y][x]
			}
		}
	}
}

func getNums(line string) [4]int {
	reg := regexp.MustCompile(`[0-9]+`)
	vals := reg.FindAllString(line, -1)

	nums := [4]int{}

	for i, val := range vals {
		num, _ := strconv.Atoi(val)
		nums[i] = num
	}

	return nums
}
