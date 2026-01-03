package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	a, b int   // idx point a and b
	d    int64 // distance between a and b
}

// disjoint set union (dsu) data structure
type DSU struct {
	parent []int
	size   []int
}

func NewDSU(n int) *DSU {
	parent := make([]int, n)
	size := make([]int, n)

	for i := range n {
		parent[i] = i
		size[i] = 1
	}

	return &DSU{
		parent: parent,
		size:   size,
	}
}

func (d *DSU) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x]) // path compression find (optimization)
	}

	return d.parent[x]
}

func (d *DSU) Union(a, b int) {
	ra, rb := d.Find(a), d.Find(b) // root a and b

	if ra == rb {
		// have same root -> already in same circuit
		return
	}

	// union by size (optimization)
	if d.size[ra] < d.size[rb] {
		d.parent[ra] = rb
		d.size[rb] += d.size[ra]
	} else {
		d.parent[rb] = ra
		d.size[ra] += d.size[rb]
	}
}

func part1() {
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

	dsu := NewDSU(1000)
	for i := range 1000 {
		a, b := pairs[i].a, pairs[i].b
		dsu.Union(a, b)
	}

	circuitSize := []int{}
	for i := range dsu.parent {
		// root
		if dsu.parent[i] == i {
			circuitSize = append(circuitSize, dsu.size[i])
		}
	}

	sort.Slice(circuitSize, func(i, j int) bool { return circuitSize[i] > circuitSize[j] })

	ans := 1
	for i := range 3 {
		ans *= circuitSize[i]
	}

	fmt.Println(ans)
}

// note
// spanning tree:
// - a connected graph using all vertices in which there are no circuits/cycles
// - the minimum number of edges is N - 1; e.g. 5 vertices -> spanning tree with 4 edges
