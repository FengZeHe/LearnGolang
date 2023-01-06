package main

import (
	"fmt"
)

/*
使用链表实现队列
*/

type queueNode struct {
	data interface{} //放数据
	next *queueNode  // 放指针

}

type queueList struct {
	length int        //存储长度
	front  *queueNode // 指向对头
	rear   *queueNode //指向队尾
}

// 链表初始化
func initQueue() *queueList {
	L := &queueList{
		length: 0,
		front:  nil,
		rear:   nil,
	}
	return L
}

// 判断队列是否为空
func (queue queueList) isNull() bool {
	return queue.length == 0
}

// 链表入队
func (queue *queueList) push(val interface{}) {
	node := &queueNode{
		data: val,
		next: nil,
	}

	//	如果是空队列
	if queue.isNull() {
		queue.front = node
		queue.rear = node
		queue.length++
		return
	}

	queue.rear.next = node
	queue.rear = queue.rear.next
	queue.length++
}

// 链表出队
func (queue *queueList) pop() {
	//	如果队列为空
	if queue.isNull() {
		fmt.Println("queue empty")
		return
	}

	//	链表中只有一个节点
	if queue.length == 1 {
		queue.front = nil
		queue.rear = nil
		queue.length--
		return
	}

	queue.front = queue.front.next
	queue.length--
}

// 遍历队列
func (queue *queueList) Traverse() (arr []interface{}) {
	pre := queue.front
	for i := 0; i < queue.length; i++ {
		arr = append(arr, pre.data, "-->")
		pre = pre.next
	}
	return
}

func main() {
	data := []interface{}{
		"holy",
		20,
		30,
	}

	L := initQueue()
	for i := range data {
		L.push(data[i])
		fmt.Println(L.Traverse())
	}

	L.pop()
	L.Traverse()

	L.pop()
	fmt.Println(L.Traverse())

	L.push("999")
	fmt.Println(L.Traverse())
}
