package main

import "fmt"

func part1() {
	redTiles := read("input.txt")

	largestRectangle := 0
	for i := 0; i < len(redTiles)-1; i++ {
		for j := i + 1; j < len(redTiles); j++ {
			ax, ay := redTiles[i][0], redTiles[i][1]
			bx, by := redTiles[j][0], redTiles[j][1]

			area := (abs(by-ay) + 1) * (abs(bx-ax) + 1)

			largestRectangle = max(largestRectangle, area)
		}
	}

	fmt.Println(largestRectangle)
}