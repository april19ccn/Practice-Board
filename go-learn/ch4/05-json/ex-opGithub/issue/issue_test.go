package issue

import (
	"fmt"
	"testing"
)

var params = Params{
	Key:       "json decoder",
	Repo:      "golang/go",
	Is:        "issue",
	State:     "open",
	TimeRange: 0,
	Sort:      "",
	Page:      1,
	PerPage:   30,
}

func TestGetIssues(t *testing.T) {
	got, err := GetIssues(params)
	if err != nil {
		t.Fatal(err)
	}
	want := 1
	if got.TotalCount != want {
		t.Errorf("got %d issues, want %d", got.TotalCount, want)
	}

	fmt.Printf("%d issues:\n", got.TotalCount)
	fmt.Println(len(got.Items))
	for _, item := range got.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}

	t.Logf("got %d issues", got.TotalCount)
}
