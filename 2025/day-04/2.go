package main

import "fmt"

func part2() {
	grid := readGrid("input.txt")

	dirs := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1}, // left-up, up, right-up
		{0, -1}, {0, 1}, // left, right
		{1, -1}, {1, 0}, {1, 1}, // left-down, down, right-down
	}

	// bfs
	ans := 0
	m, n := len(grid), len(grid[0])
	neigh := make([][]int, m)
	for i := range neigh {
		neigh[i] = make([]int, n)
	}
	queue := [][]int{}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			c := grid[i][j]

			if c == '.' {
				continue
			}

			countNeigh := 0
			for _, dir := range dirs {
				nr, nc := i+dir[0], j+dir[1]
				if 0 <= nr && nr < m && 0 <= nc && nc < n {
					if grid[nr][nc] == '@' {
						countNeigh++
					}
				}
			}

			if countNeigh < 4 {
				queue = append(queue, []int{i, j})
			}

			neigh[i][j] = countNeigh
		}
	}

	for len(queue) > 0 {
		p := queue[len(queue)-1] // point (i, j)

		queue = queue[:len(queue)-1] // pop

		r, c := p[0], p[1]
		grid[r][c] = '.'
		for _, dir := range dirs {
			nr, nc := r+dir[0], c+dir[1]
			if 0 <= nr && nr < m && 0 <= nc && nc < n {
				neigh[nr][nc]--
				if neigh[nr][nc] == 3 {
					queue = append(queue, []int{nr, nc})
				}
			}
		}

		ans++
	}

	fmt.Println(ans)
}
