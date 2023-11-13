package structs

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(rec Rectangle) float64 {
	return 2 * (rec.Width + rec.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return 0
}
