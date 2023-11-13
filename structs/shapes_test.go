package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{Width: 10.0, Height: 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	assertTest(t, got, want)
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{Width: 12.0, Height: 6.0}
	got := Area(rectangle)
	want := 72.0

	assertTest(t, got, want)
}

func assertTest(t testing.TB, got, want float64) {
	t.Helper()

	if got != want {
		t.Errorf("got %.2f, want %2.f", got, want)
	}
}
