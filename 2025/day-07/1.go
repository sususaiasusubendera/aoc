package main

import "fmt"

func part1() {
	grid := read("input.txt")

	m, n := len(grid), len(grid[0])

	var startIdx int
	for j := range n {
		if grid[0][j] == 'S' {
			startIdx = j
			break
		}
	}

	// bfs
	ans := 0
	queue := [][]int{}
	queue = append(queue, []int{1, startIdx})
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:] // pop

		x, y := curr[0], curr[1]

		if 0 <= x && x < m && 0 <= y && y < n && grid[x][y] == '|' {
			continue
		}

		if 0 <= x && x < m && 0 <= y && y < n && grid[x][y] == '^' {
			ans++
			queue = append(queue, []int{x, y - 1})
			queue = append(queue, []int{x, y + 1})
			continue
		}

		if 0 <= x && x < m && 0 <= y && y < n && grid[x][y] == '.' {
			grid[x][y] = '|'
			queue = append(queue, []int{x + 1, y})
		}
	}

	fmt.Println(ans) // ans

	// output (after the operation)
	for i := range m {
		for j := range n {
			if j == n-1 {
				fmt.Println(string(grid[i][j]))
			} else {
				fmt.Print(string(grid[i][j]))
			}
		}
	}
}
