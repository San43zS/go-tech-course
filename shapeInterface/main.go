package main

import (
	"fmt"

	"shapeInterface/flag"
)

func main() {
	f, err := flag.New()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	shape, err := f.GetShape()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	area := shape.Area()
	if err != nil {
		fmt.Printf("Calculation error: %v\n", err)
		return
	}

	fmt.Printf("Area of %s: %.2f\n", shape.GetType(), area)
}
