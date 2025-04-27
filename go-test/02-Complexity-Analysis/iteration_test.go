package complexity_analysis

import "testing"

func TestForLoop(t *testing.T) {
	got := forLoop(10)
	want := 55

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestWhileLoop(t *testing.T) {
	got := whileLoop(10)
	want := 55

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestWhileLoopII(t *testing.T) {
	got := whileLoopII(10)
	want := 15

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
