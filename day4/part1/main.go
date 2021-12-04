package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/JackFazackerley/aoc/pkg/common"
)

func main() {
	absPath, _ := filepath.Abs("day4/input.txt")

	lines := common.ReadStings(absPath)

	draws, cards := parseInput(lines)

	total, last := FindFirst(draws, cards)

	fmt.Println(total * last)
}

func FindFirst(draws []string, cards []*[5][5]string) (int, int) {
	total := 0
	last := 0
	for _, draw := range draws {
		for _, card := range cards {
			if Match(draw, card) {
				last, _ = strconv.Atoi(draw)
				for _, row := range card {
					for _, cell := range row {
						n, _ := strconv.Atoi(cell)
						total += n
					}
				}
				return total, last
			}
		}
	}

	return total, last
}

func Match(draw string, card *[5][5]string) bool {
	for y, row := range card {
		var yCount int
		var xCount int
		for x := range row {
			if card[y][x] == draw || card[y][x] == "X" {
				xCount++
				card[y][x] = "X"
			}
			if card[x][y] == draw || card[x][y] == "X" {
				yCount++
				card[x][y] = "X"
			}

			if yCount == 5 || xCount == 5 {
				return true
			}
		}
	}

	return false
}

func parseInput(lines []string) ([]string, []*[5][5]string) {
	draws := strings.Split(lines[0], ",")

	cards := make([]*[5][5]string, len(lines[1:])/6)

	for i := 1; i < len(lines); i += 6 {
		var grid [5][5]string
		for i, line := range lines[i+1 : i+6] {
			for j, cell := range strings.Fields(line) {
				grid[i][j] = cell
			}
		}
		cards[(i-1)/6] = &grid
	}

	return draws, cards
}
