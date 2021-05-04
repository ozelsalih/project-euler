// Circular primes
// How many circular primes are there below one million?
// 83.285 seconds

// Since the sieve of eratosthenes algorithm is not used, finding prime numbers decreases program performance.
// Finding all rotations of the digits takes a long time because the wrong algorithm is used.
// Simply putting numbers with 1,3,7,9 in their digits into the algorithm could speed up the performance by 60 percent.

package main

import (
	"fmt"
	"strconv"
)

func numToDigits(num int) []string {
	var digits []string = nil
	var digit string

	str := strconv.Itoa(num)

	for _, v := range str {
		digit = fmt.Sprintf("%v", string(v))
		digits = append(digits, digit)
	}

	return digits
}

func rotateDigits(num int) []int {
	var nums []int = nil
	var digits = numToDigits(num)
	var tempNum int
	var strNum string

	for i := 0; i < len(digits); i++ {

		for j := 0; j < len(digits); j++ {
			strNum += digits[(i+j)%len(digits)]
		}

		tempNum, _ = strconv.Atoi(strNum)
		strNum = ""
		nums = append(nums, tempNum)
	}

	return nums
}

func isPrime(value int) bool {
	if value == 2 {
		return true
	} else if value%2 != 0 {
		for i := 2; i <= value/2; i++ {
			if value%i == 0 {
				return false
			}
		}
		return value > 1
	}
	return false
}

func isCircularPrime(num int) bool {
	rotations := rotateDigits(num)
	circularPrime := true

	for _, i := range rotations {
		circularPrime = circularPrime && isPrime(i)
	}

	return circularPrime
}

func main() {
	var count int

	for i := 0; i < 1000000; i++ {
		if isCircularPrime(i) {
			count++
		}
	}

	fmt.Println(count)
}
