package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func part1() {
	// read input
	freshIds, availIds := read("input.txt")

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

	// optimize by merging the overlapping ranges
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

	// binary search
	ans := 0
	for _, id := range availIds {
		intId, _ := strconv.Atoi(id)

		low, high := 0, len(merged)-1
		for low <= high {
			mid := (low + high) / 2
			left, right := merged[mid][0], merged[mid][1]

			if intId < left {
				high = mid - 1
			} else if intId > right {
				low = mid + 1
			} else {
				ans++
				break
			}
		}
	}

	fmt.Println(ans)
}
