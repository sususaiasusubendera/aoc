package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	ans := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		numMax := 0
		for i := 0; i < len(line); i++ {
			for j := i + 1; j < len(line); j++ {
				digit1, digit2 := int(line[i]-'0'), int(line[j]-'0')
				num := digit1*10 + digit2
				numMax = max(numMax, num)
			}
		}

		ans += numMax
	}

	fmt.Println(ans)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
