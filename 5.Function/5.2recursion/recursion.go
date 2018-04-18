package main

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	// fmt.Println(eat(20, 0))
	testOutline()
}

type NodeType int32
type Attribute struct {
	Key, Val string
}
type Node struct {
	Type                    NodeType
	Data                    string
	Arrt                    []Attribute
	FirstChild, NextSibling *Node
}

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links

}
func findlinks() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "find links:%v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}
func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) //把标签压入栈
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = n.NextSibling {
		outline(stack, c)
	}

}
func testOutline() {
	// file, _ := os.Open("test.txt")
	// doc, err := html.Parse(file)
	// doc, err := html.Parse(os.Stdin)
	doc, err := html.Parse(bufio.NewReader(os.Stdin))
	if err != nil {
		fmt.Fprintf(os.Stderr, "outling: %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
}
func eat(count, days int) int { //吃苹果:20个,每天不能小于1个,每天吃剩下的个数的1/2+1个,问可以吃几天
	if count >= 1 {
		days++
		count -= (count/2 + 1)
		return eat(count, days)
	}
	return days
}
