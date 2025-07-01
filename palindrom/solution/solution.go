package solution

import (
	"strings"
	"unicode/utf8"
)

type palindrom string

func New(s string) Solution {
	return palindrom(s)
}

// O(n) -> O(n); MemmoryAllocated -> O(n)
func (p palindrom) IsPalindrome1() bool {
	if len(p) < 2 {
		return false
	}

	var newS strings.Builder

	for i := len(p); i > 0; {
		r, size := utf8.DecodeLastRuneInString(string(p[:i]))

		newS.WriteString(string(r))
		i -= size
	}

	return string(p) == newS.String()
}

// O(n) -> O(n); MemmoryAllocated -> O(n)
func (p palindrom) IsPalindrome2() bool {
	if len(p) < 2 {
		return false
	}

	runes := []rune(p)

	for i := 0; i < len(runes)/2; i++ {
		if runes[i] != runes[len(runes)-i-1] {
			return false
		}
	}

	return true
}

// O(n) -> O(n); MemmoryAllocated -> O(n)
func (p palindrom) IsPalindrome3() bool {
	if len(p) < 2 {
		return false
	}

	runes := []rune(p)

	return recPalindrome(runes, 0, len(runes)-1)
}

func recPalindrome(runes []rune, left, right int) bool {
	if left >= right {
		return true
	}

	if runes[left] != runes[right] {
		return false
	}

	return recPalindrome(runes, left+1, right-1)
}

// O(n) -> O(n); MemmoryAllocated -> O(1)
func (p palindrom) IsPalindrome4() bool {
	if len(p) < 2 {
		return false
	}

	left := 0
	right := len(p)

	for left < right {
		runeLeft, sizeLeft := utf8.DecodeRuneInString(string(p[left:]))
		runeRight, sizeRight := utf8.DecodeLastRuneInString(string(p[:right]))

		if runeLeft != runeRight {
			return false
		}

		left += sizeLeft
		right -= sizeRight
	}

	return true
}
