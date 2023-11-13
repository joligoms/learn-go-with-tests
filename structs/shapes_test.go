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

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}

		checkArea(t, rectangle, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}

		checkArea(t, circle, 314.1592653589793)
	})
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
