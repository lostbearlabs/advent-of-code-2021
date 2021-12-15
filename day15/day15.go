package main

import (
	"aoc21/lib"
	"image"
	"log"
	"os"
)

// Usage:
//   go run . < input

// https://adventofcode.com/2021/day/15

const max int = 1_000_000_000

func main() {
	grid := lib.ReadDigitsArray(os.Stdin)
	cost := ShortestCornerToCornerCost(grid)
	log.Println("Cost: ", cost)

	expanded := Expand(grid)
	cost = ShortestCornerToCornerCost(expanded)
	log.Println("Expanded: ", cost)
}

func ShortestCornerToCornerCost(grid [][]int) int {
	// initial path is to starting point with cost zero
	pt := image.Pt(0, 0)
	pathSet := PathSet{
		paths: map[image.Point]Path{pt: NewPath(pt, 0)},
		todo:  map[image.Point]Path{pt: NewPath(pt, 0)},
	}

	// we're trying to get to this point
	tgt := image.Pt(len(grid)-1, len(grid[0])-1)

	// here we go ...
	return search(grid, tgt, pathSet)
}

// This is just Dijkstra ...
func search(grid [][]int, tgt image.Point, pathSet PathSet) int {

	for {
		// have we reached our target?  if so, done!
		probe, ok := pathSet.paths[tgt]
		if ok {
			log.Println("found ", probe)
			return probe.cost
		}

		deltas := []image.Point{
			{X: 1, Y: 0},
			{X: 0, Y: 1},
			{X: -1, Y: 0},
			{X: 0, Y: -1},
		}

		path := cheapestTodo(&pathSet)
		//log.Println("next: ", path)

		for _, delta := range deltas {
			addIfPossible(grid, path, delta, &pathSet)
		}

		delete(pathSet.todo, path.pt)
	}
}

func addIfPossible(grid [][]int, path Path, delta image.Point, pathSet *PathSet) {

	x := path.pt.X + delta.X
	y := path.pt.Y + delta.Y
	qt := image.Pt(x, y)

	// illegal point
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) {
		return
	}

	// already processed
	_, ok := pathSet.paths[qt]
	if ok {
		return
	}

	newPath := NewPath(qt, path.cost+grid[x][y])
	//log.Println("add ", newPath)
	pathSet.paths[newPath.pt] = newPath
	pathSet.todo[newPath.pt] = newPath
}

func cheapestTodo(pathSet *PathSet) Path {
	var best Path
	first := true
	for _, path := range pathSet.todo {
		if first || path.cost < best.cost {
			best = path
			first = false
		}
	}
	if first {
		log.Fatal("never found cheapest")
	}

	return best
}

// *****************
// Path/Pathset - for keeping track of the frontier
// *****************

type PathSet struct {
	paths map[image.Point]Path
	todo  map[image.Point]Path
}

type Path struct {
	pt   image.Point
	cost int
}

func NewPath(pt image.Point, cost int) Path {
	return Path{pt: pt, cost: cost}
}

func Expand(grid [][]int) [][]int {
	n := 5
	d := make([][]int, n*len(grid))
	for i := 0; i < len(d); i++ {
		d[i] = make([]int, n*len(grid))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			for di := 0; di < n; di++ {
				for dj := 0; dj < n; dj++ {
					ii := i + di*len(grid)
					jj := j + dj*(len(grid[0]))
					d[ii][jj] = (grid[i][j]-1+di+dj)%9 + 1
				}
			}
		}
	}

	return d
}
