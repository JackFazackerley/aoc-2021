package main

import (
	"fmt"
	"math"
	"path/filepath"

	"github.com/JackFazackerley/aoc/pkg/common"
)

func main() {
	absPath, _ := filepath.Abs("day7/input.txt")

	crabs := common.ReadIntsWithSep(absPath, ",")

	fuel := math.MaxInt

	for i := 0; i < len(crabs); i++ {
		currentFuel := 0
		for _, crab := range crabs {
			pos := abs(crab - i)
			currentFuel += pos
		}
		if currentFuel < fuel {
			fuel = currentFuel
		}
	}

	fmt.Println(fuel)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
