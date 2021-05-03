// Digit fifth powers
// Find the sum of all the numbers that can be written as the sum of fifth powers of their digits.

package main

import (
	"fmt"
	"math"
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

func isSumOfPow(num, pow int) bool {
	digits := digitsArray(num)
	var sum int

	for _, digit := range digits {
		sum += int(math.Pow(float64(digit), float64(pow)))
	}

	if sum == num {
		return true
	} else {
		return false
	}
}

func findSumOfPow(pow int) []int {
	var numbers []int = nil
	maxPossible := int(math.Pow(9, float64(pow))) * 9
	for i := 2; i <= maxPossible; i++ {
		if isSumOfPow(i, pow) {
			numbers = append(numbers, i)
		}
	}

	return numbers
}

func main() {
	var sum int

	for _, v := range findSumOfPow(5) {
		sum += v
	}

	fmt.Println(sum)
}
