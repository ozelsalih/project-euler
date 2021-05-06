// Pentagon numbers
// Find the pair of pentagonal numbers, Pj and Pk, for which their sum and difference are pentagonal and D = |Pk − Pj| is minimised; what is the value of D?

package main

import (
	"fmt"
	"math"
)

func pentagonalNum(num int) int {
	return (num * ((3 * num) - 1)) / 2
}

func isPentagonalNum(num int) bool {
	pentagonal := (math.Sqrt(float64(1.0+24.0*float64(num))) + 1.0) / 6.0
	return pentagonal == float64(int(pentagonal))
}

func isPentagonalPair(j, k int) bool {
	pentagonal := false
	pentagonalJ := pentagonalNum(j)
	pentagonalK := pentagonalNum(k)
	sum := pentagonalJ + pentagonalK
	dif := pentagonalJ - pentagonalK

	if dif < 0 {
		dif = -dif
	}

	if isPentagonalNum(sum) && isPentagonalNum(dif) {
		pentagonal = true
	}

	return pentagonal
}

func main() {
	for i := 1; i < 10000; i++ {
		for j := 1; j < i; j++ {
			if isPentagonalPair(i, j) {
				fmt.Printf("D = |P%v P%v|,  D = %v", i, j, pentagonalNum(i)-pentagonalNum(j))
			}
		}
	}
}
