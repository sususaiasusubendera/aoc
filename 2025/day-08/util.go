package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func read(fileName string) [][]int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	grid := [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		p := []int{} // point
		for _, part := range parts {
			num, _ := strconv.Atoi(part)
			p = append(p, num)
		}

		grid = append(grid, p)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return grid
}

func square(n int) int64 {
	return int64(n) * int64(n)
}
