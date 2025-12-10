package main

import "fmt"

func part1() {
	grid := readGrid("input.txt")

	dirs := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1}, // left-up, up, right-up
		{0, -1}, {0, 1}, // left, right
		{1, -1}, {1, 0}, {1, 1}, // left-down, down, right-down
	}

	ans := 0
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			c := grid[i][j]

			if c == '.' {
				continue
			}

			count := 0
			for _, dir := range dirs {
				nr, nc := i+dir[0], j+dir[1]
				if 0 <= nr && nr < m && 0 <= nc && nc < n {
					if grid[nr][nc] == '@' {
						count++
					}
				}
			}

			if count < 4 {
				ans++
			}
		}
	}

	fmt.Println(ans)
}
