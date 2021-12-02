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

	for i := 0; i < len(lines)-1; i++ {
		if lines[i+1] > lines[i] {
			count++
		}
	}

	fmt.Println(count)
}
