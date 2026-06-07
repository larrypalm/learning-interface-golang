package math

type Shape interface {
	Area() float64
}

func TotalArea[T Shape](shape T) float64 {
	return shape.Area()
}
