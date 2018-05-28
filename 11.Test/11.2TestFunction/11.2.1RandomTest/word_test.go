package word

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 随机测试
// 返回一个随机回文字符串，它的长度随机生成
func randomStringMaker(num *rand.Rand) string {
	// 字符串长度为0~24
	n := num.Intn(25)
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		// 随机字符最大为`\u0999`
		r := rune(num.Intn(0x1000))
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

/* func TestIsPalindrome(t *testing.T) {
	// 初始化一个随机数生成器
	seed := time.Now().UTC().UnixNano()
	t.Logf("random seed:%d", seed)
	num := rand.New(rand.NewSource(seed))

	// 1000		0.006s
	// 10000		0.025s
	// 100000		0.192s
	// 1000000		1.910s
	// 10000000	18.860s

	for i := 0; i < 1000; i++ {
		s := randomStringMaker(num)
		if !IsPalindrome(s) {
			t.Errorf("IsPalindrome(%q)=false", s)
		}
	}
} */
// 多线程测试
/* func TestIsPalindromePro(t *testing.T) {
	ch := make(chan string, 2)
	stop := make(chan string, 1)
	// 初始化一个随机数生成器
	seed := time.Now().UTC().UnixNano()
	t.Logf("random seed:%d", seed)
	num := rand.New(rand.NewSource(seed))

		// 1000		0.003s
		// 10000		0.012s
		// 100000		0.135s	0.129s	0.516s	0.456s	出错	受机器性能影响，耗费时间越来越久
		// 1000000		出错	出错	出错	出错	出错
		// 10000000	死机

	for i := 0; i < 100000; i++ {
		go func() {
			var number = &rand.Rand{}
			number = num
			s := randomStringMaker(number)
			if !IsPalindrome(s) {
				ch <- fmt.Sprintf("IsPalindrome(%q)=false", s)
			}
		}()
	}
	stop <- ""
loop:
	for {
		select {
		case s := <-ch:
			t.Error(s)
		case <-stop:
			break loop
		default:
		}
	}
} */

// 多线程，令牌限制最大并发数

func TestIsPalindromePro(t *testing.T) {
	// 令牌个数考虑到机器性能`，不宜超过1000
	token := make(chan string, 1000)
	ch := make(chan string, 2)
	stop := make(chan string, 1)
	// 初始化一个随机数生成器
	seed := time.Now().UTC().UnixNano()
	t.Logf("random seed:%d", seed)
	num := rand.New(rand.NewSource(seed))

	// 1000		0.003s
	// 10000		0.011s
	// 100000		0.111s	会出错
	// 1000000		出错	出错	出错	出错	出错
	// 10000000

	for i := 0; i < 100000; i++ {
		token <- " "
		go func() {
			var number = &rand.Rand{}
			number = num
			s := randomStringMaker(number)
			if !IsPalindrome(s) {
				ch <- fmt.Sprintf("IsPalindrome(%q)=false", s)
			}
		}()
		<-token
	}
	stop <- ""
loop:
	for {
		select {
		case s := <-ch:
			t.Error(s)
		case <-stop:
			break loop
		default:
		}
	}
}
