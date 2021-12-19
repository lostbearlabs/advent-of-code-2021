package main

import (
	"aoc21/lib"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Usage:
//   go run . < input

// https://adventofcode.com/2021/day/19

func main() {
	input := Parse(os.Stdin)
	pts := FindAllReadingsFromViewpointZero(input)
	log.Println("num points:", len(pts))

	pts = FindAbsolutePositions(input)
	dist := BiggestManhattanDistance(pts)
	log.Println("biggest manhattan distance:", dist)
}

type Point struct {
	x, y, z int
}

func NewPoint(x, y, z int) Point {
	return Point{
		x: x,
		y: y,
		z: z,
	}
}

type Beacon struct {
	readings []Point
}

type Input struct {
	beacons []Beacon
}

func Parse(rd io.Reader) Input {
	input := Input{}
	lines := lib.ReadLines(rd)
	var cur *Beacon

	for _, line := range lines {
		if line == "" {
			continue
		}

		if line[1] == '-' {
			input.beacons = append(input.beacons, Beacon{})
			cur = &input.beacons[len(input.beacons)-1]
			continue
		}

		ar := strings.Split(line, ",")
		x, _ := strconv.Atoi(ar[0])
		y, _ := strconv.Atoi(ar[1])
		z, _ := strconv.Atoi(ar[2])
		cur.readings = append(cur.readings, NewPoint(x, y, z))
	}

	return input
}

type Rotation struct {
	matrix []int
}

func NewRotation(x ...int) Rotation {
	if len(x) != 9 {
		log.Fatalf("bad matrix")
	}

	det := x[0]*(x[4]*x[8]-x[5]*x[7]) +
		x[1]*(x[3]*x[8]-x[5]*x[6]) +
		x[2]*x[3]*x[7] - x[4]*x[6]

	if det != 1 && det != -1 {
		log.Fatal("bad det", det)
	}

	return Rotation{matrix: x}
}

type Offset struct {
	dx, dy, dz int
}

func NewOffset(dx, dy, dz int) Offset {
	return Offset{dx: dx, dy: dy, dz: dz}
}

func ApplyToPoint(point Point, rotation Rotation, offset Offset) Point {
	x := point.x*rotation.matrix[0] + point.y*rotation.matrix[3] + point.z*rotation.matrix[6]
	y := point.x*rotation.matrix[1] + point.y*rotation.matrix[4] + point.z*rotation.matrix[7]
	z := point.x*rotation.matrix[2] + point.y*rotation.matrix[5] + point.z*rotation.matrix[8]
	return Point{
		x: x + offset.dx,
		y: y + offset.dy,
		z: z + offset.dz,
	}
}

func ApplyToBeacon(beacon Beacon, rotation Rotation, offset Offset) Beacon {
	next := Beacon{}
	for _, point := range beacon.readings {
		nextPoint := ApplyToPoint(point, rotation, offset)
		next.readings = append(next.readings, nextPoint)
	}
	return next
}

type Overlap struct {
	offset         Offset
	rotation       Rotation
	beaconReadings []Point
	mappedReadings []Point
}

func FindOverlap(a Beacon, b Beacon, rotation Rotation, offset Offset) Overlap {

	overlap := Overlap{
		offset:         offset,
		rotation:       rotation,
		beaconReadings: nil,
		mappedReadings: nil,
	}

	mp := make(map[Point]bool)
	for _, pa := range a.readings {
		mp[pa] = true
	}

	for _, pb := range b.readings {
		pbMapped := ApplyToPoint(pb, rotation, offset)
		if mp[pbMapped] {
			//log.Println("hit: ", pa, pb, pbMapped)
			overlap.beaconReadings = append(overlap.beaconReadings, pb)
			overlap.mappedReadings = append(overlap.mappedReadings, pbMapped)
		}
	}

	return overlap
}

const (
	X = iota
	Y
	Z
)

func interestingOffsets(ar []Point, br []Point, axis int, dxMatch int, dyMatch int) []int {
	mp := make(map[int]int, 0)
	for _, pa := range ar {
		for _, pb := range br {
			dx := pa.x - pb.x
			dy := pa.y - pb.y
			dz := pa.z - pb.z
			var delta int

			if axis == X {
				delta = dx
			} else if axis == Y {
				delta = dy
			} else {
				delta = dz
			}

			if (axis < Z || dy == dyMatch) && (axis < Y || dx == dxMatch) {
				mp[delta]++
			}
		}
	}

	var zr []int
	for k, v := range mp {
		if v >= 12 {
			zr = append(zr, k)
		}
	}

	sort.Ints(zr)
	//log.Println("   Offsets:", axis, zr)
	//log.Printf("   len(offsets[%d])=%d\n", axis, len(zr))
	return zr
}

func AllRotations() []Rotation {
	var rotations []Rotation
	for i := 0; i <= 2; i++ {
		for j := 0; j <= 2; j++ {
			for k := 0; k <= 2; k++ {
				if i != j && j != k && i != k {
					for mx := -1; mx <= 1; mx += 2 {
						for my := -1; my <= 1; my += 2 {
							for mz := -1; mz <= 1; mz += 2 {
								ar := make([]int, 9)
								ar[i] = mx
								ar[3+j] = my
								ar[6+k] = mz
								rotations = append(rotations, NewRotation(ar...))
							}
						}
					}
				}
			}
		}
	}

	//log.Println("All rotations:", rotations)
	//log.Println("Num rotations: ", len(rotations))
	return rotations
}

func SearchForOverlap(a Beacon, b Beacon) Overlap {

	for _, rotation := range AllRotations() {
		//log.Println("Rotation", rotation)
		bRotated := ApplyToBeacon(b, rotation, NewOffset(0, 0, 0))
		offsetsX := interestingOffsets(a.readings, bRotated.readings, X, 0, 0)
		for _, dx := range offsetsX {
			offsetsY := interestingOffsets(a.readings, bRotated.readings, Y, dx, 0)
			for _, dy := range offsetsY {
				//log.Printf("i=%d/%d, j=%d/%d\n", i, len(offsetsX), j, len(offsetsY))
				offsetsZ := interestingOffsets(a.readings, bRotated.readings, Z, dx, dy)
				for _, dz := range offsetsZ {
					offset := NewOffset(dx, dy, dz)
					overlap := FindOverlap(a, b, rotation, offset)
					if len(overlap.mappedReadings) >= 12 {
						return overlap
					}
				}
			}
		}
	}

	return Overlap{}
}

// MergeBeacons
// merge points from b into a
func MergeBeacons(a Beacon, b Beacon) Beacon {

	mp := make(map[Point]bool)
	for _, pt := range a.readings {
		mp[pt] = true
	}
	for _, pt := range b.readings {
		mp[pt] = true
	}

	c := Beacon{}
	for pt := range mp {
		c.readings = append(c.readings, pt)
	}

	return c
}

// GetOverlapMatrix
// for each pair (i,j) where j>i, returns overlap from beacon i to beacon j
func GetOverlapMatrix(input Input) [][]Overlap {
	ar := make([][]Overlap, len(input.beacons))
	for i, b := range input.beacons {
		ar[i] = make([]Overlap, len(input.beacons))
		for j := 0; j < len(input.beacons); j++ { // a little wasteful, since the matrix is anti-symmetric
			if i != j {
				ar[i][j] = SearchForOverlap(b, input.beacons[j])
			}
		}
	}
	return ar
}

func GetPathFromZero(overlap [][]Overlap, tgt int, accum []int, visited map[int]bool) ([]int, bool) {
	if tgt == 0 {
		return append(accum, 0), true
	}

	if visited[tgt] {
		return nil, false
	}

	visited[tgt] = true
	for i := 0; i < len(overlap); i++ {
		if len(overlap[i][tgt].beaconReadings) > 0 {
			path, ok := GetPathFromZero(overlap, i, append(accum, tgt), visited)
			if ok {
				return path, true
			}
		}
	}
	visited[tgt] = false
	return nil, false
}

func Clone(beacon Beacon) Beacon {
	clone := Beacon{}
	for _, x := range beacon.readings {
		clone.readings = append(clone.readings, x)
	}

	return clone
}

func FindAllReadingsFromViewpointZero(input Input) []Point {
	matrix := GetOverlapMatrix(input)
	merged := Clone(input.beacons[0])

	for i := 1; i < len(matrix); i++ {
		accum := make([]int, 0)
		visited := make(map[int]bool, 0)
		path, ok := GetPathFromZero(matrix, i, accum, visited)
		//log.Println("Path from zero for ", i, path)
		if !ok || path[0] != i || path[len(path)-1] != 0 {
			log.Fatalf("bad path from %d, %v", i, path)
		}

		z := Clone(input.beacons[path[0]])
		for j := 1; j < len(path); j++ {
			overlap := matrix[path[j]][path[j-1]]
			if len(overlap.beaconReadings) == 0 {
				log.Fatalf("no overlap from %d to %d\n", path[j-1], path[j])
			}
			z = ApplyToBeacon(z, overlap.rotation, overlap.offset)
		}

		merged = MergeBeacons(merged, z)
	}

	return merged.readings
}

func FindAbsolutePositions(input Input) []Point {
	matrix := GetOverlapMatrix(input)
	ar := []Point{NewPoint(0, 0, 0)} // position of beacon zero

	for i := 1; i < len(matrix); i++ {
		accum := make([]int, 0)
		visited := make(map[int]bool, 0)
		path, _ := GetPathFromZero(matrix, i, accum, visited)

		z := Beacon{}
		z.readings = append(z.readings, NewPoint(0, 0, 0)) // beacon's position from its own viewpoint

		for j := 1; j < len(path); j++ {
			overlap := matrix[path[j]][path[j-1]]
			z = ApplyToBeacon(z, overlap.rotation, overlap.offset)
		}

		// the beacon's origin has now been translated to beacon zero's origin
		ar = append(ar, z.readings[0])
	}

	return ar
}

func abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

func ManhattanDistance(a Point, b Point) int {
	return abs(a.x, b.x) + abs(a.y, b.y) + abs(a.z, b.z)
}

func BiggestManhattanDistance(ar []Point) int {
	max := 0
	for i := 0; i < len(ar); i++ {
		for j := i + 1; j < len(ar); j++ {
			d := ManhattanDistance(ar[i], ar[j])
			if d >= max {
				max = d
			}
		}
	}
	return max
}

//func FindAllReadingsFromViewpointZero(input Input) []Point {
//
//	beacons := input.beacons
//
//	for iFrom := len(input.beacons) - 1; iFrom > 0; iFrom-- {
//		for iTo := iFrom - 1; iTo >= 0; iTo-- {
//			bFrom := beacons[iFrom]
//			bTo := beacons[iTo]
//			log.Printf("trying %d (%d) -> %d (%d) ", iFrom, len(bFrom.readings), iTo, len(bTo.readings))
//			overlap := SearchForOverlap(bTo, bFrom)
//			if len(overlap.mappedReadings) > 0 {
//				mappedFrom := ApplyToBeacon(bFrom, overlap.rotation, overlap.offset)
//				beacons[iTo] = MergeBeacons(beacons[iTo], mappedFrom)
//				log.Println("merged ", iFrom, "into ", iTo, ", len[iTo] is now", len(beacons[iTo].readings))
//			}
//		}
//	}
//
//	return beacons[0].readings
//}
