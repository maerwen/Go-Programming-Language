package main
import(
	"net/http"
	"fmt"
	"log"
)
// http://localhost:8080/ 访问此链接以测试
func main(){
	// 对于指定的访问路径的处理方法
	http.HandleFunc("/",handler)
	// 标注要监听的ip和端口号
	log.Fatal(http.ListenAndServe("localhost:8080",nil));
}
func handler(w http.ResponseWriter,r *http.Request){//处理程序回显http请求
	// 打印从请求中获取的请求方法，网址
	fmt.Fprintf(w,"%s %s %s\n", r.Method,r.URL, r.Proto)
	// 遍历请求头
	for k, v := range r.Header {
		fmt.Fprintf(w,"header[%q] = %q\n", k, v)
	}
	// 打印主机地址
	fmt.Fprintf(w,"host: %q\n", r.Host)
	// 打印客户端地址
	fmt.Fprintf(w,"remoteAddr: %q\n", r.RemoteAddr)
	// 如果请求解析parseform中出错，打印错误日志信息
	if err := r.ParseForm(); err != nil {//if语句中嵌套使用parseform。缩小了error变量的作用域
		log.Println(err)
	}
	// 打印请求的form
	for k, v := range r.Form {
		fmt.Fprintf(w,"form[%q]: %q\n]",k,v)
	}
}