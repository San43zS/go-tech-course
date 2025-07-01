package figure

import (
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
	return model.PI * float64(c.Radius) * float64(c.Radius)
}

func (c Circle) GetType() string {
	return "circle"
}
