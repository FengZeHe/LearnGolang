package main

import (
	"fmt"
	"sync"
)

/*
sync.Map是在并发过程中使用的。
*/
func main() {
	m := sync.Map{}
	m.Store("cat", "Tom")
	m.Store("mouse", "Jerry")
	m.Store("person1", "Chan")

	val, ok := m.Load("mouse")
	if ok {
		fmt.Println(val)
	}
	//	store是往里面塞东西，Load是从里面读取数据m,
	// Load里面要写Key值

	//删除某个值为key的键值对
	m.Delete("person1")

	//遍历 sync.Map{}
	m.Range(func(key, value interface{}) bool {
		fmt.Println("-->", key, value)
		return true
	})
}
