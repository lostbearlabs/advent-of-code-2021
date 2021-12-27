package main

import (
	"aoc21/lib"
	"fmt"
	"io"
	"log"
	"os"
)

// Usage:
//   go run . < input

// https://adventofcode.com/2021/day/11

func main() {
	grid := ReadGrid(os.Stdin)
	n := grid.NumSteps()
	log.Println("steps to settle: ", n)
}

// ****************
// Grid
// ****************

type Grid struct {
	nx int
	ny int
	ar []int
}

func NewGrid(nx, ny int) Grid {
	g := Grid{
		nx: nx,
		ny: ny,
		ar: make([]int, nx*ny),
	}
	return g
}

func (g Grid) idx(x, y int) int {
	// to allow for indexes of -1 ... -n
	x = (x + g.nx) % g.nx
	y = (y + g.ny) % g.ny
	return y*g.nx + x
}

func (g Grid) get(x, y int) int {
	return g.ar[g.idx(x, y)]
}

func (g Grid) set(x, y, val int) {
	g.ar[g.idx(x, y)] = val
}

// ****************
// Read
// ****************

const (
	EMPTY = iota
	RIGHT
	DOWN
)

func charToEnum(c byte) int {
	switch c {
	case '.':
		return EMPTY
	case '>':
		return RIGHT
	case 'v':
		return DOWN
	default:
		log.Printf("BAD CHARACTER %c\n", c)
		os.Exit(1)
		return EMPTY
	}
}

func ReadGrid(rd io.Reader) Grid {
	lines := lib.ReadLines(rd)
	grid := NewGrid(len(lines[0]), len(lines))
	for y, line := range lines {
		for x, c := range line {
			grid.set(x, y, charToEnum(byte(c)))
		}
	}

	return grid
}

// ****************
// Step
// ****************

func (g Grid) move2(dx, dy, tgt int) (Grid, bool) {
	moved := false
	next := NewGrid(g.nx, g.ny)
	for x := 0; x < g.nx; x++ {
		for y := 0; y < g.ny; y++ {
			cur := g.get(x, y)
			if cur == tgt {
				if g.get(x+dx, y+dy) == EMPTY {
					next.set(x+dx, y+dy, cur)
					moved = true
				} else {
					next.set(x, y, cur)
				}
			} else if cur != EMPTY {
				next.set(x, y, cur)
			}
		}
	}
	return next, moved
}

func (g Grid) display() {
	fmt.Println()
	for y := 0; y < g.ny; y++ {
		for x := 0; x < g.nx; x++ {
			c := g.get(x, y)
			if c == EMPTY {
				fmt.Print(".")
			} else if c == RIGHT {
				fmt.Print(">")
			} else {
				fmt.Print("v")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (g Grid) step() (Grid, bool) {
	//g.display()
	g1, a := g.move2(1, 0, RIGHT)
	//g1.display()
	g2, b := g1.move2(0, 1, DOWN)
	//g2.display()
	return g2, a || b
}

func (g Grid) NumSteps() int {
	i := 0
	b := true
	for b {
		i++
		if i%1000 == 0 {
			log.Println("STEP ", i)
		}
		g, b = g.step()
	}

	return i
}
