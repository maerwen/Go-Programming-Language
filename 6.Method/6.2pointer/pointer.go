package main

import (
	"fmt"
	"math"
)

func main() {
	// test()
	// var kvs Values //创建一个nil map,无法存入键值对,否则会panic
	var v = Values{"a": {"a-avlue"}}
	fmt.Println(v.Get("key"))
	v.add("key", "value")
	fmt.Println(v.Get("key"))
	v.add("b", "b-value")
	fmt.Println(v.Get("b"))
}

type Values map[string][]string

func (v Values) Get(key string) []string { //返回具有给定key的value
	if value := v[key]; value != nil && len(value) > 0 {
		return value
	}
	return []string{}
}
func (v Values) add(key string, value string) {
	if v[key] == nil {
		// v[key] = make([]string, 1, 1)
		v[key] = []string{}
	}
	v[key] = append(v[key], value)
}

/*
type P *int
func (p P) test() {}//编译错误,不允许本身是指针的类型去进行方法声明
*/
type Point struct {
	X, Y float64
}

func (p *Point) Radius() { //到原点的距离
	fmt.Printf("radius: %v\n", math.Hypot(p.X, p.Y))
}
func test() {
	p := Point{5, 5}
	p.Radius()
	q := &p
	q.Radius()
	// &Point{5, 5}.Radius() //编译错误,临时变量,无法获取地址
	// Point{5, 5}.Radius() //编译错误,临时变量,无法获取地址
}

type IntList struct { //简单的整数链表
	//IntList的nil代表空列表
	Value int
	Tail  *IntList
}

func (i *IntList) sum() int {
	if i.Tail == nil {
		return i.Value
	}
	return i.Value + i.Tail.sum()
}
