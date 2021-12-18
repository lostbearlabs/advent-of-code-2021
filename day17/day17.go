package main

import (
	"image"
	"log"
	"strconv"
	"strings"
)

// Usage:
//   go run

// https://adventofcode.com/2021/day/17

func main() {
	const input = "target area: x=281..311, y=-74..-54"
	target := ParseInput(input)
	ymax, found := FindHighestHit(target)
	log.Println("ymax ", ymax)
	log.Println("len(found) ", len(found))
}

type TargetArea struct {
	x0 int
	y0 int
	x1 int
	y1 int
}

func ParseInput(input string) TargetArea {
	st := strings.Replace(input, "target area: x=", "", -1)
	st = strings.Replace(st, "y=", "", -1)
	st = strings.Replace(st, "..", ", ", -1)
	ar := strings.Split(st, ", ")

	x0, _ := strconv.Atoi(ar[0])
	x1, _ := strconv.Atoi(ar[1])
	y0, _ := strconv.Atoi(ar[2])
	y1, _ := strconv.Atoi(ar[3])

	if x1 < x0 {
		x0, x1 = x1, x0
	}
	if y1 < y0 {
		y0, y1 = y1, y0
	}

	return TargetArea{
		x0: x0,
		y0: y0,
		x1: x1,
		y1: y1,
	}
}

const (
	LT = iota
	EQ
	GT
)

func ltEqGt(amin int, amax int, actual int) int {
	if amin > amax {
		log.Fatal("amin > amax", amin, amax)
	}
	if actual < amin {
		return LT
	} else if actual > amax {
		return GT
	} else {
		return EQ
	}
}

// TraceTrajectory
// returns highest altitude reached and LT/EQ/GT for X and Y when the trajectory falls below the target area and
// whether the trajectory ever hit the target area
func TraceTrajectory(vx int, vy int, target TargetArea) (int, int, int, bool) {
	x := 0
	y := 0

	max := -1_000_000_000
	hit := false

	for x <= target.x1 && y >= target.y0 {
		if y > max {
			max = y
		}
		if x >= target.x0 && x <= target.x1 && y >= target.y0 && y <= target.y1 {
			hit = true
		}
		x += vx
		y += vy
		if vx > 0 {
			vx -= 1
		} else if vx < 0 {
			vx += 1
		}
		vy -= 1
	}

	return max, ltEqGt(target.x0, target.x1, x), ltEqGt(target.y0, target.y1, y), hit
}

func FindHighestHit(target TargetArea) (int, map[image.Point]bool) {
	ymax := -1_000_000_000
	attempted := make(map[image.Point]bool)
	found := make(map[image.Point]bool)

	ymax = search(target, ymax, 0, 0, attempted, found)
	return ymax, found
}

func search(target TargetArea, ymax int, vx int, vy int, attempted map[image.Point]bool, found map[image.Point]bool) int {
	// exclude silly velocities
	max := 500
	if vx > max || vy > max || vy < -max {
		return ymax
	}

	// to avoid duplicated effort, we memoize the velocities we attempt
	apt := image.Pt(vx, vy)
	_, ok := attempted[apt]
	if ok {
		return ymax
	}
	attempted[apt] = true

	yy, _, _, hit := TraceTrajectory(vx, vy, target)
	//log.Printf("vx=%d, vy=%d, hit=%t, yy=%d, endx=%d, endy=%d\n", vx, vy, hit, yy, endx, endy)

	// if we got a hit, we can update ymax
	if hit {
		log.Printf("found: %d,%d\n", vx, vy)
		found[apt] = true
		if yy > ymax {
			ymax = yy
		}
	}

	// TODO: we could probably be smart here to reduce the seach time, but it runs
	// plenty fast without bothering.
	//	// if we didn't overshoot the target, then we could try faster velocities
	//	if endx != GT || endy != LT {
	ymax = search(target, ymax, vx+1, vy, attempted, found)
	ymax = search(target, ymax, vx, vy+1, attempted, found)
	ymax = search(target, ymax, vx, vy-1, attempted, found)
	//	}

	return ymax
}
