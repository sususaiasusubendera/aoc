package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func part2() {
	freshIds, _ := read("input.txt")

	rangeIds := [][]int{}
	for _, s := range freshIds {
		parts := strings.Split(s, "-")

		left, _ := strconv.Atoi(parts[0])
		right, _ := strconv.Atoi(parts[1])

		rangeIds = append(rangeIds, []int{left, right})
	}

	// sorting by left range
	sort.Slice(rangeIds, func(i, j int) bool {
		if rangeIds[i][0] == rangeIds[j][0] {
			return rangeIds[i][1] < rangeIds[j][1]
		}

		return rangeIds[i][0] < rangeIds[j][0]
	})

	merged := [][]int{}
	for _, rid := range rangeIds {
		if len(merged) == 0 {
			merged = append(merged, rid)
			continue
		}

		rightMerged := merged[len(merged)-1][1]
		left, right := rid[0], rid[1]

		if left <= rightMerged {
			merged[len(merged)-1][1] = max(merged[len(merged)-1][1], right)
		} else {
			merged = append(merged, rid)
		}
	}

	ans := 0
	for _, m := range merged {
		left, right := m[0], m[1]
		ans += right - left + 1
	}

	fmt.Println(ans)
}
