package main

import (
	"fmt"
	"reverse/reverseString"
)

func main() {
	str := "привет мир /     как дела"

	reverse := reverseString.New(str)

	fmt.Println(reverse.ReverseStandart())
	fmt.Println(reverse.ReverseNotStandart())
}
