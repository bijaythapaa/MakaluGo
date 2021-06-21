package perimeter

import "math"

//type Rectangle struct {
//	Height float64
//	Width  float64
//}
//
//func Perimeter(rectangle Rectangle) float64 {
//	return 2 * (rectangle.Width + rectangle.Height)
//}
//
//func Area(rectangle Rectangle) float64 {
//	return rectangle.Height * rectangle.Width
//}

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Height float64
	Width  float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Height * rectangle.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
