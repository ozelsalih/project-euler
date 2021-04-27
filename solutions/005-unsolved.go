package main

import "fmt"

func IsPrime(value int) bool {
	for i := 2; i <= value/2; i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func makePrimeRange(min, max int) []int {
	var a []int = nil
	for i := min; i <= max; i++ {
		if IsPrime(i) {
			a = append(a, i)
		}
	}
	return a
}

func main() {
	sum := 1
	for _, v := range makePrimeRange(1, 10) {
		sum *= v
	}
	fmt.Println(sum)
}
