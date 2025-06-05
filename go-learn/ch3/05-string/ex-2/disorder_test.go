package main

import (
	"testing"
)

type Disorder struct {
	left  string
	right string
}

var dataList = []struct {
	id   int8
	got  Disorder
	want bool
}{
	{
		id:   1,
		got:  Disorder{left: "abc123中文", right: "文123abc中"},
		want: true,
	},
	{
		id:   2,
		got:  Disorder{"852963741", "159784623"},
		want: true,
	},
	{
		id:   3,
		got:  Disorder{"abc123中文12", "文123abc中"},
		want: false,
	},
	{
		id:   4,
		got:  Disorder{"123", "123"},
		want: false,
	},
	{
		id:   5,
		got:  Disorder{"123,*!@#$", ",*!@#$123"},
		want: true,
	},
	{
		id:   6,
		got:  Disorder{"123,*!@#$aaaaaabbbb", ",*!@#$123bbbbaaaaa"},
		want: false,
	},
	{
		id:   7,
		got:  Disorder{"aabb", "aacb"},
		want: false,
	},
}

var funcList = []func(string, string) bool{
	disorderArr,
	disorderMap,
}

func TestDisorder(t *testing.T) {
	for fid, f := range funcList {
		for _, v := range dataList {
			if f(v.got.left, v.got.right) != v.want {
				t.Errorf("func %d, id: %d, got %v, want %v", fid, v.id, v.got, v.want)
			}
		}
	}
}

// 1483370               831.4 ns/op           320 B/op          4 allocs/op
func BenchmarkDisorderArr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		disorderArr("852963741", "159784623")
	}
}

// 2100723               545.0 ns/op           328 B/op          3 allocs/op
func BenchmarkDisorderMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		disorderMap("852963741", "159784623")
	}
}
