// 1000-digit Fibonacci number
// What is the index of the first term in the Fibonacci sequence to contain 1000 digits?

package main

import (
	"fmt"
	"math/big"
)

func fib(n int) string {
	a := big.NewInt(1)
	b := big.NewInt(1)
	for i := 3; i <= n+1; i++ {
		a.Add(a, b)
		a, b = b, a
	}
	return a.String()
}

func findFibNumber(n int) int {
	maxLen := 0
	num := 3
	for maxLen <= n {
		maxLen = len(fib(num))
		num++
	}
	return num - 1
}

func main() {
	fmt.Println(findFibNumber(999))
}
