package main

import "strings"

func basename1(s string) string { //移除路径部分和.后缀-----------------告诉我,哪里下标越界了
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for j := len(s) - 1; j >= 0; j-- {
		if s[j] == '.' {
			s = s[:j]
			break
		}
	}
	return s

}
func basename2(s string) string { //基于strings的lastindex做的basename
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
