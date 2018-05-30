package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"unicode"
	"unicode/utf8"
)

func main() {
	// kvs1 := map[string]int{"ca": 1, "b": 2, "fc": 3, "d": 4, "e": 5}
	// kvs2 := map[string]int{"ca": 1, "b": 2, "fc": 3, "d": 4, "e": 6}
	// sortKeys(map[string]int{"ca": 1, "b": 2, "fc": 3, "d": 4, "e": 5})
	// fmt.Println(equal(kvs1, kvs2))
	// dedup()
	charcount()
}
func charcount() { //计算unicode字符的个数
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0
	in := bufio.NewReader(os.Stdin)
	for {
		r, n, error := in.ReadRune()
		if error == io.EOF {
			break
		}
		if error != nil {
			fmt.Fprintf(os.Stderr, "charcount:\t%v\n", error)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invaliud UTF-8 characters\n", invalid)
	}

}
func dedup() { //读取一系列的行,并且只输出每个行一次
	//创建map
	kvs := make(map[string]bool)
	// 获取控制台标准输入流
	input := bufio.NewScanner(os.Stdin)
	// 当有输入是,如果map不存在该输入内容,则存储,并打印到控制台
	for input.Scan() { //
		s := input.Text() //获取输出文本
		if !kvs[s] {
			kvs[s] = true
			fmt.Println(s)
		}
	}
	// 错误处理
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err) //%v为内置格式
		os.Exit(1)                          //返回0意味着成功,否则即为出错
	}
}

//map作为参数传入函数中时,kvs map[string]int
func sortKeys(kvs map[string]int) { //利用sort包的strings函数来进行键排序
	var keys []string
	for key, _ := range kvs {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Printf("%s\t%d\n", key, kvs[key])
	}
}
func equal(kvs1, kvs2 map[string]int) bool { //判断两个map是否相等
	// 先判断长度
	if len(kvs1) != len(kvs2) {
		return false
	}
	// 在判断key-value
	for k1, v1 := range kvs1 {
		if v2, ok := kvs2[k1]; !ok || v1 != v2 {
			return false
		}
	}
	return true
}
