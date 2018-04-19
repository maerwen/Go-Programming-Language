package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}
type Path []Point //简单slice类型声明

func main() {
	q := Point{1, 1} //变量定义
	p := Point{1, 6}
	q.distance(p) //方法调用
	path := Path{q, p}
	path.distance()
}
func (p Point) distance(q Point) {
	fmt.Println(math.Hypot(p.X-q.X, p.Y-q.Y))
}
func (p Path) distance() {
	sum := 0.0
	for i, j := range p {
		if i > 0 {
			sum += math.Hypot(j.X-p[i-1].X, j.Y-p[i-1].Y)
		}
	}
	fmt.Println(sum)
}
