package main

import (
	"fmt"
	"strconv"
)

func part1() {
	// read file and parse it
	worksheet := readParse("input.txt")

	m, n := len(worksheet), len(worksheet[0])
	ans := 0
	for j := range n {
		switch worksheet[m-1][j] {
		case "*":
			tempAns := 1
			for i := 0; i < m-1; i++ {
				num, _ := strconv.Atoi(worksheet[i][j])
				tempAns *= num
			}
			ans += tempAns
		case "+":
			tempAns := 0
			for i := 0; i < m-1; i++ {
				num, _ := strconv.Atoi(worksheet[i][j])
				tempAns += num
			}
			ans += tempAns
		}
	}

	fmt.Println(ans)
}
