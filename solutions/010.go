// Summation of primes
// Find the sum of all the primes below two million.

// It needs to be improved. The result can be reached using a primitive algorithm.

package main

import "fmt"

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

func main() {
	var sum int
	for i := 0; i < 2000000; i++ {
		if isPrime(i) {
			sum += i
			fmt.Println(i)
		}
	}
	fmt.Println(sum)
}
