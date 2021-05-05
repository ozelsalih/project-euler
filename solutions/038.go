// Pandigital multiples
// What is the largest 1 to 9 pandigital 9-digit number that can be formed as the concatenated product of an integer with (1,2, ... , n) where n > 1?

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isPandigital(num, maxDigitCount int) bool {
	var multiples string
	pandigital := true

	multiples = pandigitalMultiples(num, maxDigitCount)

	pandigital = pandigital && strings.Contains(multiples, "1")
	pandigital = pandigital && strings.Contains(multiples, "2")
	pandigital = pandigital && strings.Contains(multiples, "3")
	pandigital = pandigital && strings.Contains(multiples, "4")
	pandigital = pandigital && strings.Contains(multiples, "5")
	pandigital = pandigital && strings.Contains(multiples, "6")
	pandigital = pandigital && strings.Contains(multiples, "7")
	pandigital = pandigital && strings.Contains(multiples, "8")
	pandigital = pandigital && strings.Contains(multiples, "9")

	return pandigital
}

func pandigitalMultiples(num, maxDigitCount int) string {
	var panDigitalString, tempPanDigitalString string
	var panDigitalInt int
	i := 1
	for i <= 9 {
		panDigitalInt = i * num
		tempPanDigitalString += strconv.Itoa(panDigitalInt)
		if len(tempPanDigitalString) <= maxDigitCount {
			panDigitalString = tempPanDigitalString
		} else {
			i = 9
		}
		i++
	}

	return panDigitalString
}

func main() {
	biggest, to := 0, 9

	for i := 0; i < 10000; i++ {
		if isPandigital(i, to) {
			tempPandigital, _ := strconv.Atoi(pandigitalMultiples(i, to))
			if tempPandigital > biggest {
				biggest = tempPandigital
			}
		}
	}
	fmt.Println("largest 1 to 9 pandigital 9-digit number:", biggest)
}
