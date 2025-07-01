package flag

import (
	"flag"
	"fmt"
	"shapeInterface/figure"
	"shapeInterface/shape"
)

type Params struct {
	ShapeType string
	Width     int
	Height    int
	Radius    int
}

func New() (Iflag, error) {
	p := &Params{}

	flag.StringVar(&p.ShapeType, "shape", "", "shape type (circle/rectangle)")
	flag.IntVar(&p.Width, "width", 0, "rectangle width")
	flag.IntVar(&p.Height, "height", 0, "rectangle height")
	flag.IntVar(&p.Radius, "radius", 0, "circle radius")
	flag.Parse()

	if p.ShapeType != "circle" && p.ShapeType != "rectangle" {
		return Params{}, fmt.Errorf("invalid shape type: must be 'circle' or 'rectangle'")
	}

	return p, nil
}

func (p Params) GetShape() (shape.Shape, error) {
	switch p.ShapeType {
	case "circle":
		if p.Radius <= 0 {
			return nil, fmt.Errorf("radius must be positive")
		}
		return figure.NewCircle(p.Radius), nil
	case "rectangle":
		if p.Width <= 0 || p.Height <= 0 {
			return nil, fmt.Errorf("width and height must be positive")
		}
		return figure.NewRectangle(p.Width, p.Height), nil
	default:
		return nil, fmt.Errorf("unknown shape type")
	}
}
