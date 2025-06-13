package main

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	dataList := []struct {
		id        int
		direction string
		s         []int
		n         int
		want      []int
	}{
		{
			id:        1,
			direction: "left",
			s:         []int{0, 1, 2, 3, 4, 5},
			n:         2,
			want:      []int{2, 3, 4, 5, 0, 1},
		},
		{
			id:        2,
			direction: "right",
			s:         []int{0, 1, 2, 3, 4, 5},
			n:         2,
			want:      []int{4, 5, 0, 1, 2, 3},
		},
	}

	for _, v := range dataList {
		got := rotate(v.direction, v.s, v.n)
		if !reflect.DeepEqual(got, v.want) {
			t.Errorf("id: %d, got: %v, want: %v", v.id, got, v.want)
		}
	}
}

func BenchmarkRotate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rotate("left", []int{0, 1, 2, 3, 4, 5}, 2)
	}
}
