// Largest prime factor
// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ?

package main

import "fmt"

func smallestDivisor(x int) int {
	if x%2 == 0 {
		return 2
	} else {
		i := 3
		for i*i <= x {
			if x%i == 0 {
				return i
			}
			i += 2
		}
		return x
	}
}

func findPrimeFactor(x int) []int {
	var primeFactors []int = nil
	for x != 1 {
		smallest := smallestDivisor(x)
		primeFactors = append(primeFactors, smallest)
		x = x / smallest
	}
	return primeFactors
}

func findMax(a []int) int {
	max := a[0]
	for _, value := range a {
		if value > max {
			max = value
		}
	}
	return max
}

func main() {
	numberToCalculate := 600851475143

	primeFactors := findPrimeFactor(numberToCalculate)
	largestPrimeFactor := findMax(primeFactors)

	fmt.Printf("largest prime factor of the number 600851475143 is, %v", largestPrimeFactor)
}
