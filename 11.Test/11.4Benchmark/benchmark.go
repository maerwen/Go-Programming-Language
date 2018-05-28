package benchmark

import "unicode"

func IsPalindrome1(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
func IsPalindrome2(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func IsPalindrome3(s string) bool {
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
func IsPalindrome4(s string) bool {
	var letters = make([]rune, len(s))
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
