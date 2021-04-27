// Multiples of 3 and 5
// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

// Find the sum of all the multiples of 3 or 5 below 1000.
package main

import "fmt"

func getMultiples(x, r int) []int {
	var multiples []int = nil

	for i := 1; i < r; i++ {
		if i%x == 0 {
			multiples = append(multiples, i)
		}
	}
	return multiples
}

func getSum(x []int) int {
	var sum int
	for _, v := range x {
		sum += v
	}
	return sum
}

func getIntersection(x, y []int) []int {
	var intersection []int = nil

	for _, i := range x {
		for _, j := range y {
			if i == j {
				intersection = append(intersection, j)
			}
		}
	}

	return intersection
}

func main() {
	multiples5 := getMultiples(5, 1000)
	multiples3 := getMultiples(3, 1000)
	intersection := getIntersection(multiples5, multiples3)

	sum5 := getSum(multiples5)
	sum3 := getSum(multiples3)
	sumIntersection := getSum(intersection)

	fmt.Printf("sum of all the multiples of 3 or 5 below 1000: %v", sum5+sum3-sumIntersection)
}
