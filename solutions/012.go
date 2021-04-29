// Highly divisible triangular number
// What is the value of the first triangle number to have over five hundred divisors?

package main

import "fmt"

func findFactors(number int) []int {
	var factors []int = nil
	i := 1

	for i*i <= number {
		if number%i == 0 {
			factors = append(factors, i)

			if number/i != i {
				factors = append(factors, number/i)
			}
		}
		i++
	}
	return factors
}

func triangleNumber(number int) int {
	var sum int
	for i := 0; i <= number; i++ {
		sum += i
	}
	return sum
}

func findHighestTriangular(divisorCount int) int {
	archived := false
	i := 1
	var highest int

	for archived == false {
		triangle := triangleNumber(i)
		factors := findFactors(triangle)

		if len(factors) >= divisorCount {
			archived = true
			highest = triangle
		}
		i++
	}
	return highest
}

func main() {
	answer := findHighestTriangular(500)
	fmt.Printf("the value of the first triangle number to have over five hundred divisors is: %v", answer)
}
