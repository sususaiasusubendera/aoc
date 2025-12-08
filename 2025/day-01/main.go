package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	dial := 50 // the dial starts by pointing at 50
	countZero := 0
	for scanner.Scan() {
		line := scanner.Text()

		dir, num := parseLine(line)
		switch dir {
		case 'R':
			dial = (dial + num) % 100
		case 'L':
			dial = (dial - num + 100) % 100
		}

		if dial == 0 {
			countZero++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(countZero)
}

func parseLine(line string) (byte, int) {
	dir := line[0]
	num, err := strconv.Atoi(line[1:])
	if err != nil {
		log.Fatal(err)
	}

	return dir, num
}
