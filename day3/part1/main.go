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

	var gammaRate string
	var epsilonRate string

	for x := 0; x < len(lines[0]); x++ {
		oneCount := 0
		zeroCount := 0
		for y := 0; y < len(lines)-1; y++ {
			line := string(lines[y][x])
			if line == "0" {
				zeroCount += 1
			} else {
				oneCount += 1
			}
		}
		if zeroCount > oneCount {
			gammaRate += "0"
			epsilonRate += "1"
		} else {
			gammaRate += "1"
			epsilonRate += "0"
		}
	}

	gamma, _ := strconv.ParseInt(gammaRate, 2, 64)
	epsilon, _ := strconv.ParseInt(epsilonRate, 2, 64)

	fmt.Println(gamma * epsilon)
}
