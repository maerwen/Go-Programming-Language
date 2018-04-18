package main

import "bytes"

func comm(s string) string { //向十进制非负整数的字符串中从右侧开始每三位插入一个逗号
	// 递归做法
	/*
		n := len(s)
		if n <= 3 {
			return s
		}
		return comm(s[:n-3]) + "," + s[n-3:]
	*/
	// 反转
	n := 0
	var str, st bytes.Buffer
	for j := len(s) - 1; j >= 0; j-- {
		n++
		str.WriteByte(s[j])
		if n == 3 {
			str.WriteByte(',')
			n = 0
		}
	}
	s = str.String()
	for i := len(s) - 1; i >= 0; i-- {
		st.WriteByte(s[i])
	}
	return st.String()
}
