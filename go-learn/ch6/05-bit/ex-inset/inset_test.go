package inset

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func TestOperate(t *testing.T) {
	t.Run("test Add and Has", func(t *testing.T) {
		t.Parallel()
		var set IntSet
		set.Add(1)
		set.Add(55)
		set.Add(64)
		set.Add(188)

		got := set.String()
		want := "{1 55 64 188}"
		if got != want {
			t.Errorf("set.String() = %q, want %q", got, want)
		}

		if !set.Has(1) {
			t.Errorf("set.Has(1) = false, want true")
		}
		if !set.Has(55) {
			t.Errorf("set.Has(55) = false, want true")
		}
		if !set.Has(64) {
			t.Errorf("set.Has(64) = false, want true")
		}
		if !set.Has(188) {
			t.Errorf("set.Has(188) = false, want true")
		}
	})

	t.Run("test Len", func(t *testing.T) {
		t.Parallel()
		var set IntSet
		set.Add(1)
		set.Add(55)
		set.Add(64)
		set.Add(188)
		set.Add(1889)

		got := set.Len()
		want := 5
		if got != want {
			t.Errorf("set.Len() = %d, want %d", got, want)
		}
	})

	t.Run("test Remove", func(t *testing.T) {
		t.Parallel()
		var set IntSet
		set.Add(1)
		set.Add(55)
		set.Add(64)
		set.Add(188)

		// 移除存在的数据
		set.Remove(55)
		got := set.String()
		want := "{1 64 188}"
		if got != want {
			t.Errorf("set.String() = %q, want %q", got, want)
		}

		// 移除不存在的数据（检测是否有对越界的处理）
		set.Remove(990)
		got = set.String()
		want = "{1 64 188}"
		if got != want {
			t.Errorf("set.String() = %q, want %q", got, want)
		}
	})

	t.Run("test Clear", func(t *testing.T) {
		t.Parallel()
		var set IntSet
		set.Add(1)
		set.Add(55)
		set.Add(64)
		set.Add(188)

		set.Clear()
		got := set.String()
		want := "{}"
		if got != want {
			t.Errorf("set.String() = %q, want %q", got, want)
		}
	})

	t.Run("test Copy", func(t *testing.T) {
		t.Parallel()
		var set IntSet
		set.Add(1)
		set.Add(55)
		set.Add(64)
		set.Add(188)

		got := set.Copy()
		want := "{1 55 64 188}"
		if got.String() != want {
			t.Errorf("set.String() = %q, want %q", got, want)
		}

		setPtr := (*reflect.SliceHeader)((unsafe.Pointer(&set.words)))
		gotPtr := (*reflect.SliceHeader)((unsafe.Pointer(&got.words)))
		fmt.Println(setPtr, gotPtr)
		if setPtr.Data == gotPtr.Data {
			t.Errorf("got is not a copy of set")
		}
	})

	t.Run("test AddAll", func(t *testing.T) {
		t.Parallel()
		var set IntSet
		set.AddAll(1, 55, 64, 188)

		got := set.String()
		want := "{1 55 64 188}"
		if got != want {
			t.Errorf("set.String() = %q, want %q", got, want)
		}
	})
}

func TestSet(t *testing.T) {
	t.Run("test UnionWith", func(t *testing.T) {
		t.Parallel()
		var set IntSet
		set.Add(1)
		set.Add(55)
		set.Add(64)
		set.Add(188)

		var set2 IntSet
		set2.Add(8)
		set2.Add(55)
		set2.Add(64)
		set2.Add(298)

		set.UnionWith(&set2)

		got := set.String()
		want := "{1 8 55 64 188 298}"
		if got != want {
			t.Errorf("set.String() = %q, want %q", got, want)
		}
	})

	t.Run("test IntersectWith", func(t *testing.T) {
		t.Parallel()
		var set IntSet
		set.Add(1)
		set.Add(55)
		set.Add(64)
		set.Add(188)
		set.Add(375)

		var set2 IntSet
		set2.Add(8)
		set2.Add(55)
		set2.Add(64)
		set2.Add(298)

		set.IntersectWith(&set2)

		got := set.String()
		want := "{55 64}"
		if got != want {
			t.Errorf("set.String() = %q, want %q", got, want)
		}
	})

	t.Run("test DifferenceWith", func(t *testing.T) {
		t.Parallel()
		var set IntSet
		set.Add(1)
		set.Add(55)
		set.Add(64)
		set.Add(188)
		set.Add(375)

		var set2 IntSet
		set2.Add(8)
		set2.Add(55)
		set2.Add(64)
		set2.Add(298)

		set.DifferenceWith(&set2)

		got := set.String()
		want := "{1 188 375}"
		if got != want {
			t.Errorf("set.String() = %q, want %q", got, want)
		}
	})

	t.Run("test SymmetricDifference", func(t *testing.T) {
		t.Parallel()
		var set IntSet
		set.Add(1)
		set.Add(55)
		set.Add(64)
		set.Add(188)
		set.Add(375)

		var set2 IntSet
		set2.Add(8)
		set2.Add(55)
		set2.Add(64)
		set2.Add(298)

		set.SymmetricDifference(&set2)

		got := set.String()
		want := "{1 8 188 298 375}"
		if got != want {
			t.Errorf("set.String() = %q, want %q", got, want)
		}
	})
}

func TestElems(t *testing.T) {
	t.Parallel()
	var set IntSet
	set.Add(1)
	set.Add(55)
	set.Add(64)
	set.Add(188)

	got := set.Elems()
	want := []int{1, 55, 64, 188}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("set.Elems() = %v, want %v", got, want)
	}
}

// 验证同一个包里的封装性，可以访问到小写的words
func TestEcapsulation(t *testing.T) {
	var got = IntSet{}
	got.Add(1)
	got.Add(55)
	got.Add(64)
	got.Add(188)
	fmt.Println(got.words)
}

func BenchmarkIntSetAdd(b *testing.B) {
	var s IntSet
	for i := 0; i < b.N; i++ {
		s.Add(i)
	}
}

func BenchmarkIntSetAddPro(b *testing.B) {
	var s IntSet
	for i := 0; i < b.N; i++ {
		s.AddPro(i)
	}
}
