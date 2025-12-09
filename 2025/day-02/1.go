package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func part1() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	ans := 0
	pd := parseData(data)
	for _, d := range pd {
		num1, num2 := parseNumber(d)
		for num := num1; num <= num2; num++ {
			freqDigit := countFreq(num)
			allDigitEven := true
			for _, v := range freqDigit {
				if v%2 == 1 {
					allDigitEven = false
					break
				}
			}

			if allDigitEven {
				stringNum := strconv.Itoa(num)
				mid := len(stringNum) / 2
				left, right := stringNum[:mid], stringNum[mid:]
				if left == right {
					ans += num
				}
			}
		}
	}

	fmt.Println(ans)
}
