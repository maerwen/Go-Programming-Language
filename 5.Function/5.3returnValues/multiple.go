package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	testFindlinks()
}
func testFindlinks() {
	s, err := findlinks("https://www.baidu.com/s?wd=%E4%BB%8A%E6%97%A5%E6%96%B0%E9%B2%9C%E4%BA%8B&tn=SE_PclogoS_8whnvm25&sa=ire_dl_gh_logo&rsv_dl=igh_logo_pcs")
	if err != nil {
		fmt.Printf("error:\t%v\n", err)
	}
	fmt.Printf("%d\n", len(s))
}
func findlinks(url string) ([]string, error) { //发起一个http的get请求,解析返回的html页面,并返回所有连接
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("Parsing %s as HTML:%v", url, err)
	}
	return visit([]string{}, doc), nil

}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
				fmt.Println(a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links

}

// 多返回值函数
func test() (value string, ok bool) {
	return //同 return "",false
	// kvs := make(map[string]string)
	// value, ok = kvs["key"]
	// return // 同于 return value,ok
}
