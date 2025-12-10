package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	ans := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		s := best12(line)
		num, _ := strconv.Atoi(s)
		ans += num
	}

	fmt.Println(ans)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
