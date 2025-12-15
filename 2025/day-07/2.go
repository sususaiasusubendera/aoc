package main

import "fmt"

type State struct {
	x, y int
}

func part2() {
	grid := read("input.txt")
	m, n := len(grid), len(grid[0])
	memo := map[State]int{}

	var startIdx int // y
	for j := range n {
		if grid[0][j] == 'S' {
			startIdx = j
			break
		}
	}

	// dp (top-down)
	var dfs func(x, y int) int
	dfs = func(x, y int) int {
		if x < 0 || x >= m || y < 0 || y >= n {
			return 1
		}

		s := State{x, y}
		if v, ok := memo[s]; ok {
			return v
		}

		res := 0
		if grid[x][y] == '^' {
			res += dfs(x, y-1)
			res += dfs(x, y+1)
		} else { // '.'
			res = dfs(x+1, y)
		}

		memo[s] = res
		return res
	}

	ans := dfs(1, startIdx) // x, y

	fmt.Println(ans)
}
