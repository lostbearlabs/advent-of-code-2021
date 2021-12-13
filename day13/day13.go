package main

import (
	"aoc21/lib"
	"fmt"
	"image"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// Usage:
//   go run . < input

// https://adventofcode.com/2021/day/13

func main() {
	fnord := ReadFnord(os.Stdin)

	mp := fnord.points
	for i, in := range fnord.instructions {
		if i == 0 {
			mx, my := MaxXY(mp)
			log.Println("init count ", len(mp), ", max-x-y ", mx, my)
			log.Println()
		}
		mp = DoFold(mp, in)
		mx, my := MaxXY(mp)
		log.Println("fold ", i, ", count ", len(mp), ", max-x-y ", mx, my)
		log.Println()
	}

	Display(mp)
}

type Instruction struct {
	coord    int
	isCoordX bool
}

type Fnord struct {
	// points on the page
	points map[image.Point]bool

	// list of fold line, true=x or false=y
	instructions []Instruction
}

func NewFnord() Fnord {
	return Fnord{
		points:       make(map[image.Point]bool),
		instructions: nil,
	}
}

func ReadFnord(rd io.Reader) Fnord {
	lines := lib.ReadLines(rd)
	part1 := true
	fnord := NewFnord()

	for _, line := range lines {
		if line == "" {
			part1 = false
		} else if part1 {
			ar := strings.Split(line, ",")
			x, _ := strconv.Atoi(ar[0])
			y, _ := strconv.Atoi(ar[1])
			pt := image.Point{
				X: x,
				Y: y,
			}
			fnord.points[pt] = true
		} else {
			ar := strings.Split(line, "=")
			isX := ar[0] == "fold along x"
			c, _ := strconv.Atoi(ar[1])
			in := Instruction{
				coord:    c,
				isCoordX: isX,
			}
			fnord.instructions = append(fnord.instructions, in)
		}
	}

	return fnord
}

func DoFold(points map[image.Point]bool, in Instruction) map[image.Point]bool {

	if in.isCoordX {
		log.Println("Fold along x=", in.coord)
	} else {
		log.Println("Fold along y=", in.coord)
	}

	after := make(map[image.Point]bool)
	for pt, _ := range points {
		qt := Fold(pt, in)
		after[qt] = true
	}
	return after
}

func Fold(pt image.Point, in Instruction) image.Point {
	if in.isCoordX {
		if pt.X > in.coord {
			return image.Point{
				X: 2*in.coord - pt.X,
				Y: pt.Y,
			}
		} else {
			return pt
		}
	} else {
		if pt.Y > in.coord {
			return image.Point{
				X: pt.X,
				Y: 2*in.coord - pt.Y,
			}
		} else {
			return pt
		}
	}
}

func MaxXY(points map[image.Point]bool) (int, int) {

	maxX := 0
	maxY := 0
	for pt, _ := range points {
		if pt.X > maxX {
			maxX = pt.X
		}
		if pt.Y > maxY {
			maxY = pt.Y
		}
	}

	return maxX, maxY
}

func Display(points map[image.Point]bool) {

	fmt.Println()

	maxX, maxY := MaxXY(points)

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			qt := image.Point{
				X: x,
				Y: y,
			}
			if points[qt] {
				os.Stdout.WriteString("*")
			} else {
				os.Stdout.WriteString(" ")
			}
		}
		os.Stdout.WriteString("\n")
	}

	fmt.Println()
}
