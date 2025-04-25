package complexity_analysis

import "testing"

func TestRecur(t *testing.T) {
	got := recur(10)
	want := 55

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestTailRecur(t *testing.T) {
	got := tailRecur(10, 0)
	want := 55

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestFib(t *testing.T) {
	got := fib(10)
	want := 34

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestForLoopRecur(t *testing.T) {
	got := forLoopRecur(10)
	want := 55

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
