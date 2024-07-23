package arraysandhashing

import (
	"math/rand"
	"testing"
)

func getInputForDuplicates() []int {
	var num []int
	for i := 0; i < 10000; i++ {
		// get a random number
		num = append(num, rand.Intn(100))

	}
	return num
}

func BenchmarkFindDuplicatesSlow(b *testing.B) {
	num := getInputForDuplicates()
	for i := 0; i < b.N; i++ {
		FindDuplicatesSlow(num)
	}
}

func BenchmarkFindDuplciatesFast(b *testing.B) {
	num := getInputForDuplicates()
	for i := 0; i < b.N; i++ {
		FindDuplciatesFast(num)
	}
}
