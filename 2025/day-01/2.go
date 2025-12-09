package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func part2() {
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
			for i := 0; i < num; i++ {
				dial = (dial + 1) % 100
				if dial == 0 {
					countZero++
				}
			}
		case 'L':
			for i := 0; i < num; i++ {
				dial = (dial - 1 + 100) % 100
				if dial == 0 {
					countZero++
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(countZero)
}
