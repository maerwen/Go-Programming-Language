package main

import "sync"

func main() {

}

var v = make(map[string]string)
var mu sync.Mutex
var rwmu sync.RWMutex
var once sync.Once

func load() {
	v = map[string]string{"name": "maerwen", "do you": "yes,we can"}
}
func primary(key string) string { //一次性初始化,并发不安全
	if v == nil {
		load()
	}
	return v[key]
}
func middle(key string) string { //使用sync.Mutex，并发安全，但初始化后效率较低，只能单线程访问
	mu.Lock()
	defer mu.Unlock()
	if v == nil {
		load()
	}
	return v[key]
}
func high(key string) string { //使用sync.RWMutex.在第一次初始化时使用读写锁，第二次使用互斥锁
	// 读写锁，是为了应对互斥锁在写操作之外的低效率情况
	// 读操作用读锁，写操作用互斥锁
	rwmu.RLock()
	if v != nil {
		value := v[key]
		rwmu.RUnlock()
		return value
	}
	rwmu.RUnlock()
	rwmu.Lock()
	if v == nil { //必须重新检查nil值
		load()
	}
	value := v[key]
	rwmu.Unlock()
	return value
}
func final(key string) string { //使用sync.Once，专门解决一次性初始化问题
	once.Do(load) //load为初始化函数
	return v[key]
}
