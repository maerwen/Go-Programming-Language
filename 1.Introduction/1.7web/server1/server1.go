package main
import (
	"fmt"
	"log"
	"net/http"
)
// 搭建一个迷你回声服务器
func main(){
	// http://localhost:8080/adasfnas	可在浏览器中访问该地址
	http.HandleFunc("/",handler)//回声请求调用处理程序
	log.Fatal(http.ListenAndServe("localhost:8080",nil))
}
func handler(w http.ResponseWriter,r *http.Request){//处理程序回显请求URL r的路径部分
	fmt.Fprintf(w,"URL.Path = %q\n",r.URL.Path)
}