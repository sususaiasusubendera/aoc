package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func part2() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	ans := 0
	pd := parseData(data)
	for _, d := range pd {
		num1, num2 := parseNumber(d)
		for num := num1; num <= num2; num++ {
			stringNum := strconv.Itoa(num)

			if isRepeated(stringNum) {
				ans += num
			}
		}
	}

	fmt.Println(ans)
}
