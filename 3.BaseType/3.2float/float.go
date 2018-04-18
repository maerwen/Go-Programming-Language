// 根据一个三维曲面函数计算并生成svg
package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
)

const (
	width, height = 600, 320            //像素表示的画布大小
	cells         = 100                 //网格单元的个数
	xyrange       = 30.0                //坐标轴的范围（-xyrange~+xyrange）
	xyscale       = width / 2 / xyrange //x或y轴上每个单位长度的像素
	zscale        = height * 0.4        //z轴上每个单位长度的像素
	angle         = math.Pi / 6         //x、y轴的角度（30度）
)

var (
	sin30 = math.Sin(angle)
	cos30 = math.Cos(angle)
)

func main() {
	web()
}
func f(x, y float64) float64 {
	r := math.Hypot(x, y) //该坐标到原点的距离
	return math.Sin(r)
}
func corner(i, j int) (float64, float64) {
	// 求出网格单元（i，j）的顶点坐标
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// 计算曲面高度z
	z := f(x, y)
	// 将（x，y，z）等角投射到二维svg绘图平面上，坐标为（sx，sy）
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyrange - z*zscale
	return sx, sy
}
func showSvgOnWeb(w http.ResponseWriter, r *http.Request) { //显示svg图
	fmt.Fprintf(w,
		"<svg xmlns ='http://www.w3.org/2000/svg' "+
			"style='stroke: grey; fill: white;stroke-width: 0.7' "+
			"width='%d' height='%d'>\n",
		width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Fprintf(w,
				"\t<polugon points='%g, %g %g, %g %g, %g %g, %g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintf(w, "</svg>")
}
func web() {
	http.HandleFunc("/", showSvgOnWeb)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
