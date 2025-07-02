package figure

import (
	"math"

	"shapeInterface/figure/model"
	"shapeInterface/shape"
)

type Circle struct {
	Radius int
}

func NewCircle(radius int) shape.Shape {
	return &Circle{Radius: radius}
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(float64(c.Radius), 2)
}

func (c Circle) GetType() string {
	return model.Circle
}
