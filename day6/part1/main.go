package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/JackFazackerley/aoc/pkg/common"
)

func main() {
	absPath, _ := filepath.Abs("day6/input.txt")

	lines := common.ReadStings(absPath)

	fish := parseInput(lines[0])

	for day := 0; day < 80; day++ {
		var next [9]int
		next[8] += fish[0]
		next[6] += fish[0]
		for i := 0; i < 8; i++ {
			next[i] += fish[i+1]
		}
		fish = next
	}

	res := 0
	for i := 0; i < 9; i++ {
		res += fish[i]
	}
	fmt.Println(res)
}

func parseInput(input string) [9]int {
	var fishes [9]int

	split := strings.Split(input, ",")

	for _, f := range split {
		day, _ := strconv.Atoi(f)

		fishes[day]++
	}

	return fishes
}
