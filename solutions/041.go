// Pandigital prime
// What is the largest n-digit pandigital prime that exists?

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isPandigital(num int) bool {
	numAsString := strconv.Itoa(num)
	digitCount := len(numAsString)
	pandigital := true

	if digitCount < 10 {
		var digit string
		for i := 1; i <= digitCount; i++ {
			digit = strconv.Itoa(i)
			if strings.Count(numAsString, digit) != 1 {
				pandigital = false
			}
		}
	} else {
		pandigital = false
	}

	return pandigital
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

func main() {
	var largest int

	//Because of the division by three rule, the maximum digit must be 7
	for i := 0; i < 10000000; i++ {
		if isPandigital(i) {
			if isPrime(i) && i > largest {
				largest = i
			}
		}
	}
	fmt.Println("largest n-digit pandigital prime is:", largest)
}
