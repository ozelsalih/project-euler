// Sum square difference

// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

package main

import "fmt"

func sqrSum(n int) int {
	return (n * (n + 1) * ((2 * n) + 1)) / 6
}

func sumSqr(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum * sum
}

func sumSqrDifference(x int) int {
	return sumSqr(x) - sqrSum(x)
}

func main() {
	fmt.Printf("difference between the sum of the squares of the first one hundred natural numbers and the square of the sum is %v", sumSqrDifference(100))
}
