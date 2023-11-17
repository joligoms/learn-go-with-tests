package mocks

import (
	"bytes"
	"reflect"
	"testing"
)

type SpyCountdownCheck struct {
	Calls []string
}

const (
	sleep = "sleep"
	write = "write"
)

func (s *SpyCountdownCheck) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownCheck) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spyCountdown := &SpyCountdownCheck{}

		Countdown(spyCountdown, spyCountdown)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownCheck{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("got %v want %v", spySleepPrinter.Calls, want)
		}
	})

}
