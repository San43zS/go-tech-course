package solution

import (
	"testing"
)

var (
	strShort = "22\\2\\22"
	strLong  = "фвфвф"
)

func BenchmarkIsPalindrome1Short(b *testing.B) {
	p := New(strShort)
	for i := 0; i < b.N; i++ {
		_ = p.IsPalindrome1()
	}
}

func BenchmarkIsPalindrome1Long(b *testing.B) {
	p := New(strLong)
	for i := 0; i < b.N; i++ {
		_ = p.IsPalindrome1()
	}
}

func BenchmarkIsPalindrome2Short(b *testing.B) {
	p := New(strShort)
	for i := 0; i < b.N; i++ {
		_ = p.IsPalindrome2()
	}
}

func BenchmarkIsPalindrome2Long(b *testing.B) {
	p := New(strLong)
	for i := 0; i < b.N; i++ {
		_ = p.IsPalindrome2()
	}
}

func BenchmarkIsPalindrome3Short(b *testing.B) {
	p := New(strShort)
	for i := 0; i < b.N; i++ {
		_ = p.IsPalindrome3()
	}
}

func BenchmarkIsPalindrome3Long(b *testing.B) {
	p := New(strLong)
	for i := 0; i < b.N; i++ {
		_ = p.IsPalindrome3()
	}
}

func BenchmarkIsPalindrome4Short(b *testing.B) {
	p := New(strShort)
	for i := 0; i < b.N; i++ {
		_ = p.IsPalindrome4()
	}
}

func BenchmarkIsPalindrome4Long(b *testing.B) {
	p := New(strLong)
	for i := 0; i < b.N; i++ {
		_ = p.IsPalindrome4()
	}
}
