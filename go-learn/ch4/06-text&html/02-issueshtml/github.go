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

const templ = `
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`

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

	// 。template.Must辅助函数可以简化这个致命错误的处理：它接受一个模板和一个error类型的参数，检测error是否为nil（如果不是nil则发出panic异常），然后返回传入的模板。
	issueList := template.Must(template.New("issuelist").Parse(templ))

	// Execute 将解析后的模板应用到指定的数据对象，并将输出写入 wr。如果在执行模板或写入输出时发生错误，执行将停止，但部分结果可能已写入输出写入器。模板可以安全地并行执行，但如果并行执行共享一个写入器，输出可能会交错。
	// 如果数据是一个 reflect.Value，模板就会应用到 reflect.Value 所持有的具体值，如 fmt.Print.Print 模板。
	if err := issueList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

// 注意： >issues.html
// go run .\github.go repo:golang/go is:open json decoder >issues.html

// go run .\github.go repo:golang/go 3133 10535 >issues2.html
