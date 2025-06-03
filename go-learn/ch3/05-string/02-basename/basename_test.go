package main

import "testing"

func TestBasename(t *testing.T) {
	funcList := []struct {
		handlerName string
		handler     func(string) string
	}{
		{
			handlerName: "basenameV1",
			handler:     basenameV1,
		},
		{
			handlerName: "basenameV2",
			handler:     basenameV2,
		},
		{
			handlerName: "basenameV3",
			handler:     basenameV3,
		},
		{
			handlerName: "basenameV4",
			handler:     basenameV4,
		},
	}

	dataList := []struct {
		data string
		want string
	}{
		{
			data: "a/b/c.go",
			want: "c",
		},
		{
			data: "c.d.go",
			want: "c.d",
		},
		{
			data: "abc",
			want: "abc",
		},
		{
			data: "d:/中文A/中文B/Test.go",
			want: "Test",
		},
	}

	for _, f := range funcList {
		for _, d := range dataList {
			got := f.handler(d.data)
			want := d.want
			if got != want {
				t.Errorf("func %s got %q want %q", f.handlerName, got, want)
			}
		}
	}
}
