package main
import (
	"fmt"
	"os"
	// os提供一些函数和变量
)
func main(){
	echo3()
}
// 输出命令行参数
func echo1(){
	var s, sep string
	// for循环
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
func echo2(){
	// 并行声明
	a,b,c,d := []int{1,2,3},"index","element"," "
	// range范围使用
	for i,j := range a {
		fmt.Printf("%s-%d%s%s-%d\n",b,i,d,c,j)
	}
}
func echo3(){
	a := []int{1,2,3}
	// 空标识符弃用无用的数据
	for _,j := range a {
		fmt.Println(j)
	}
}