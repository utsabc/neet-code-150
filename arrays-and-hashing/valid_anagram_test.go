package arraysandhashing

import "testing"

func getInputForAnagram(n int) string {
	return RandomStringRunes(n)
}

func BenchmarkValidAnagramSlow(b *testing.B) {

	for i := 0; i < b.N; i++ {
		inputString1 := getInputForAnagram(i)
		inputString2 := getInputForAnagram(i)
		ValidAnagramSlow(inputString1, inputString2)
	}
}

func BenchmarkValidAnagramFast(b *testing.B) {

	for i := 0; i < b.N; i++ {
		inputString1 := getInputForAnagram(i)
		inputString2 := getInputForAnagram(i)
		ValidAnagramFast(inputString1, inputString2)
	}
}
