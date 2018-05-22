package word

import "testing"

// 对于非GOROOT或者GOPATH路径下的test，必须使用cd命令进入当前目录才可执行go test命令进行测试
func TestPalindrome(t *testing.T) {
	if !IsPalindrome("deded") {
		t.Error(`IsPalindrome("deded")=false`)
	}
	if !IsPalindrome("katak") {
		t.Error(`IsPalindrome("katak")=false`)
	}
}
func TestNonPalindrome(t *testing.T) {
	if IsPalindrome("query") {
		t.Error(`IsPalindrome("query")=true`)
	}
}

// 法语字符识别
func TestFrenchPalindrome(t *testing.T) {
	if !IsPalindrome("été") {
		t.Error(`IsPalindrome("été")=false`)
	}
}

// 语句识别
func TestCanalPalindrome(t *testing.T) {
	if !IsPalindrome("ok,wow.ko!") {
		t.Error(`IsPalindrome("ok,wow.ko!")=false`)
	}
}

// 更全面的测试用例
func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"deded", true},
		{"katak", true},
		{"query", false},
		{"été", true},
		{"ok,wow.ko!", false},
		{"lalala", true},
	}
	for _, ele := range tests {
		if got := IsPalindrome(ele.input); got != ele.want {
			// %q像源代码一样带双引号的输出
			t.Errorf("IsPalindrome(%q)=%v", ele.input, got)
			// t.Fatalf("IsPalindrome(%q)=%v", ele.input, got)
		}
	}
}
