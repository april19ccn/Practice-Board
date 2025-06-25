package opissue_test

import (
	wIssue "example/learn/ch4/05-json/ex-opGithub/wIssue"
	"fmt"
	"log"
	"testing"
)

var params = CreateParams{
	Title: "test12",
	Body:  "test test",
}

func TestGetIssue(t *testing.T) {
	result, err := GetIssue(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

func TestCreateIssue(t *testing.T) {
	result, err := CreateIssue(params)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("issues number: %d \n", result.Number)
	fmt.Println(result)

	checkSuccess(t, result)
}

func TestUpdateIssue(t *testing.T) {
	params := CreateParams{
		Title: "test-patch-3",
		Body:  "test test test",
	}
	result, err := UpdateIssue(5, params)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("issues number: %d \n", result.Number)
	fmt.Println(result)

	checkSuccess(t, result)
}

func TestCloseIssue(t *testing.T) {
	result, err := CloseIssue(3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("issues number: %d \n", result.Number)
	fmt.Println(result)

	checkSuccessClose(t, result)
}

func TestCreateIssueWithEditor(t *testing.T) {
	params.Body, _ = wIssue.GetIssueDescription()
	result, err := CreateIssue(params)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("issues number: %d \n", result.Number)
	fmt.Println(result)

	checkSuccess(t, result)
}

func checkSuccess(t *testing.T, issue *IssueData) {
	t.Helper()
	want, err := GetIssue(issue.Number)
	if err != nil {
		t.Error("get issue failed")
	}

	if want.Title != issue.Title {
		t.Error("create issue failed")
	}
}

func checkSuccessClose(t *testing.T, issue *IssueData) {
	t.Helper()
	want, err := GetIssue(issue.Number)
	if err != nil {
		t.Error("get issue failed")
	}

	if want.Title != issue.Title {
		t.Error("create issue failed")
	}

	if want.State != "closed" {
		t.Error("close issue failed")
	}
}
