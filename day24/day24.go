package main

import (
	"log"
	"sort"
)

// Usage:
//   go run .

// https://adventofcode.com/2021/day/11

// day 25 ...
// TODO: this is not 100% right yet;  as it is, it gives both correct min and
// correct max, but when I first ran it with d in [1,9] instead of d in [9,1]
// it gave incorrect values, something must be wibbly in my memoization

func main() {
	params := [][]int{
		{1, 12, 4},
		{1, 11, 10},
		{1, 14, 12},
		{26, -6, 14},
		{1, 15, 6},
		{1, 12, 16},
		{26, -9, 1},
		{1, 14, 7},
		{1, 14, 8},
		{26, -5, 11},
		{26, -9, 8},
		{26, -5, 3},
		{26, -2, 1},
		{26, -7, 8},
	}

	log.Println("start...")
	stats := find(params)
	if len(stats.results) > 0 {
		sort.Ints(stats.results)
		log.Println("MIN: ", stats.results[0])
		log.Println("MAX: ", stats.results[len(stats.results)-1])
	} else {
		log.Println("NO RESULTS")
	}

}

type Key struct {
	idx int
	d   int
	z   int
}

func NewKey(idx, d, z int) Key {
	return Key{
		idx: idx,
		d:   d,
		z:   z,
	}
}

type Visited struct {
	canSucceed bool
}

type Stats struct {
	results []int
}

func calcResult(params [][]int, digits []int) int {
	z := 0
	for i := 0; i < len(params); i++ {
		z = compNextZ(params, i, z, digits[i])
	}
	return z
}

func find(params [][]int) Stats {
	stats := Stats{results: make([]int, 0)}
	memo := map[Key]Visited{}
	digits := make([]int, len(params))
	zInit := 0
	idxInit := 0
	dfs(params, &stats, memo, digits, idxInit, zInit)
	return stats
}

// return indicates whether this idx/z can get to a success for any value of d
func dfs(params [][]int, stats *Stats, memo map[Key]Visited, digits []int, idx int, z int) bool {

	if idx >= len(params) {
		if z == 0 {
			log.Println(digits)
			v := calcResult(params, digits)
			if v != 0 {
				log.Fatalf("something went wrong - expected zero, got %d, %v", v, digits)
			}
			n := toInt(digits)
			stats.results = append(stats.results, n)
			return true
		} else {
			v := calcResult(params, digits)
			if v == 0 {
				log.Fatalf("something went wrong - expected non-zero, got %d, %v", v, digits)
			}
			return false
		}
	}

	canSucceed := false
	for d := 9; d >= 1; d-- {
		if idx < 3 {
			log.Printf("... %v\n", digits[:idx])
		}

		key := NewKey(idx, d, z)
		v, ok := memo[key]
		if ok {
			canSucceed = canSucceed || v.canSucceed
		} else {
			digits[idx] = d
			zNext := compNextZ(params, idx, z, d)
			success := dfs(params, stats, memo, digits, idx+1, zNext)
			memo[key] = Visited{canSucceed: success}
			canSucceed = canSucceed || success
		}
	}

	return canSucceed
}

func toInt(digits []int) int {
	n := 0
	for _, d := range digits {
		n = n*10 + d
	}
	return n
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func compNextZ(params [][]int, idx int, zPrev int, d int) int {
	n1 := params[idx][0]
	n2 := params[idx][1]
	n3 := params[idx][2]

	z := zPrev

	//inp w
	w := d
	//mul x 0
	x := 0
	//add x z
	x += z
	//mod x 26
	x %= 26
	//div z (1)
	z /= n1
	//add x (12)
	x += n2
	//eql x w
	x = boolToInt(x == w)
	//eql x 0
	x = boolToInt(x == 0)
	//mul y 0
	y := 0
	//add y 25
	y += 25
	//mul y x
	y *= x
	//add y 1
	y += 1
	//mul z y
	z *= y
	//mul y 0
	y *= 0
	//add y w
	y += w
	//add y (4)
	y += n3
	//mul y x
	y *= x
	//add z y
	z += y

	return z
}
