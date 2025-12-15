package main

import (
	"bufio"
	"log"
	"os"
)

func read(fileName string) [][]rune {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	grid := [][]rune{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return grid
}