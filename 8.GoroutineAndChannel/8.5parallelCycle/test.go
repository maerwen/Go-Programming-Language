package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	test()
}
func print(i int) {
	//文件写入
	file, err := os.OpenFile("test.txt", os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	_, e := file.WriteString(fmt.Sprintf("%d\t", i))
	if e != nil {
		log.Fatal(e)
	}
}
func test() {
	arr := make([]int, 100)
	for i := 0; i < 100; i++ {
		arr[i] = i
	}
	for _, i := range arr {
		go print(i)
		// j := i
		// go print(j)
	}
}
