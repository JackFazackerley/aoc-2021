package common

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
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

func ReadStings(path string) []string {
	file, _ := os.Open(path)
	defer file.Close()

	data, _ := io.ReadAll(file)

	return strings.Split(string(data), "\n")
}

func ReadIntsWithSep(path string, sep string) []int {
	file, _ := os.Open(path)
	defer file.Close()

	data, _ := io.ReadAll(file)

	strs := strings.Split(string(data), sep)

	ints := make([]int, len(strs))

	for i, str := range strs {
		ints[i], _ = strconv.Atoi(str)
	}

	return ints
}
