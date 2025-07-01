package flag

import "shapeInterface/shape"

type Iflag interface {
	GetShape() (shape.Shape, error)
}
