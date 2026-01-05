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

	redTiles := [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		p := []int{} // coordinate point
		for _, part := range parts {
			num, _ := strconv.Atoi(part)
			p = append(p, num)
		}

		redTiles = append(redTiles, p)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return redTiles
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}