// Triangular, pentagonal, and hexagonal
// Find the next triangle number that is also pentagonal and hexagonal.

package main

import "fmt"

func triangleNum(num int) int {
	return (num * (num + 1)) / 2
}

func pentagonalNum(num int) int {
	return (num * ((3 * num) - 1)) / 2
}

func hexagonalNum(num int) int {
	return num * ((2 * num) - 1)
}

func main() {

	for i := 286; i < 100000; i++ {
		for j := 0; j < i; j++ {
			t := triangleNum(i)
			p := pentagonalNum(j)
			if t == p {
				for k := 0; k < j; k++ {
					h := hexagonalNum(k)
					if t == h {
						fmt.Println("triangle number that is also pentagonal and hexagonal is:", h)
					}
				}
			}
		}
	}
}
