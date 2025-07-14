package main

import (
	"testing"
)

func TestJoin(t *testing.T) {
	tests := []struct {
		name     string
		sep      string
		elements []string
		want     string
	}{
		{ // 单个元素
			name:     "single element",
			sep:      ",",
			elements: []string{"hello"},
			want:     "hello",
		},
		{ // 多个元素
			name:     "two elements",
			sep:      ",",
			elements: []string{"hello", "world"},
			want:     "hello,world",
		},
		{ // 自定义分隔符
			name:     "three elements with custom separator",
			sep:      "---",
			elements: []string{"a", "b", "c"},
			want:     "a---b---c",
		},
		{ // 空分隔符
			name:     "empty separator",
			sep:      "",
			elements: []string{"a", "b", "c"},
			want:     "abc",
		},
		{ // 空元素切片
			name:     "empty elements",
			sep:      ",",
			elements: []string{},
			want:     "",
		},
		{ // 包含空字符串的元素1
			name:     "one empty element",
			sep:      ",",
			elements: []string{""},
			want:     "",
		},
		{ // 包含空字符串的元素2
			name:     "multiple empty elements",
			sep:      ",",
			elements: []string{"", "", ""},
			want:     ",,",
		},
		{ // 混合空字符串和非空字符串
			name:     "mixed elements",
			sep:      "|",
			elements: []string{"a", "", "b", "c", ""},
			want:     "a||b|c|",
		},
		{ // 特殊字符分隔符
			name:     "separator with special characters",
			sep:      "🌧️",
			elements: []string{"rain", "clouds"},
			want:     "rain🌧️clouds",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Join(tt.sep, tt.elements...)
			if got != tt.want {
				t.Errorf("Join() = %q, want %q", got, tt.want)
			}
		})
	}
}

// 基准测试
func BenchmarkJoin(b *testing.B) {
	elements := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	sep := ","

	for i := 0; i < b.N; i++ {
		Join(sep, elements...)
	}
}

// 示例用法
func ExampleJoin() {
	// Example usage
	result := Join(" ", "hello", "world", "!")
	println(result)
	// Output: hello world !
}
