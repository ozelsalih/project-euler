// Truncatable primes
// Find the sum of the only eleven primes that are both truncatable from left to right and right to left.

package main

import (
	"fmt"
	"strconv"
)

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

func truncNumber(num int) []int {
	var numbers []int = nil
	numAsString := strconv.Itoa(num)
	var ltr, rtl int

	for i := 1; i <= len(numAsString); i++ {
		ltr, _ = strconv.Atoi(numAsString[:i])
		numbers = append(numbers, ltr)
	}

	for i := 1; i < len(numAsString); i++ {
		rtl, _ = strconv.Atoi(numAsString[i:])
		numbers = append(numbers, rtl)
	}

	return numbers
}

func isTruncatablePrime(num int) bool {
	if num < 10 {
		return false
	} else {
		if isPrime(num) {
			numbers := truncNumber(num)
			truncateblePrime := true

			for _, v := range numbers {
				if truncateblePrime {
					truncateblePrime = truncateblePrime && isPrime(v)
				}
			}

			return truncateblePrime
		} else {
			return false
		}
	}
}

func main() {
	var count, i, sum int

	for count < 11 {
		if isTruncatablePrime(i) {
			count++
			sum += i
		}
		i++
	}

	fmt.Println("sum of the only eleven primes that are both truncatable:", sum)
}
