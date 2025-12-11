package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func readParse(fileName string) [][]string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	grid := [][]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Fields(line)
		grid = append(grid, s)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return grid
}
