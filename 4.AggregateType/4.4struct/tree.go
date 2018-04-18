package main

// 利用二叉树来实现排序
type tree struct {
	value       int
	left, right *tree
}
