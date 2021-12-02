package main

import (
	"fmt"
	"path/filepath"

	"github.com/JackFazackerley/aoc/pkg/common"
)

func main() {
	absPath, _ := filepath.Abs("day1/input.txt")

	lines := common.ReadInts(absPath)

	count := 0

	for i := 0; i < len(lines)-3; i++ {
		line1 := lines[i]+ lines[i+1] + lines[i+2]
		line2 := lines[i+1] + lines[i+2] + lines[i+3]

		if line2 > line1 {
			count++
		}
	}

	fmt.Println(count)
}
