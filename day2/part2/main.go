package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/JackFazackerley/aoc/pkg/common"
)

func main() {
	absPath, _ := filepath.Abs("day2/input.txt")

	lines := common.ReadStings(absPath)

	depth, aim, zPos := 0, 0, 0

	for _, line := range lines {
		x := strings.Split(line, " ")

		num, _ := strconv.Atoi(x[1])

		switch x[0] {
		case "forward":
			zPos += num
			depth += aim * num
		case "down":
			aim += num
		case "up":
			aim -= num
		}
	}

	fmt.Println(zPos * depth)
}
