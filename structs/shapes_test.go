package structs

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0

	assertTest(t, got, want)
}

func TestArea(t *testing.T) {
	got := Area(12.0, 6.0)
	want := 72.0

	assertTest(t, got, want)
}

func assertTest(t testing.TB, got, want float64) {
	t.Helper()

	if got != want {
		t.Errorf("got %.2f, want %2.f", got, want)
	}
}
