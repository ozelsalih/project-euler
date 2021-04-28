// Special Pythagorean triplet

// There exists exactly one Pythagorean triplet for which a + b + c = 1000.
// Find the product abc.

//Brute Force Approach

package main

import "fmt"

func isPythangoean(a, b, c int) bool {
	if c < b || c < a {
		return false
	} else {
		return a*a+b*b == c*c
	}
}

func share(denominator int) []int {
	var a, b, c int
	var triplet []int = nil
	for i := 2; i < denominator; i++ {
		c = denominator - i
		for j := 1; j < i; j++ {
			b = j
			a = i - j
			if isPythangoean(a, b, c) && a < b {
				triplet = append(triplet, a, b, c)
			}
		}
	}
	return triplet
}

func main() {
	triplet := share(1000)
	product := triplet[0] * triplet[1] * triplet[2]
	fmt.Printf("The pythagorean triple is %v %v %v, product is %v", triplet[0], triplet[1], triplet[2], product)
}
