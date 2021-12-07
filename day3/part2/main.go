package main

import (
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/JackFazackerley/aoc/pkg/common"
)

func main() {
	absPath, _ := filepath.Abs("day3/input.txt")

	lines := common.ReadStings(absPath)

	oxygen := getNumbers(lines, 0, true)
	c02 := getNumbers(lines, 0, false)

	ox, _ := strconv.ParseInt(oxygen, 2, 64)
	c0, _ := strconv.ParseInt(c02, 2, 64)

	fmt.Println(ox)
	fmt.Println(c0)

	fmt.Println(ox * c0)
}

func getMost(line []string, x int) (uint8, uint8) {
	var one, zero int
	for _, line := range line {
		if line[x] == '0' {
			zero++
		} else {
			one++
		}
	}

	if zero > one {
		return '0', '1'
	}

	return '1', '0'
}

func getNumbers(lines []string, x int, mostCommon bool) string {
	if len(lines) == 1 {
		return lines[0]
	}

	most, least := getMost(lines, x)
	comparator := least
	if mostCommon {
		comparator = most
	}

	var filtered []string

	for _, line := range lines {
		if line[x] == comparator {
			filtered = append(filtered, line)
		}
	}

	return getNumbers(filtered, x+1, mostCommon)
}
