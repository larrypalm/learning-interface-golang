package math

type Shape interface {
	Area() float64
	Perimeter() float64
}

func TotalArea[T Shape](shape T) float64 {
	return shape.Area()
}

func TotalPerimeter[T Shape](shape T) float64 {
	return shape.Perimeter()
}

func SizeCategory[T Shape](shape T) string {
	area := shape.Area()
	var category string

	if area < 5 {
		category = "small"
	} else if area <= 20 {
		category = "medium"
	} else {
		category = "large"
	}

	return category
}
