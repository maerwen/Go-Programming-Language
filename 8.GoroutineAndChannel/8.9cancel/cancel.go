package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
)

func main() {
	// test1()
	test2()
}

// 1.先创建一个取消通道done,用来在各gouroutine中来传递取消状态和讯息
var done = make(chan struct{})

// 2.再定义一个取消函数,在调用时检测或者轮询取消状态
func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}
func test1() {
	// 3.然后创建一个goroutine,在获取到输入的情况下来通过关闭取消通道来传播取消事件
	go func() {
		defer close(done)
		os.Stdin.Read(make([]byte, 1))
		done <- struct{}{}
	}()
	// var roots = []string{"/"}
	var roots = [1000]string{}
	for i, _ := range roots {
		roots[i] = "/home/wzx/Documents"
	}
	//传送文件大小的通道构建
	filesizes := make(chan int64)
	//开始便利目录集合,同时进行递归操作
	var n sync.WaitGroup
	for _, filename := range roots {
		//计数器增1
		n.Add(1)
		go walkDir1(&n, filename, filesizes)
	}
	//关闭通道
	go func() {
		n.Wait() //计数器为0时关闭通道
		close(filesizes)
	}()
	//结果输出
	var nfiles, nbytes int64
	// 4.接下来在主goroutine中来运用select多路复用,在里面包含之前通道接收操作的同时,新增一种从取消通道done来接收信息的情况.
	// 		4.----新增的case实现在从取消通道接收消息后,(在确认原有情况的接收通道已耗尽且关闭后,但在代码加入后会造成死锁...)直接return
loop:
	for {
		select {
		case size, ok := <-filesizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-done:
			// for range filesizes {//加入该代码会报错??????
			// }
			fmt.Println("已取消操作!")
			return
		}
	}
	printInfo(nfiles, nbytes)

}
func test2() { // 4,5替代解决方案:在核心关键代码位置,通过轮询取消状态或者在select语句中来新增一种接收done通道的状况,来返回一些导致程序取消的值
	go func() {
		defer close(done)
		os.Stdin.Read(make([]byte, 1))
		done <- struct{}{}
	}()
	// var roots = []string{"/"}
	var roots = [1000]string{}
	for i, _ := range roots {
		roots[i] = "/home/wzx/Documents"
	}
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
loop:
	for {
		select {
		case size, ok := <-filesizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-done:
			// for range filesizes { //加入该代码会报错??????
			// }
			// fmt.Println("已取消操作!")
			return
		}
	}
	printInfo(nfiles, nbytes)
}

// 高并发下线程同时开启太多,会导致内存耗尽,所以对并发数进行限制
// 对最底层的方法dirents2进行控制更有效
var sema = make(chan struct{}, 20)

//递归遍历以dir为根目录的整个文件树
//发送每个已经找到的文件的大小到通道上面去
//多线程同时进行
func walkDir1(n *sync.WaitGroup, dir string, filesizes chan<- int64) {
	// 5.在我们想要执行/响应取消操作的goroutine中的核心代码位置,轮询取消状态.如果状态已设置,则return.
	if cancelled() {
		return
	}
	defer n.Done() //该goroutine执行完成后,计数器num-1
	//利用dirents来获取单个条目
	for _, entry := range dirents1(dir) {
		//判断是否是目录
		if entry.IsDir() {
			//计数器增1
			n.Add(1)
			//是目录,进行递归操作
			go walkDir1(n, filepath.Join(dir, entry.Name()), filesizes)
		} else {
			//不是目录.发送文件大小到通道
			filesizes <- entry.Size()
		}
	}
}
func walkDir2(n *sync.WaitGroup, dir string, filesizes chan<- int64) {
	if cancelled() {
		return
	}
	defer n.Done() //该goroutine执行完成后,计数器num-1
	//利用dirents来获取单个条目
	// 4,5替代解决方案:在核心关键代码位置,通过轮询取消状态或者在select语句中来新增一种接收done通道的状况,来返回一些导致程序取消的值
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

//dirents返回dir中目录的条目
func dirents2(dir string) []os.FileInfo {
	// 4,5替代解决方案:在核心关键代码位置,通过轮询取消状态或者在select语句中来新增一种接收done通道的状况,来返回一些导致程序取消的值
	select {
	case sema <- struct{}{}:
		//获取令牌
	case <-done:
		return nil
	}
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
