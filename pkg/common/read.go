package common

import (
	"bufio"
	"os"
	"strconv"
)

func ReadInts(path string) []int {
	file, _ := os.Open(path)
	defer file.Close()

	 var v []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(line)
		v = append(v, num)
	}

	return v
}