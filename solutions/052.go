// Permuted multiples
// Find the smallest positive integer, x, such that 2x, 3x, 4x, 5x, and 6x, contain the same digits.

package main

import (
	"fmt"
	"strconv"
)

func digitsArray(num int) []int {
	var digits []int = nil
	var digit int

	str := strconv.Itoa(num)

	for _, v := range str {
		digit, _ = strconv.Atoi(fmt.Sprintf("%v", string(v)))

		digits = append(digits, digit)
	}

	return digits
}

func mutiplesArray(num, to int) []int {
	var multiples []int = nil

	for i := 1; i <= to; i++ {
		multiples = append(multiples, i*num)
	}

	return multiples
}

func unorderedEqual(first, second []int) bool {
	if len(first) != len(second) {
		return false
	}
	exists := make(map[int]bool)
	for _, value := range first {
		exists[value] = true
	}
	for _, value := range second {
		if !exists[value] {
			return false
		}
	}
	return true
}

func hasSameDigits(nums []int) bool {
	same := true

	for i := 0; i < len(nums)-1; i++ {
		same = same && unorderedEqual(digitsArray(nums[i]), digitsArray(nums[i+1]))
	}

	return same
}

func main() {
	for i := 1; i < 1000000; i++ {
		if hasSameDigits(mutiplesArray(i, 6)) {
			fmt.Println(i)
		}
	}
}
