package main

import (
	arraysandhashing "github.com/utsabc/neet-code-150/arrays-and-hashing"
)

func main() {
	// array of random numbers
	numbers := []int{1, 2, 3, 1, 2, 4, 5, 6, 7, 8, 9, 10}
	println(arraysandhashing.FindDuplicatesSlow(numbers))
	println(arraysandhashing.FindDuplciatesFast(numbers))
}
