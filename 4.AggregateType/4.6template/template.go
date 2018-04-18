package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func main() {
	testTemplate()
}

var report, _ = template.New("report").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(templ) //模板对象创建
const templ = `{{}.TotalCount} issues:` +                                                       //输出符合条件的issues数量
	`{{range .Items}}-------------------` + //创建一个循环,展开内部值
	`Number:{{.Number}}` + //序号
	`User:	{{.User.Login}}` + //用户
	`Title:	{{.Title|printf "%.64s"}}` + //标题
	`Age:	{{.CreatedAt|daysAgo}} days` + //时间
	`{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
func testTemplate() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

const IssueURL = "https://api.github.com/search/issues"

// const IssueURL = "https://developer.github.com/v3/search/#search-issues"

type IssueSearchResult struct {
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
	Body      string
}
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func SearchIssues(terms []string) (*IssueSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssueURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed:%s", resp.Status)
	}
	var result IssueSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
