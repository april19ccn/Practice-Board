// 创建一个web服务器，查询一次GitHub，然后生成BUG报告、里程碑和对应的用户信息。
package main

import (
	"example/learn/ch4/05-json/ex-opGithub/issue"
	"io"
	"log"
	"net/http"
	"text/template"
)

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

func GetIssues(out io.Writer) {
	params := issue.Params{
		Key:       "json decoder",
		Repo:      "golang/go",
		Is:        "issue",
		State:     "open",
		TimeRange: 0,
		Sort:      "",
		Page:      1,
		PerPage:   30,
	}

	result, err := issue.GetIssues(params)
	if err != nil {
		log.Fatal(err)
	}

	issueList := template.Must(template.New("issuelist").Parse(templ))
	if err := issueList.Execute(out, result); err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		GetIssues(w)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
