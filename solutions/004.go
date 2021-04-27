// Largest palindrome product
// A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

// Find the largest palindrome made from the product of two 3-digit numbers.

package main

import (
	"fmt"
	"strconv"
)

func isPalindromic(x int) bool {
	s := strconv.Itoa(x)

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var biggest, num1, num2 int

	for i := 999; i > 900; i-- {
		for j := i; j > 0; j-- {
			num := i * j
			if isPalindromic(num) && num > biggest {
				biggest, num1, num2 = num, i, j
			}
		}
	}

	fmt.Printf("largest palindrome made from the product of two 3-digit numbers is %v * %v = %v", num1, num2, biggest)
}
