package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func main() {
	test3()
}
func test1() { //初始版本
	//根目录获取
	// 在当前打开的终端命令行所在目录开始遍历
	// flag.Parse()
	// roots := flag.Args()
	// if len(roots) == 0 {
	// 	roots = []string{"."}
	// }
	// var roots = []string{"/"} //系统盘根目录
	var roots = []string{"/home/wzx"}
	//传送文件大小的通道构建
	filesizes := make(chan int64)
	go func() {
		//开始便利目录集合,同时进行递归操作
		for _, filename := range roots {
			walkDir1(filename, filesizes)
		}
		//关闭通道
		close(filesizes)
	}()
	//结果输出
	var nfiles, nbytes int64
	for size := range filesizes {
		nfiles++
		nbytes += size
	}
	printInfo(nfiles, nbytes)
}

// test2
// 高并发下线程同时开启太多,会导致内存耗尽,所以对并发数进行限制
// 对最底层的方法dirents2进行控制更有效
var sema = make(chan struct{}, 20)

func test2() { //并发版本
	// var roots = []string{"/"}
	var roots = []string{"/home/wzx/Documents"}
	//传送文件大小的通道构建
	filesizes := make(chan int64)
	//开始便利目录集合,同时进行递归操作
	var n sync.WaitGroup
	for _, filename := range roots {
		//计数器增1
		n.Add(1)
		go walkDir2(&n, filename, filesizes)
	}
	//关闭通道
	go func() {
		n.Wait() //计数器为0时关闭通道
		close(filesizes)
	}()
	//结果输出
	var nfiles, nbytes int64
	for size := range filesizes {
		nfiles++
		nbytes += size
	}
	printInfo(nfiles, nbytes)
}

// test3
// 设置提示符及意义
var v = flag.Bool("v", false, "显示过程进展")

func test3() { //不使用高并发,单线程持续进行统计,并在输入-v命令符执行时间歇性打印目前结果
	var roots = []string{"/home/wzx"}
	//传送文件大小的通道构建
	filesizes := make(chan int64)
	go func() {
		//开始便利目录集合,同时进行递归操作
		for _, filename := range roots {
			walkDir1(filename, filesizes)
		}
		//关闭通道
		close(filesizes)
	}()
	//结果定期输出
	//接收通道
	var tick <-chan time.Time
	//如果输入该命令符
	var nfiles, nbytes int64
	if *v {
		fmt.Println(11111)
		tick = time.Tick(500 * time.Millisecond)
	loop:
		for {
			select {
			// 非时间通道,执行后台操作
			case size, ok := <-filesizes:
				if !ok {
					break loop
				}
				nfiles++
				nbytes += size
				// 时间通道,打印目前进展
			case <-tick:
				printInfo(nfiles, nbytes)
			}
		}
	} else {
		for size := range filesizes {
			nfiles++
			nbytes += size
		}
	}
	printInfo(nfiles, nbytes)
}

//递归遍历以dir为根目录的整个文件树
//发送每个已经找到的文件的大小到通道上面去
func walkDir1(dir string, filesizes chan<- int64) {
	//利用dirents来获取单个条目
	for _, entry := range dirents1(dir) {
		//判断是否是目录
		if entry.IsDir() {
			//是目录,进行递归操作
			walkDir1(filepath.Join(dir, entry.Name()), filesizes)
		} else {
			//不是目录.发送文件大小到通道
			filesizes <- entry.Size()
		}
	}
}

//多线程同时进行
func walkDir2(n *sync.WaitGroup, dir string, filesizes chan<- int64) {
	defer n.Done() //该goroutine执行完成后,计数器num-1
	//利用dirents来获取单个条目
	for _, entry := range dirents2(dir) {
		//判断是否是目录
		if entry.IsDir() {
			//计数器增1
			n.Add(1)
			//是目录,进行递归操作
			go walkDir2(n, filepath.Join(dir, entry.Name()), filesizes)
		} else {
			//不是目录.发送文件大小到通道
			filesizes <- entry.Size()
		}
	}
}

//dirents返回dir中目录的条目
func dirents1(dir string) []os.FileInfo {
	// 利用ioutil来读取目录信息
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Printf("%v\n", err)
	}
	return fileInfos
}
func dirents2(dir string) []os.FileInfo {
	//获取令牌
	sema <- (struct{}{})
	//释放令牌
	defer func() {
		<-sema
	}()
	// 利用ioutil来读取目录信息
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Printf("%v\n", err)
	}
	return fileInfos
}
func printInfo(nfiles, nbytes int64) {
	// fmt.Printf("文件个数:\t%d\n总大小\t%.1fGB\n", nfiles, float64(nbytes)/1e9)
	fmt.Printf("文件个数:\t%d\n总大小:\t%.1fMB\n", nfiles, float64(nbytes)/1e6)
}
