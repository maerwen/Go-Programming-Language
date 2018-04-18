package main
import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"net/http"
	"log"
	"time"
	"strconv"
)
var palette = []color.Color{color.White,color.Black}
const (
	whiteIndex = 0
	blackIndex = 1
)
	
func main(){//产生随机李萨如图形的gif动画
	testWeb2()
}
func testWeb1(){//在浏览器显示李萨如图形
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		lissajous1(w)
	})
	log.Fatal(http.ListenAndServe("localhost:8080",nil))

}
func testWeb2(){//在浏览器显示李萨如图形
	// http://localhost:8080/asdsaf?res=0.005&cycles=20&size=50&nframes=32&delay=4
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		// 必须先去解析url
		// r.ParseForm();
		// 寻找参数key-value中key为“a”的values
		/* if len(r.Form["a"]) > 0 {
			values := r.Form["a"]
			for i, j := range values {
				fmt.Printf("%d\t%s",i,j)
			}
			} 
		*/
		query := r.URL.Query()
		// lissajous2(w， res， cycles, size, nframes, delay)
		res,_ := strconv.ParseFloat(query["res"][0],64)
		cycles,_ := strconv.Atoi(query["cycles"][0])
		size,_ := strconv.Atoi(query["size"][0])
		nframes,_ := strconv.Atoi(query["nframes"][0])
		delay,_ := strconv.Atoi(query["delay"][0])
		lissajous2(w,res,cycles,size,nframes,delay)
	})
	log.Fatal(http.ListenAndServe("localhost:8080",nil))

}
func testLocal(){//在本地通过传入的os.args来生成随机的李萨如图形
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous1(w)
		}
		http.HandleFunc("/",handler)
		log.Fatal(http.ListenAndServe("localhost:8080", nil))
		return
	}
	lissajous1(os.Stdout)
}
func lissajous1 (out io.Writer) {
	const (
		cycles =5//个数
		res = 0.001//分辨率
		size = 100//振幅
		nframes = 64//帧数
		delay = 8//延迟
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 1; i < nframes; i++ {
		rect := image.Rect(0,0,2*size+1,2*size+1)
		img := image.NewPaletted(rect,palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5),size+int(y*size+0.5),blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out,&anim)

}
func lissajous2(out io.Writer, res float64, cycles, size, nframes, delay int){
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 1; i < nframes; i++ {
		rect := image.Rect(0,0,2*size+1,2*size+1)
		img := image.NewPaletted(rect,palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			// 类型转换
			img.SetColorIndex(size+int(float64(x)*float64(size)+0.5),size+int(float64(y)*float64(size)+0.5),blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out,&anim)
}



