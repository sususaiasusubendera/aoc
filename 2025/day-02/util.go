package main

import (
	"strconv"
	"strings"
)

func parseData(d []byte) []string {
	s := string(d)
	return strings.Split(s, ",")
}

func parseNumber(s string) (int, int) {
	nums := strings.Split(s, "-")

	num1, err := strconv.Atoi(nums[0])
	if err != nil {
		return -1, -1
	}

	num2, err := strconv.Atoi(nums[1])
	if err != nil {
		return -1, -1
	}

	return num1, num2
}

func countFreq(n int) map[int]int {
	m := map[int]int{}
	for n > 0 {
		digit := n % 10
		m[digit]++
		n /= 10
	}

	return m
}
