package benchmark

import "testing"

var s = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaababaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

func BenchmarkTest1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome1(s)
	}
}
func BenchmarkTest2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome2(s)
	}
}
func BenchmarkTest3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome3(s)
	}
}
func BenchmarkTest4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome4(s)
	}
}
