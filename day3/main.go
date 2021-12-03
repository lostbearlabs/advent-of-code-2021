package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// https://adventofcode.com/2021/day/3

// Usage:
//   go run . < input

func main() {
	nums, bits := LinesToNumbers(os.Stdin)
	gamma, epsilon := CalcGammaEpsilon(nums, bits)

	fmt.Printf("gamma: %d\n", gamma)
	fmt.Printf("epsilon: %d\n", epsilon)
	fmt.Printf("gamma * epsilon: %d\n", gamma*epsilon)

	o2, co2 := ComputeO2AndCO2(nums, bits)

	fmt.Printf("o2: %d\n", o2)
	fmt.Printf("co2: %d\n", co2)
	fmt.Printf("o2 * co2: %d\n", o2*co2)

}

func LinesToNumbers(rd io.Reader) ([]int, int) {
	var nums []int
	var bits int

	scanner := bufio.NewScanner(rd)
	for scanner.Scan() {
		line := scanner.Text()
		bits = len(line)
		num, _ := strconv.ParseInt(line, 2, 64)
		nums = append(nums, int(num))
	}

	return nums, bits
}

func ComputeO2AndCO2(nums []int, bits int) (int, int) {
	return o2(nums, bits), co2(nums, bits)
}

func o2(nums []int, bits int) int {

	nextNums := nums
	for i := 0; i < bits; i++ {
		nums = nextNums
		nextNums = make([]int, 0)

		mask := 1 << (bits - i - 1)
		gamma, _ := CalcGammaEpsilon(nums, bits)
		for _, num := range nums {
			if num&mask == gamma&mask {
				nextNums = append(nextNums, num)
			}
		}

		if len(nextNums) == 1 {
			return nextNums[0]
		}
	}

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
	epsilon := 0

	zeros := make([]int, bits)
	ones := make([]int, bits)

	for _, num := range nums {
		for i := 0; i < bits; i++ {
			if (num & (1 << i)) == 0 {
				zeros[bits-i-1]++
			} else {
				ones[bits-i-1]++
			}
		}
	}

	for i := 0; i < bits; i++ {
		gamma <<= 1
		epsilon <<= 1
		if ones[i] >= zeros[i] {
			gamma |= 1
		} else {
			epsilon |= 1
		}
	}

	return gamma, epsilon
}
