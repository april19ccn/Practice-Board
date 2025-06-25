package main

import (
	// issue "example/learn/ch4/05-json/ex-opGithub/issue"
	opIssue "example/learn/ch4/05-json/ex-opGithub/opIssue"
	"example/learn/ch4/05-json/ex-opGithub/wIssue"
	"fmt"
	"log"
)

// var params = issue.Params{
// 	Key:       "json decoder",
// 	Repo:      "golang/go",
// 	Is:        "issue",
// 	State:     "open",
// 	TimeRange: 3,
// 	Sort:      "",
// 	Page:      1,
// 	PerPage:   30,
// }

var params = opIssue.CreateParams{
	Title: "test10",
	Body:  "test test",
}

func main() {
	// result, err := issue.GetIssues(params)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%d issues:\n", result.TotalCount)
	// fmt.Println(len(result.Items))
	// for _, item := range result.Items {
	// 	fmt.Printf("#%-5d %9.9s %.55s\n",
	// 		item.Number, item.User.Login, item.Title)
	// }

	params.Body, _ = wIssue.GetIssueDescription()
	fmt.Println(params)
	result, err := opIssue.CreateIssue(params)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("issues number: %d \n", result.Number)
	fmt.Println(result)
}
