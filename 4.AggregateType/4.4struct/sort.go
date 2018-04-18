package main

// 往二叉树里面存入一个数据
func add(t *tree, i int) *tree {
	if t == nil {
		t = new(tree)
		t.value = i
		return t
	}
	if t.value > i {
		t.left = add(t.left, i)
	} else {
		t.right = add(t.right, i)
	}
	return t
}

// 把tree里面的存储的value按照顺序追加到values里面,返回slice
func appendValues(values []int, t *tree) []int {
	if t == nil {
		return values
	}
	values = appendValues(values, t.left)
	values = append(values, t.value)
	values = appendValues(values, t.right)
	return values
}
func sort(values []int) []int {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	return appendValues(values[:0], root)

}
