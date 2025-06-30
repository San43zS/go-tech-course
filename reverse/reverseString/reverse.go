package reverseString

import (
	"fmt"
	"strings"
)

type reverseStr struct {
	s string
}

func New(s string) Reverse {
	return reverseStr{
		s: s,
	}
}

// MemmoryAllocated -> O(n); O(n) -> O(n)
func (r reverseStr) ReverseStandart() string {
	runes := []rune(r.s)
	fmt.Println("runes = ", string(runes))

	var tmp int
	for i := 0; i <= len(runes); i++ {
		if (i == len(runes)) || (runes[i] == ' ') {
			for k, j := tmp, i-1; k < j; k, j = k+1, j-1 {
				runes[k], runes[j] = runes[j], runes[k]
			}
			tmp = i + 1
		}
	}

	return string(runes)
}

// MemmoryAllocated -> O(n); O(n) -> O(n)
func (r reverseStr) ReverseNotStandart() string {
	f := strings.Split(r.s, " ")
	for i := 0; i < len(f); i++ {
		slice := []rune(f[i])

		for j, k := 0, len(slice)-1; j < k; j, k = j+1, k-1 {
			slice[j], slice[k] = slice[k], slice[j]
		}
		f[i] = string(slice)
	}

	return strings.Join(f, " ")
}
