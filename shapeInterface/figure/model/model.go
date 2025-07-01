package model

import "errors"

const (
	Shape      = "shape"
	RectWidth  = "width"
	RectHeight = "height"
	CircRadius = "radius"
	Rectangle  = "rectangle"
	Circle     = "circle"
)

var (
	ErrInvalidShapeType        = errors.New("invalid shape type: must be 'circle' or 'rectangle'")
	ErrRadiusMustBePositive    = errors.New("radius must be positive")
	ErrRectParamMustBePositive = errors.New("width and height must be positive")
	ErrUnknownShapeType        = errors.New("unknown shape type")
)
