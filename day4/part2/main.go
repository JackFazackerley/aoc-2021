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

	total, last := FindLast(draws, cards)

	fmt.Println(total)
	fmt.Println(last)

	fmt.Println(total * last)
}

func FindLast(draws []string, cards []*[5][5]string) (int, int) {
	for _, draw := range draws {
		for i := 0; i < len(cards); i++ {
			if Match(draw, cards[i]) {
				total := 0
				last, _ := strconv.Atoi(draw)
				for _, row := range cards[0] {
					for _, cell := range row {
						n, _ := strconv.Atoi(cell)
						total += n
					}
				}

				cards = append(cards[:i], cards[i+1:]...)
				i--

				if len(cards) == 0 {
					return total, last
				}
			}
		}
	}

	return 0, 0
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
