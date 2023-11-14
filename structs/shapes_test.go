package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{Width: 10.0, Height: 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	assertTest(t, got, want)
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()

		got := shape.Area()

		assertTestG(t, got, want)
	}

	areaTests := map[string]struct {
		shape Shape
		want  float64
	}{
		"rectangles": {
			shape: Rectangle{12, 6},
			want:  72.0,
		},
		"circles": {
			shape: Circle{10},
			want:  314.1592653589793,
		},
		"triangles": {
			shape: Triangle{12, 6},
			want:  36.0,
		},
	}

	for testName, test := range areaTests {
		t.Run(testName, func(t *testing.T) {
			checkArea(t, test.shape, test.want)
		})
	}
}

func assertTest(t testing.TB, got, want float64) {
	t.Helper()

	if got != want {
		t.Errorf("got %.2f, want %2.f", got, want)
	}
}

func assertTestG(t testing.TB, got, want float64) {
	t.Helper()

	if got != want {
		t.Errorf("got %g, want %g", got, want)
	}
}
