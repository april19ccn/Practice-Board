package main

import "testing"

func TestGetId(t *testing.T) {
	doc, err := getHTML(true, "./template.html")
	if err != nil {
		t.Fatal(err)
	}

	got := ElementByID(doc, "test1")
	want := "span"

	if got == nil {
		t.Fatalf("ElementByID() = %v, want %v", got, want)
	}

	if got.Data != want {
		t.Errorf("ElementByID() = %v, want %v", got.Data, want)
	}
}
