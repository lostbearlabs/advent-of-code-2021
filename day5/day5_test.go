package main

import (
	"aoc21/lib"
	"fmt"
	"image"
	"strings"
	"testing"
)

func givenTestCoordinates() [][]image.Point {
	var testInput = `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`

	rd := strings.NewReader(testInput)
	return ReadCoordinates(rd)
}

func TestReadCoordinates(t *testing.T) {
	coords := givenTestCoordinates()
	lib.AssertEq(t, "num coords", 10, len(coords))
	lib.AssertEq(t, "first coord", 0, coords[0][0].X)
	lib.AssertEq(t, "last coord", 2, coords[9][1].Y)
}

func TestAddPointToMapHorz(t *testing.T) {
	dx := make(map[image.Point]int)
	AddSegmentToMap(dx, NewPoint(0, 0), NewPoint(2, 0), true)
	lib.AssertEq(t, "map size", 3, len(dx))
	lib.AssertEq(t, "0,0", 1, dx[NewPoint(0, 0)])
	lib.AssertEq(t, "1,0", 1, dx[NewPoint(1, 0)])
	lib.AssertEq(t, "2,0", 1, dx[NewPoint(2, 0)])
}

func TestAddPointToMapVert(t *testing.T) {
	dx := make(map[image.Point]int)
	AddSegmentToMap(dx, NewPoint(0, 0), NewPoint(0, 2), true)
	lib.AssertEq(t, "map size", 3, len(dx))
	lib.AssertEq(t, "0,0", 1, dx[NewPoint(0, 0)])
	lib.AssertEq(t, "0,1", 1, dx[NewPoint(0, 1)])
	lib.AssertEq(t, "0,2", 1, dx[NewPoint(0, 2)])
}

func TestAddPointToMapDiag1(t *testing.T) {
	dx := make(map[image.Point]int)
	AddSegmentToMap(dx, NewPoint(0, 0), NewPoint(2, 2), true)
	lib.AssertEq(t, "map size", 3, len(dx))
	lib.AssertEq(t, "0,0", 1, dx[NewPoint(0, 0)])
	lib.AssertEq(t, "1,1", 1, dx[NewPoint(1, 1)])
	lib.AssertEq(t, "2,2", 1, dx[NewPoint(2, 2)])
}

func TestAddPointToMapDiag2(t *testing.T) {
	dx := make(map[image.Point]int)
	AddSegmentToMap(dx, NewPoint(2, 2), NewPoint(0, 0), true)
	lib.AssertEq(t, "map size", 3, len(dx))
	lib.AssertEq(t, "0,0", 1, dx[NewPoint(0, 0)])
	lib.AssertEq(t, "1,1", 1, dx[NewPoint(1, 1)])
	lib.AssertEq(t, "2,2", 1, dx[NewPoint(2, 2)])
}

func TestAddPointToMapDiag3(t *testing.T) {
	dx := make(map[image.Point]int)
	AddSegmentToMap(dx, NewPoint(0, 2), NewPoint(2, 0), true)
	lib.AssertEq(t, "map size", 3, len(dx))
	lib.AssertEq(t, "0,2", 1, dx[NewPoint(0, 2)])
	lib.AssertEq(t, "1,1", 1, dx[NewPoint(1, 1)])
	lib.AssertEq(t, "2,0", 1, dx[NewPoint(2, 0)])
}

func TestAddPointToMapDiag4(t *testing.T) {
	dx := make(map[image.Point]int)
	AddSegmentToMap(dx, NewPoint(2, 0), NewPoint(0, 2), true)
	lib.AssertEq(t, "map size", 3, len(dx))
	lib.AssertEq(t, "0,2", 1, dx[NewPoint(0, 2)])
	lib.AssertEq(t, "1,1", 1, dx[NewPoint(1, 1)])
	lib.AssertEq(t, "2,0", 1, dx[NewPoint(2, 0)])
}

func TestCoordsToMapNoDiag(t *testing.T) {
	coords := givenTestCoordinates()
	dx := CoordsToMap(coords, false)
	fmt.Println(dx)
	lib.AssertEq(t, "map size", 21, len(dx))
}

func TestNumPointsNoDiag(t *testing.T) {
	coords := givenTestCoordinates()
	dx := CoordsToMap(coords, false)
	n := NumPointsWithCountGT(dx)
	lib.AssertEq(t, "num points", 5, n)
}

func TestNumPointsDiag(t *testing.T) {
	coords := givenTestCoordinates()
	dx := CoordsToMap(coords, true)
	n := NumPointsWithCountGT(dx)
	lib.AssertEq(t, "num points", 12, n)
}
