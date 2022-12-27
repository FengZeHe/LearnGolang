package main

import (
	"fmt"
)

type Node struct {
	Value    int
	Previous *Node
	Next     *Node
}

var head = new(Node)

// 添加节点
func addNode(t *Node, v int) int {
	//如果这是唯一的节点，那么前驱、后继指针都赋值为nil
	if head == nil {
		t = &Node{v, nil, nil}
		head = t
		return 0
	}

	if v == t.Value {
		fmt.Println("节点已经存在", v)
		return -1
	}

	//如果当前节点的下一个节点为空
	if t.Next == nil {
		//每个节点还要维护前驱节点指针
		temp := t
		t.Next = &Node{v, temp, nil}
		return -2
	}
	return addNode(t.Next, v)
}

// 遍历链表
func traverse(t *Node) {
	if t == nil {
		fmt.Println("-> 空链表")
		return
	}

	for t != nil {
		fmt.Printf("%d ->", t.Value)
		t = t.Next
	}
	fmt.Println()
}

// 反向遍历链表
func reverse(t *Node) {
	if t == nil {
		fmt.Println("-> 空链表")
		return
	}

	temp := t
	for t != nil {
		temp = t
		t = t.Next
	}

	for temp.Previous != nil {
		fmt.Printf("%d ->", temp.Value)
		temp = temp.Previous
	}

	fmt.Printf("%d ->", temp.Value)
	fmt.Println()
}

// 获取长度
func getsize(t *Node) int {
	if t == nil {
		fmt.Println("-> 空链表")
		return 0
	}

	n := 0
	for t != nil {
		n++
		t = t.Next
	}
	return n
}

// 查找节点
func lookupNode(t *Node, v int) bool {
	if head == nil {
		return false
	}

	if v == t.Value {
		return true
	}

	if t.Next == nil {
		return false
	}
	return lookupNode(t.Next, v)
}

func main() {
	fmt.Print(head)
	head = nil

	traverse(head)
	addNode(head, 1)
	traverse(head)

	addNode(head, 20)
	traverse(head)
	reverse(head)

	if lookupNode(head, 1) {
		fmt.Println("该节点存在")
	} else {
		fmt.Println("该节点不存在")
	}

	fmt.Println("size = ", getsize(head))

}
