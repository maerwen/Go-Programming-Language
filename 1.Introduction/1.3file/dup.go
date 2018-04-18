package main
import(
	"bufio"
	"fmt"
	"os"
	"io/ioutil"
	"strings"
)
func main(){
	dup3()
}

func dup1(){// 输出标准输入中出现次数大于1的行，前面是次数
	// 用make函数创建一个map
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		// counts[input.Text()]++
		/* 
			上面代码等价于
			line := input.Text()
			count[line] += 1
		*/
		line := input.Text()
		counts[line] = counts[line] + 1
		fmt.Println(counts[line])
		/* 
			fmt.Println(counts[input.Text()])
			if counts[input.Text()] > 3 {
				break
			}
		 */
	}
	for k,v := range counts{
		fmt.Printf("%s-%d\n",k,v)
	}
}


func dup2(){//打印输入中多次出现的行的个数和文本
	counts := make(map[string]int)
	// 新建一个数组，从os的参数列表以第二个开始顺序截取
	files := os.Args[1:]
	// len函数用于求取数组的长度
	if len(files) == 0 {
		countLines(os.Stdin,counts)
	} else {
		for _,file := range files {
			// os.open方法有两个返回值
			f,error := os.Open(file)
			if error != nil {
				fmt.Fprintf(os.Stderr,"dup2: %v\n",error)
				continue
			}
			countLines(f,counts)
			f.Close()
		}
	}
	for k,v := range counts {
		if v > 1 {
			fmt.Printf("%s\t%d\n",k,v)
		}
	}

}
func countLines(f *os.File, counts map[string]int){
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func dup3(){//
	counts := make(map[string]int)
	// 往map里面插入数据
	counts["dup.go"] = 1
	for _, filename := range os.Args[1:]{
		// ioutil的ReadFile函数返回一个可以转换成字符串的字节slice
		data, error := ioutil.ReadFile(filename)
		if error != nil {
			fmt.Fprintf(os.Stderr,"dup3: %v\n",error)
			continue
		}
		// strings的split函数与join函数功用相反
		for _, line  := range strings.Split(string(data),"\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%s\t%d\n",line,n)
		}
	}
}

