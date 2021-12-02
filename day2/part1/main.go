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

	xPos, yPos := 0, 0

	for _, line := range lines {
		x := strings.Split(line, " ")

		num, _ := strconv.Atoi(x[1])

		switch x[0] {
		case "forward":
			xPos+= num
		case "down":
			yPos+= num
		case "up":
			yPos-= num

		}
	}

	fmt.Println(xPos * yPos)
}
