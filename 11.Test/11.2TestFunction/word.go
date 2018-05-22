package word

import (
	"unicode"
)

// 判断一个字符串是否是回文字符串

// 引发bug的版本
/* func IsPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
} */

// 重写后的版本
func IsPalindrome(s string) bool {
	var letters []rune
	// 忽略字母大小写以及非字母字符
	for _, j := range s {
		if unicode.IsLetter(j) {
			letters = append(letters, unicode.ToLower(j))
		}
	}
	for i := 0; i < len(letters)/2; i++ {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}
