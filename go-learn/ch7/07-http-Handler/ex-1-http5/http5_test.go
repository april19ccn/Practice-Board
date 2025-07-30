// http 怎么通过依赖注入实现测试？
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttp(t *testing.T) {
	t.Run("Test Get List", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "list", nil)
		response := httptest.NewRecorder()

		db := database{"shoes": 50, "socks": 5}
		db.list(response, request)

		got := response.Body.String()
		want := "shoes: $50.00\nsocks: $5.00\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("Test Create List", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/create?item=hat&price=25.5", nil)
		response := httptest.NewRecorder()

		db := database{"shoes": 50, "socks": 5}
		db.create(response, request)

		got := response.Body.String()
		want := "create success! hat: $25.50\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("Test Update List", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/update?item=hat&price=45", nil)
		response := httptest.NewRecorder()

		db := database{"shoes": 50, "socks": 5}
		db.update(response, request)

		got := response.Body.String()
		want := "update success! hat: $45.00\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		request, _ = http.NewRequest("GET", "/update?item=hat&price=80", nil)
		response = httptest.NewRecorder()

		db.update(response, request)

		got = response.Body.String()
		want = "update success! hat: $80.00\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("Test Delete List", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/delete?item=hat", nil)
		response := httptest.NewRecorder()

		db := database{"shoes": 50, "socks": 5}
		db.delete(response, request)

		got := response.Body.String()
		want := "no such item: \"hat\"\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("Test Read List", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/read?item=shoes", nil)
		response := httptest.NewRecorder()

		db := database{"shoes": 50, "socks": 5}
		db.read(response, request)

		got := response.Body.String()
		want := "shoes: $50.00\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
