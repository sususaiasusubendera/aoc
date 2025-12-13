package main

import "fmt"

func part2() {
	worksheet := readRaw("input.txt")

	m, n := len(worksheet), len(worksheet[0])
	ans := int64(0)

	j := 0
	for j < n {
		sign := worksheet[m-1][j]

		nextj := j + 1
		for nextj < n && worksheet[m-1][nextj] == ' ' {
			nextj++
		}

		// edge case: last col
		if nextj == n {
			nextj += 1
		}

		var tempAns int64
		switch sign {
		case '+':
			tempAns = 0
		case '*':
			tempAns = 1
		}

		for col := j; col < nextj-1; col++ {
			num := int64(0)
			for row := 0; row < m-1; row++ {
				if worksheet[row][col] != ' ' {
					num = (num * 10) + int64(worksheet[row][col]-'0')
				}
			}

			switch sign {
			case '+':
				tempAns += num
			case '*':
				tempAns *= num
			}
		}

		ans += tempAns
		j = nextj
	}

	fmt.Println(ans)
}
