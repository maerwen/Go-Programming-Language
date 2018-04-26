package main

func main() {
	// x()
}
func x() { //数据竞态
	var x []int
	go func() {
		x = make([]int, 10)
	}()
	go func() {
		x = make([]int, 10000)
	}()
	x[9999] = 1 //未定义行为,可能造成内存异常
	// fmt.Println(x[9999]) //未定义行为,可能造成内存异常
}
