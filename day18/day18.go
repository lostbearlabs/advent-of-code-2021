package main

import (
	"aoc21/lib"
	"fmt"
	"log"
	"os"
)

// Usage:
//   go run . < input

// https://adventofcode.com/2021/day/18

func main() {
	lines := lib.ReadLines(os.Stdin)
	ar := make([]*SnailfishNumber, len(lines))
	for i, line := range lines {
		ar[i] = Parse(line)
	}
	num := AddList(ar)
	log.Println("Magnitude of sum:", Magnitude(num))

	log.Println("LargestPairwiseMagnitude:", LargestPairwiseMagnitude(ar))

}

type SnailfishNumber struct {
	leaf  bool
	x     int
	left  *SnailfishNumber
	right *SnailfishNumber
}

func Render(num *SnailfishNumber) string {
	if num.leaf {
		return fmt.Sprintf("%d", num.x)
	} else {
		return fmt.Sprintf("[%s,%s]", Render(num.left), Render(num.right))
	}
}

func Clone(num *SnailfishNumber) *SnailfishNumber {
	if num == nil {
		return nil
	}

	return &SnailfishNumber{
		leaf:  num.leaf,
		x:     num.x,
		left:  Clone(num.left),
		right: Clone(num.right),
	}
}

func Magnitude(num *SnailfishNumber) int {
	if num.leaf {
		return num.x
	}
	return 3*Magnitude(num.left) + 2*Magnitude(num.right)
}

func Add(a *SnailfishNumber, b *SnailfishNumber) *SnailfishNumber {
	pair := SnailfishNumber{
		leaf:  false,
		x:     0,
		left:  Clone(a),
		right: Clone(b),
	}

	Reduce(&pair)
	return &pair
}

func AddList(nums []*SnailfishNumber) *SnailfishNumber {
	num := Add(Clone(nums[0]), nums[1])
	for i := 2; i < len(nums); i++ {
		num = Add(num, nums[i])
	}
	return num
}

func LargestPairwiseMagnitude(nums []*SnailfishNumber) int {
	max := 0

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i != j {
				ii := nums[i]
				jj := nums[j]
				a := Magnitude(Add(ii, jj))
				if a > max {
					max = a
				}
				fmt.Printf("i=%d, j=%d, %s + \n  %s\n    = %d  (%d)\n", i, j, Render(ii), Render(jj), a, max)
			}
		}
	}

	return max
}

// ********************
// Parse
// ********************

func Parse(st string) *SnailfishNumber {
	num, _ := parseFrom(st, 0)
	return num
}

func parseFrom(st string, pos int) (*SnailfishNumber, int) {

	if st[pos] >= '0' && st[pos] <= '9' {
		a, pos := parseDigit(st, pos)
		return &SnailfishNumber{
			leaf:  true,
			x:     a,
			left:  nil,
			right: nil,
		}, pos
	}

	pos = parseExpect(st, pos, '[')
	left, pos := parseFrom(st, pos)
	pos = parseExpect(st, pos, ',')
	right, pos := parseFrom(st, pos)
	pos = parseExpect(st, pos, ']')

	return &SnailfishNumber{
		leaf:  false,
		left:  left,
		right: right,
	}, pos

}

func parseExpect(st string, pos int, expected byte) int {
	if st[pos] != expected {
		log.Fatalf("did not find %c at pos %d in %s\n", expected, pos, st)
	}
	return pos + 1
}

func parseDigit(st string, pos int) (int, int) {
	a := st[pos] - '0'
	if a < 0 || a > 9 {
		log.Fatalf("did not find digit at pos %d in %s\n", pos, st)
	}
	return int(a), pos + 1
}

// ********************
// Reduce
// ********************

func Reduce(num *SnailfishNumber) {
	for Explode(num) || Split(num) {
		// noop
	}
}

func Explode(num *SnailfishNumber) bool {
	pred, tgt, succ := findExplode(num, nil, nil, nil, 0)
	if tgt == nil {
		return false
	}

	if pred != nil {
		pred.x += tgt.left.x
	}
	if succ != nil {
		succ.x += tgt.right.x
	}
	tgt.left = nil
	tgt.right = nil
	tgt.leaf = true
	tgt.x = 0
	return true
}

// depth-first search of the tree, accumulating target pair and its predecessor and successor along the way
func findExplode(num *SnailfishNumber, pred *SnailfishNumber, tgt *SnailfishNumber, succ *SnailfishNumber, depth int) (*SnailfishNumber, *SnailfishNumber, *SnailfishNumber) {

	// tgt is first pair with depth >=4 composed of two normal numbers
	if depth >= 4 && tgt == nil && !num.leaf && num.left.leaf && num.right.leaf {
		return pred, num, succ
	}

	// pred is the first leaf we come to before finding tgt
	if tgt == nil && num.leaf {
		pred = num
	}

	// succ is the first leaf we come to after finding tgt
	if tgt != nil && succ == nil && num.leaf {
		succ = num
	}

	// recur until succ found (or we run out of tree)
	if !num.leaf && succ == nil {
		pred, tgt, succ = findExplode(num.left, pred, tgt, succ, depth+1)
		pred, tgt, succ = findExplode(num.right, pred, tgt, succ, depth+1)
	}

	return pred, tgt, succ
}

func Split(num *SnailfishNumber) bool {
	nd := findSplit(num)
	if nd == nil {
		return false
	}

	nd.left = &SnailfishNumber{
		leaf:  true,
		x:     nd.x / 2,
		left:  nil,
		right: nil,
	}

	nd.right = &SnailfishNumber{
		leaf:  true,
		x:     nd.x/2 + nd.x%2,
		left:  nil,
		right: nil,
	}

	nd.leaf = false
	nd.x = 0

	return true
}

func findSplit(num *SnailfishNumber) *SnailfishNumber {
	if num.leaf {
		if num.x >= 10 {
			return num
		}
		return nil
	}

	nd := findSplit(num.left)
	if nd != nil {
		return nd
	}

	return findSplit(num.right)
}
