package main

import (
	"fmt"

	"palindrom/solution"
)

func main() {
	str := "фаф"

	sol := solution.New(str)
	fmt.Println(sol.IsPalindrome1())
	fmt.Println(sol.IsPalindrome2())
	fmt.Println(sol.IsPalindrome3())
	fmt.Println(sol.IsPalindrome4())
}
