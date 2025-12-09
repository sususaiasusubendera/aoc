package main

import (
	"log"
	"strconv"
)

func parseLine(line string) (byte, int) {
	dir := line[0]
	num, err := strconv.Atoi(line[1:])
	if err != nil {
		log.Fatal(err)
	}

	return dir, num
}
