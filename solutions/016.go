// Power digit sum
// What is the sum of the digits of the number 2^1000?

package main

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
)

func scientificPow(x, y float64) string {
	input := fmt.Sprintf("%v", math.Pow(x, y))
	flt, _, err := big.ParseFloat(input, 10, 0, big.ToNearestEven)

	if err != nil {
		panic(err)
	}

	var i = new(big.Int)
	i, _ = flt.Int(i)

	return fmt.Sprintf("%v", i)
}

func sumDigits(num string) int {
	var sum int

	for i := 0; i < len(num); i++ {
		digit, _ := strconv.Atoi(string(num[i]))
		sum += digit
	}

	return sum
}

func main() {
	number := scientificPow(2, 1000)
	fmt.Println(sumDigits(number))
}
