package figure

import "shapeInterface/shape"

type rectangle struct {
	width  int
	height int
}

func NewRectangle(width int, height int) shape.Shape {
	return &rectangle{width: width, height: height}
}

func (r rectangle) Area() float64 {
	return float64(r.width) * float64(r.height)
}

func (r rectangle) GetType() string {
	return "rectangle"
}
