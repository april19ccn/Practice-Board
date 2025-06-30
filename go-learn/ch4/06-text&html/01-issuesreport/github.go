// Package github provides a Go API for the GitHub issue tracker.
// See https://developer.github.com/v3/search/#search-issues.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"text/template"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	fmt.Printf("Search query: %s\n", q)
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	// We must close resp.Body on all execution paths.我们必须关闭所有执行路径上的 resp.Body。
	// (Chapter 5 presents 'defer', which makes this simpler.)  第 5 章提出了 'defer'，这使得这更简单。
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%d issues:\n", result.TotalCount)
	// fmt.Println(len(result.Items))
	// for _, item := range result.Items {
	// 	fmt.Printf("#%-5d %9.9s %.55s\n",
	// 		item.Number, item.User.Login, item.Title)
	// }

	report, err := template.New("report").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(templ)
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

// go run .\github.go repo:golang/go is:open json decoder
