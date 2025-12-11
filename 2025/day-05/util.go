package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func read(fileName string) ([]string, []string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	block1 := []string{}
	block2 := []string{}
	curr := &block1

	for scanner.Scan() {
		line := scanner.Text()

		if len(strings.TrimSpace(line)) == 0 {
			curr = &block2
			continue
		}

		*curr = append(*curr, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return block1, block2
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
