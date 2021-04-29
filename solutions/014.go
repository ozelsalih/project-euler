// Longest Collatz sequence
// Which starting number, under one million, produces the longest chain?

package main

import "fmt"

func collatzSeq(n int) []int {
	var sequence []int = nil
	sequence = append(sequence, n)

	for n != 1 {
		if n%2 == 0 {
			n /= 2
			sequence = append(sequence, n)
		} else {
			n = 3*n + 1
			sequence = append(sequence, n)
		}
	}

	return sequence
}

func findLongestSeq(start int) int {
	var longestChain int
	var longest int

	for i := 1; i < start; i++ {
		seq := len(collatzSeq(i))
		if seq > longestChain {
			longestChain = seq
			longest = i
		}
	}

	return longest
}

func main() {
	longest := findLongestSeq(1000000)
	fmt.Printf("under one million, %v produces the longest chain.", longest)
}
