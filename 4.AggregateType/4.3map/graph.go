package main

var graph = make(map[string]map[string]bool)

//建立一个字符串到字符串集合的映射
func addEdge(from, to string) {
	edge := graph[from]
	if edge == nil {
		edge = make(map[string]bool)
	}
	edge[to] = true
}
func hasEdge(from, to string) bool {
	return graph[from][to]
}
