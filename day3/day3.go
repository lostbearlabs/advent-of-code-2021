package main

import (
	"aoc21/lib"
	"fmt"
	"os"
)

// https://adventofcode.com/2021/day/3

// Usage:
//   go run . < input

func main() {
	lines := lib.ReadLines(os.Stdin)
	bits := len(lines[0])
	nums := lib.LinesToNumbers(lines, 2)
	gamma, epsilon := CalcGammaEpsilon(nums, bits)

	fmt.Printf("gamma: %d\n", gamma)
	fmt.Printf("epsilon: %d\n", epsilon)
	fmt.Printf("gamma * epsilon: %d\n", gamma*epsilon)

	o2, co2 := ComputeO2AndCO2(nums, bits)

	fmt.Printf("o2: %d\n", o2)
	fmt.Printf("co2: %d\n", co2)
	fmt.Printf("o2 * co2: %d\n", o2*co2)

}

func ComputeO2AndCO2(nums []int, bits int) (int, int) {
	return o2(nums, bits, true), o2(nums, bits, false)
}

func o2(nums []int, bits int, useGamma bool) int {

	nextNums := nums
	for i := 0; i < bits; i++ {
		nums = nextNums
		nextNums = make([]int, 0)

		// which bit are we looking at this time?
		mask := 1 << (bits - i - 1)

		// mask gamma (or epsilon) with this mask
		var compareBit int
		gamma, epsilon := CalcGammaEpsilon(nums, bits)
		if useGamma {
			compareBit = gamma & mask
		} else {
			compareBit = epsilon & mask
		}

		// get all the remaining numbers that have the right value for compareBit
		for _, num := range nums {
			if num&mask == compareBit {
				nextNums = append(nextNums, num)
			}
		}

		// are we down to one?
		if len(nextNums) == 1 {
			return nextNums[0]
		}
	}

	fmt.Errorf("Whoops, did not converge")
	return 0
}

func co2(nums []int, bits int) int {

	nextNums := nums
	for i := 0; i < bits; i++ {
		nums = nextNums
		nextNums = make([]int, 0)

		mask := 1 << (bits - i - 1)
		_, epsilon := CalcGammaEpsilon(nums, bits)
		for _, num := range nums {
			if num&mask == epsilon&mask {
				nextNums = append(nextNums, num)
			}
		}

		if len(nextNums) == 1 {
			return nextNums[0]
		}
	}

	return 0
}

func CalcGammaEpsilon(nums []int, bits int) (int, int) {
	gamma := 0

	zeros := make([]int, bits)
	n := len(nums)

	for _, num := range nums {
		for i := 0; i < bits; i++ {
			if (num & (1 << i)) == 0 {
				zeros[bits-i-1]++
			}
		}
	}

	for i := 0; i < bits; i++ {
		gamma <<= 1
		if zeros[i] <= n/2 {
			gamma |= 1
		}
	}

	epsilon := (1<<bits - 1) & ^gamma
	return gamma, epsilon
}
