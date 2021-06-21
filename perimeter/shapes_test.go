package perimeter

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10, 20}
	got := Perimeter(rectangle)
	want := 60.0

	if got != want {
		t.Errorf("want: %.2f, got: %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	//rectangle := Rectangle{10, 20}
	//got := Area(rectangle)
	//want := 200.0
	//
	//if got != want {
	//	t.Errorf("got: %.2f, want: %.2f", got, want)
	//}

	//t.Run("rectangles", func(t *testing.T) {
	//	rectangle := Rectangle{10, 20}
	//	//got := Area(rectangle)
	//	got := rectangle.Area()
	//	want := 200.0
	//
	//	if got != want {
	//		t.Errorf("got: %.2f, want: %.2f", got, want)
	//	}
	//})
	//
	//t.Run("circles", func(t *testing.T) {
	//	circle := Circle{10}
	//	got := circle.Area()
	//	want := 314.1592653589793
	//	if got != want {
	//		t.Errorf("got: %g, want: %g", got, want)
	//	}
	//})

	//checkArea := func(t testing.TB, shape Shape, want float64) {
	//	t.Helper()
	//	got := shape.Area()
	//	if got != want {
	//		t.Errorf("got: %g, want: %g", got, want)
	//	}
	//}
	//t.Run("rectangles", func(t *testing.T) {
	//	rectangle := Rectangle{12, 6}
	//	checkArea(t, rectangle, 72.0)
	//})
	//t.Run("circles", func(t *testing.T) {
	//	circle := Circle{10}
	//	checkArea(t, circle, 314.1592653589793)
	//})

	areaTests := []struct {
		name  string
		shape Shape
		//want  float64
		hasArea float64
	}{
		{"Rectangle", Rectangle{12, 6}, 72.0},
		{"Circle", Circle{10}, 314.1592653589793},
		{"Triangle", Triangle{12, 6}, 36.0},
	}
	for _, tt := range areaTests {
		//got := tt.shape.Area()
		//if got != tt.want {
		//	t.Errorf("got: %g, want: %g", got, tt.want)
		//}

		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got: %g, want: %g", tt.shape, got, tt.hasArea)
			}
		})
	}

}
