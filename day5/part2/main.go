package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/JackFazackerley/aoc/pkg/common"
)

func main() {
	absPath, _ := filepath.Abs("day5/input.txt")

	lines := common.ReadStings(absPath)

	var diagram [1000][1000]int

	for _, line := range lines {
		x1, y1, x2, y2 := parseLines(line)
		if x1 == x2 {
			setPoints(y1, y2, x1, true, &diagram)
		} else if y1 == y2 {
			setPoints(x1, x2, y1, false, &diagram)
		} else {
			setHorizontalPoints(x1, y1, x2, y2, &diagram)
		}
	}

	overlapCount := 0

	for _, x := range diagram {
		for _, y := range x {
			if y == 0 {
				fmt.Printf(".")
			} else {
				fmt.Printf("%d", y)
			}
			if y > 1 {
				overlapCount++
			}
		}
		fmt.Println()
	}

	fmt.Println(overlapCount)
}

func setHorizontalPoints(x1, y1, x2, y2 int, diagram *[1000][1000]int) {
	x := x1
	y := y1
	for {
		diagram[y][x]++
		if x == x2 && y == y2 {
			break
		}
		if x >= x2 {
			x--
		} else if x <= x2 {
			x++
		}

		if y >= y2 {
			y--
		} else if y <= y2 {
			y++
		}

	}
}

func setPoints(x1, x2, y int, vertical bool, diagram *[1000][1000]int) {
	smallest := x1
	biggest := x2
	if x2 < x1 {
		smallest = x2
		biggest = x1
	}

	for i := smallest; i <= biggest; i++ {
		if vertical {
			diagram[i][y] += 1
		} else {
			diagram[y][i] += 1
		}
	}
}

func parseLines(line string) (int, int, int, int) {
	startFinish := strings.Split(line, " -> ")

	xy1 := strings.Split(startFinish[0], ",")
	xy2 := strings.Split(startFinish[1], ",")

	x1, _ := strconv.Atoi(xy1[0])
	y1, _ := strconv.Atoi(xy1[1])
	x2, _ := strconv.Atoi(xy2[0])
	y2, _ := strconv.Atoi(xy2[1])

	return x1, y1, x2, y2
}
