package main

import (
	"fmt"
	"sliceString/slice"
)

func main() {
	data := []string{"cat", "dog", "cat", "bird", "dog", "fish"}

	s := slice.New(data)
	result := s.Set()

	for _, v := range result {
		fmt.Print(" ", v)
	}
}
