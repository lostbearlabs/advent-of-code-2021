package main

import (
	"aoc21/lib"
	"fmt"
	"image"
	"log"
	"os"
	"sort"
)

// Usage:
//   go run . < input

// https://adventofcode.com/2021/day/9

func main() {
	ar := lib.ReadDigitsArray(os.Stdin)
	riskLevel := ComputeRiskLevel(ar)
	log.Println("riskLevel: ", riskLevel)

	basinProduct := ComputeBasinProduct(ar)
	log.Println("basinProduct: ", basinProduct)
}

func isLowPoint(ar [][]int, i int, j int, ni int, nj int) bool {
	x := ar[i][j]
	if i > 0 && x >= ar[i-1][j] {
		return false
	}

	if i < ni-1 && x >= ar[i+1][j] {
		return false
	}

	if j > 0 && x >= ar[i][j-1] {
		return false
	}

	if j < nj-1 && x >= ar[i][j+1] {
		return false
	}

	fmt.Println("hit: i=", i, ", j=", j, ", x=", x)
	return true
}

func ComputeRiskLevel(ar [][]int) int {
	log.Println(ar)

	ni := len(ar)
	nj := len(ar[0])
	risk := 0

	log.Println("ni: ", ni)
	log.Println("nj: ", nj)

	for i := 0; i < ni; i++ {
		for j := 0; j < nj; j++ {
			if isLowPoint(ar, i, j, ni, nj) {
				risk += ar[i][j] + 1
			}
		}
	}

	return risk
}

type Basins struct {
	// source array
	Grid [][]int

	// grid dimensions
	NI int
	NJ int

	// cells that have been visited
	Visited map[image.Point]bool

	// initial basin point -> reachable points
	Roots map[image.Point]int
}

func dxy(p image.Point, dx int, dy int) image.Point {
	return image.Point{
		X: p.X + dx,
		Y: p.Y + dy,
	}
}

func addPoint(b *Basins, root image.Point, pt image.Point) {
	if pt.X < 0 || pt.Y < 0 || pt.X >= b.NI || pt.Y >= b.NJ {
		return // point out of bounds
	}
	if b.Visited[pt] {
		return // point already processed
	}
	if b.Grid[pt.X][pt.Y] == 9 {
		return // point not in any basin
	}

	// update root for this crawl
	b.Visited[pt] = true
	b.Roots[root] += 1

	// crawl in all directions
	addPoint(b, root, dxy(pt, 1, 0))
	addPoint(b, root, dxy(pt, -1, 0))
	addPoint(b, root, dxy(pt, 0, 1))
	addPoint(b, root, dxy(pt, 0, -1))
}

func makeBasins(ar [][]int) Basins {
	basins := Basins{
		Grid:    ar,
		NI:      len(ar),
		NJ:      len(ar[0]),
		Visited: map[image.Point]bool{},
		Roots:   map[image.Point]int{},
	}

	for i := 0; i < basins.NI; i++ {
		for j := 0; j < basins.NJ; j++ {
			pt := image.Point{
				X: i,
				Y: j,
			}
			addPoint(&basins, pt, pt)
		}
	}

	return basins
}

func ComputeBasinProduct(ar [][]int) int {
	basins := makeBasins(ar)
	log.Println(basins)

	var sizes []int
	for _, n := range basins.Roots {
		sizes = append(sizes, n)
	}

	sort.Ints(sizes)
	n := len(sizes)
	log.Println("sizes: ", sizes)
	return sizes[n-1] * sizes[n-2] * sizes[n-3]
}
