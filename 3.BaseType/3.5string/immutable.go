package main

import (
	"fmt"
)

func immutable() { //不可变字符序列
	s := "hello"
	fmt.Printf("s:	\t%s\t\t地址:\t%q\n", s, &s)
	t1 := s
	s += " world"
	t2 := s
	fmt.Printf("s:	\t%s\t地址:\t%q\n", s, &s)
	fmt.Printf("t1:	\t%s\t\t地址:\t%q\n", t1, &t1)
	fmt.Printf("t2:	\t%s\t地址:\t%q\n", t2, &t2)

}
