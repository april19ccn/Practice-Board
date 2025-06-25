package issue

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
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

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Params struct {
	Key       string
	Repo      string
	Is        string
	State     string
	TimeRange int // 查询范围
	Sort      string
	Page      int
	PerPage   int
}

func GetTime() []string {
	t := time.Now()

	result := []string{}

	result = append(result, "")
	result = append(result, "created:>"+t.AddDate(0, -1, 0).Format("2006-01-02"))  // 本月
	result = append(result, "created:>"+t.AddDate(0, -6, 0).Format("2006-01-02"))  // 6个月
	result = append(result, "created:>"+t.AddDate(-1, 0, 0).Format("2006-01-02"))  // 1年
	result = append(result, "created:<="+t.AddDate(-1, 0, 0).Format("2006-01-02")) // 超过1年

	return result
}

func SearchIssues(timeRanges []string, params Params) (*IssuesSearchResult, error) {
	q := fmt.Sprintf(
		"is:%s repo:%s state:%s %s %s",
		params.Is,
		params.Repo,
		params.State,
		timeRanges[params.TimeRange],
		params.Key,
	)

	// 创建 url.Values 存储所有查询参数
	values := url.Values{}
	values.Set("q", q)                                   // 设置 q 参数
	values.Set("page", strconv.Itoa(params.Page))        // 设置 page 参数
	values.Set("per_page", strconv.Itoa(params.PerPage)) // 设置 per_page 参数
	values.Set("sort", params.Sort)

	fmt.Println(IssuesURL + "?" + values.Encode())
	resp, err := http.Get(IssuesURL + "?" + values.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	return &result, nil
}

func GetIssues(params Params) (*IssuesSearchResult, error) {
	return SearchIssues(GetTime(), params)
}
