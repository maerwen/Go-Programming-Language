package main

import "fmt"

func main() {
	testCounter()
}

type Buffer struct { //模拟bytes.Buffer类型缓存实现
	buf     []byte
	initial [64]byte
}

func (b Buffer) Grow(n int) { //新增n个数据,自动扩展容量
	if b.buf == nil {
		b.buf = b.initial[:0]
	}
	if len(b.buf)+n > cap(b.buf) {
		buf := make([]byte, len(b.buf)+n, cap(b.buf)*2+n)
		copy(buf, b.buf)
		b.buf = buf
	}
	fmt.Printf("%d\t%d\n", len(b.buf), cap(b.buf))
}
func testBuffer() {
	var b Buffer
	for i := 0; i < 6; i++ {
		b.Grow(i)
	}
}

type Counter struct { //模拟简单的计数器
	n int
}

func (c Counter) N() int { //获取值
	return c.n
}
func (c Counter) Increase1() { //自增,不影响c本身
	c.n++
}
func (c *Counter) Increase2() { //自增,改变c本身
	c.n++
}
func (c *Counter) Reset() { //重置,改变c本身
	c.n = 0
}
func testCounter() {
	var c Counter
	fmt.Println(c.N())
	c.Increase1()
	fmt.Println(c.N())
	c.Increase2()
	fmt.Println(c.N())
	c.Reset()
	fmt.Println(c.N())
}
