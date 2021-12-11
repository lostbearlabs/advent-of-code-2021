package main

import (
	"aoc21/lib"
	"log"
	"os"
)

// Usage:
//   go run . < input

// https://adventofcode.com/2021/day/11

func main() {
	ar := lib.ReadDigitsArray(os.Stdin)

	count := CountFlashes(ar, 100)
	log.Println("count at 100:", count)

	all := AllFlash(ar)
	log.Println("all flash after an additional: ", all)
}

func AllFlash(ar [][]int) int {
	n := 1
	for {
		Increment(ar)
		Propagate(ar)
		flashes := CountAndReset(ar)
		if flashes == len(ar)*len(ar) {
			return n
		}
		n++
	}
}

func CountFlashes(ar [][]int, numSteps int) int {

	flashes := 0
	for i := 0; i < numSteps; i++ {
		Increment(ar)
		Propagate(ar)
		n := CountAndReset(ar)
		flashes += n
	}
	return flashes

}

func Increment(ar [][]int) {
	for i := 0; i < len(ar); i++ {
		for j := 0; j < len(ar[i]); j++ {
			ar[i][j]++
		}
	}
}

// CountAndReset
// at end of each turn, count cells that have flashed and set them to zero
func CountAndReset(ar [][]int) int {
	count := 0
	for i := 0; i < len(ar); i++ {
		for j := 0; j < len(ar[i]); j++ {
			if ar[i][j] > 9 {
				ar[i][j] = 0
				count++
			}
		}
	}
	return count
}

func Propagate(ar [][]int) {

	for i := 0; i < len(ar); i++ {
		for j := 0; j < len(ar[i]); j++ {
			spreadFrom(ar, i, j, false, " ")
		}
	}
}

func spreadFrom(ar [][]int, i int, j int, incr bool, pad string) {

	// not in the grid
	if i < 0 || j < 0 || i >= len(ar) || j >= len(ar[i]) {
		return
	}

	//log.Println(pad, "visit i=", i, ", j=", j, ", val", ar[i][j])

	if ar[i][j] == 999 {
		return // already flashed
	}

	if incr {
		//log.Println(pad, "incr i=", i, ", j=", j)
		ar[i][j]++
	}

	if ar[i][j] <= 9 {
		return // not ready to flash
	}

	//log.Println(pad, "flash! i=", i, ", j=", j)
	ar[i][j] = 999 // flash now

	for di := -1; di <= 1; di++ {
		for dj := -1; dj <= 1; dj++ {
			if di != 0 || dj != 0 {
				spreadFrom(ar, i+di, j+dj, true, pad+"  ")
			}
		}
	}
}
