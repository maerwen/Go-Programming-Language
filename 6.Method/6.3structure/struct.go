package main

import (
	"fmt"
	"sync"
)

func main() {
	// testMan()
	fmt.Println(LookUp("key"))
	Put("key", "value")
	fmt.Println(LookUp("key"))
}

/* 原代码
var (
	mu      sync.Mutex //锁
	mapping = make(map[string]string)
)

func LookUp(key string) string {
	mu.Lock()
	v := mapping[key]
	mu.Unlock()
	return v
} */
// cache实现
var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func LookUp(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}
func Put(key, value string) {
	cache.Lock()
	cache.mapping[key] = value
	cache.Unlock()
}

type Point struct {
	int
	*uint
}
type Name string

func (n Name) show() {
	fmt.Println(n)
}

type Age int

func (a Age) show() {
	fmt.Println(a)
}

type Man struct {
	Name Name
	Age
	Sex bool
}

func (m Man) show() {
	fmt.Println(m)
}
func testMan() {
	// n := Man{
	// 	"git", 17, true}
	// n.show()
	m := Man{
		"git", 17, true, //如果在{}内部要换行,必须有,
	}
	m.show()
	m.Age.show()
	m.Name.show()
}
