package main

// strings 包底层实现
func contains(s, subStr string) bool {
	for i := 0; i < len(s); i++ {
		if hasSuffix(s[i:], subStr) {
			return true
		}
	}
	return false
}
func len(s string) int {
	n := 0
	// for _, _ = range s { //如果返回值全部弃用,则不能用短变量声明,直接写赋值号
	// 	n++
	// }
	for range s { //最简洁写法
		n++
	}
	// return 111
	return n
}
func hasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}
func hasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}
