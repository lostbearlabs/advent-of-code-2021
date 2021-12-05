package main

import (
	"aoc21/lib"
	"fmt"
	"image"
	"io"
	"os"
	"strings"
)

// Usage:
//   go run . < input

// https://adventofcode.com/2021/day/5

func main() {
	coords := ReadCoordinates(os.Stdin)

	// part 1
	dx := CoordsToMap(coords, false)
	n := NumPointsWithCountGT(dx)
	fmt.Printf("no diag: %d\n", n)

	// part 2
	dx = CoordsToMap(coords, true)
	n = NumPointsWithCountGT(dx)
	fmt.Printf("yes diag: %d\n", n)
}

func ReadCoordinates(read io.Reader) [][]image.Point {
	lines := lib.ReadLines(read)
	var coords [][]image.Point

	for _, line := range lines {
		st := strings.Replace(line, ",", " ", -1)
		st = strings.Replace(st, " -> ", " ", -1)
		ar := strings.Fields(st)
		if len(ar) != 4 {
			fmt.Printf("bad input, len(ar)=%d: %s\n", len(ar), st)
			os.Exit(1)
		}
		br := lib.LinesToNumbers(ar, 10)
		coords = append(coords, []image.Point{
			NewPoint(br[0], br[1]),
			NewPoint(br[2], br[3]),
		})
	}

	return coords
}

func NewPoint(x int, y int) image.Point {
	return image.Point{
		X: x,
		Y: y,
	}
}

func getIncr(d int) int {
	if d > 0 {
		return 1
	} else if d < 0 {
		return -1
	} else {
		return 0
	}
}

func AddSegmentToMap(mp map[image.Point]int, p1 image.Point, p2 image.Point, allowDiagonal bool) {
	x1 := p1.X
	x2 := p2.X
	y1 := p1.Y
	y2 := p2.Y

	dx := x2 - x1
	dy := y2 - y1

	if dx != 0 && dy != 0 && dx != dy && dx != -dy {
		fmt.Printf("illegal: %d,%d -> %d,%d", x1, y1, x2, y2)
		os.Exit(1)
	}

	if dx != 0 && dy != 0 && !allowDiagonal {
		return
	}

	ix := getIncr(dx)
	iy := getIncr(dy)

	x := x1
	y := y1
	pt := NewPoint(x, y)
	mp[pt] += 1

	for x != x2 || y != y2 {
		x += ix
		y += iy

		pt := NewPoint(x, y)
		mp[pt] += 1
	}

}

func CoordsToMap(coords [][]image.Point, allowDiagonal bool) map[image.Point]int {
	dx := make(map[image.Point]int)

	for _, coord := range coords {
		AddSegmentToMap(dx, coord[0], coord[1], allowDiagonal)
	}

	return dx
}

func NumPointsWithCountGT(dx map[image.Point]int) int {
	n := 0
	for _, v := range dx {
		if v > 1 {
			n++
		}
	}
	return n
}
