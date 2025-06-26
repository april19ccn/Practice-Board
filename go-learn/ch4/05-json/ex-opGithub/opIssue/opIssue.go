package opissue_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// 不应提交的部分
const testOWNER = ""
const testREPO = ""
const testToken = ""

const IssuesURL = "https://api.github.com/" + "repos/" + testOWNER + "/" + testREPO + "/issues"

type IssueParams struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type IssueData struct {
	Id     int
	Number int
	Title  string
	State  string
}

// 统一处理
func handleResponse(req *http.Request, okCode int) (*IssueData, error) {
	req.Header.Add("Authorization", testToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Body)
	if resp.StatusCode != okCode {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssueData
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

// 创建 issue
func CreateIssue(params IssueParams) (*IssueData, error) {
	// 正确序列化
	requestBody := IssueParams{
		Title: params.Title,
		Body:  params.Body,
	}

	data, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	// data := []byte(fmt.Sprintf(`{"title": "%s", "body": "%s"}`, params.Title, params.Body))
	fmt.Println("SENDING JSON:", string(data))

	req, err := http.NewRequest("POST", IssuesURL, bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}

	return handleResponse(req, 201)
}

// 获取 issue
func GetIssue(issueNumber int) (*IssueData, error) {
	req, err := http.NewRequest("GET", IssuesURL+"/"+strconv.Itoa(issueNumber), nil)
	if err != nil {
		log.Fatal(err)
	}

	return handleResponse(req, 200)
}

// 更新 issue
func UpdateIssue(issueNumber int, params IssueParams) (*IssueData, error) {
	data := []byte(fmt.Sprintf(`{"title": "%s", "body": "%s"}`, params.Title, params.Body))

	req, err := http.NewRequest("PATCH", IssuesURL+"/"+strconv.Itoa(issueNumber), bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}

	return handleResponse(req, 200)
}

// 关闭 issue
func CloseIssue(issueNumber int) (*IssueData, error) {
	data := []byte(`{"state": "closed"}`)
	req, err := http.NewRequest("PATCH", IssuesURL+"/"+strconv.Itoa(issueNumber), bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}

	return handleResponse(req, 200)
}
