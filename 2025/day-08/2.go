package main

import (
	"fmt"
	"sort"
)

// use Pair from 1.go
// use DSU from 1.go

func part2() {
	junctionBoxes := read("input.txt")

	pairs := []Pair{}
	for i := 0; i < len(junctionBoxes)-1; i++ {
		for j := i + 1; j < len(junctionBoxes); j++ {
			ax, ay, az := junctionBoxes[i][0], junctionBoxes[i][1], junctionBoxes[i][2]
			bx, by, bz := junctionBoxes[j][0], junctionBoxes[j][1], junctionBoxes[j][2]
			distance := square(ax-bx) + square(ay-by) + square(az-bz)

			p := Pair{i, j, distance}
			pairs = append(pairs, p)
		}
	}

	sort.Slice(pairs, func(i, j int) bool { return pairs[i].d < pairs[j].d })

	dsu := NewDSU(len(junctionBoxes))
	countCircuit := len(junctionBoxes)
	for _, p := range pairs {
		a, b := p.a, p.b

		ok := dsu.Union(a, b)
		if ok {
			countCircuit--
		}

		if countCircuit == 1 {
			xa, xb := junctionBoxes[a][0], junctionBoxes[b][0]
			fmt.Println(xa * xb)
			break
		}
	}
}
