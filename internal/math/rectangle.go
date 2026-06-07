package math

type Rectangle struct {
	Width  float64
	Height float64
}

func NewRectangle(width, height float64) Rectangle {
	return Rectangle{
		Width:  width,
		Height: height,
	}
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
