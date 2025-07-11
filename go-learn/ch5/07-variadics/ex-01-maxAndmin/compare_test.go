package main

import "testing"

func TestMax(t *testing.T) {
	t.Parallel()
	// got, want := max(1, 2, 3, 4), 4 这是什么写法

	got, err := Max(1, 2, 3, 4)
	if err != nil {
		t.Fatal(err)
	}
	want := 4

	if got != want {
		t.Errorf("max(1, 2, 3, 4) = %d; want %d", got, want)
	}
}

func TestMin(t *testing.T) {
	t.Parallel()
	// got, err := Min(1, 2, 3, 4)
	got, err := Min()
	if err != nil {
		t.Fatal(err)
	}
	want := 1

	if got != want {
		t.Errorf("min(1, 2, 3, 4) = %d; want %d", got, want)
	}
}
