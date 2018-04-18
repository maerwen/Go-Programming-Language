package main

func nonempty1(x []string) []string { //返回已有的slice并剔除空字符串
	//创建一个新索引,每次出现一个非空字符串,就把它自增,同时对应设置为该元素的索引,最后返回slice
	i := 0
	for _, s := range x {
		if s != "" {
			x[i] = s
			i++
		}
	}
	return x[:i]
}
func nonempty2(x []string) []string { //返回已有的slice并剔除空字符串
	//利用append来实现
	r := x[:0]
	for _, s := range x {
		if s != "" {
			r = append(r, s)
		}
	}
	return r
}
func remove1(s []int, index int) []int { //从slice中移除某个索引位置处的元素,已有顺序不变
	copy(s[index:], s[index+1:])
	return s[:len(s)-1]
}
func remove2(s []int, index int) []int { //从slice中移除某个索引位置处的元素,顺序无要求
	s[index] = s[len(s)-1]
	return s[:len(s)-1]
}
