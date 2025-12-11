package main

import (
	"bufio"
	"log"
	"os"
)

func readGrid(fileName string) [][]byte {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	grid := [][]byte{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []byte(line))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return grid
}
