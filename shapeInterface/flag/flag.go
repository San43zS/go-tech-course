package flag

import (
	"flag"
	"shapeInterface/figure"
	"shapeInterface/figure/model"
	"shapeInterface/shape"
)

type Params struct {
	ShapeType string
	Width     int
	Height    int
	Radius    int
}

func New() (Iflag, error) {
	p := Params{}

	flag.StringVar(&p.ShapeType, model.Shape, "", "shape type (circle/rectangle)")
	flag.IntVar(&p.Width, model.RectWidth, 0, "rectangle width")
	flag.IntVar(&p.Height, model.RectHeight, 0, "rectangle height")
	flag.IntVar(&p.Radius, model.CircRadius, 0, "circle radius")
	flag.Parse()

	if p.ShapeType != model.Circle && p.ShapeType != model.Rectangle {
		return Params{}, model.ErrInvalidShapeType
	}

	return p, nil
}

func (p Params) GetShape() (shape.Shape, error) {
	switch p.ShapeType {
	case model.Circle:
		if p.Radius <= 0 {
			return nil, model.ErrRadiusMustBePositive
		}

		return figure.NewCircle(p.Radius), nil

	case model.Rectangle:
		if p.Width <= 0 || p.Height <= 0 {
			return nil, model.ErrRectParamMustBePositive
		}

		return figure.NewRectangle(p.Width, p.Height), nil

	default:

		return nil, model.ErrUnknownShapeType
	}
}
