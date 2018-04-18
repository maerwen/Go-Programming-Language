// 生成一个png格式的mandelbrot图形
package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math/cmplx"
	"net/http"
	"os"
)

func main() {
	// 	http.HandleFunc("/", network)
	// 	log.Fatal(http.ListenAndServe("localhost:8080", nil))
	local("test.png")
}
func network(w http.ResponseWriter, r *http.Request) {
	show(w)
}
func local(s string) {
	// 打开流
	// file, _ := os.Open("test.png")
	file, _ := os.OpenFile("test.png", os.O_CREATE, os.ModePerm)
	show(file)
}
func show(w io.Writer) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// 点（x,y）表示复数值
			img.Set(px, py, mandelbrot(z))
		}
	}
	// 输出内容到流
	png.Encode(w, img)
	// 输出内容到文件
	// png.Encode(file, img)
	// png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
