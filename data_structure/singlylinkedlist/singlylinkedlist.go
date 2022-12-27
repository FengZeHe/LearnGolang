package main

import (
	"fmt"
)

// 定义结点
type Node struct {
	Value int
	Next  *Node
}

// 初始化头节点
var head = new(Node)

func addNode(t *Node, v int) int {
	// 如果这里就是尾部
	if head == nil {
		t = &Node{v, nil}
		head = t
		return 0
	}

	if v == t.Value {
		fmt.Println("节点已经存在", v)
		return -1
	}

	// 如果当前结点的下一个结点为nil
	if t.Next == nil {
		t.Next = &Node{v, nil}
		return -2
	}
	// 如果当前节点下一个节点不为空
	return addNode(t.Next, v)

}

// 遍历链表
func traverse(t *Node) {
	if t == nil {
		fmt.Println("-> 空链表！")
		return
	}

	//一直遍历到尾部
	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}

// 查找节点
func lookupNode(t *Node, v int) bool {
	if head == nil {
		t = &Node{v, nil}
		head = t
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

// 获取链表长度
func size(t *Node) int {
	if t == nil {
		fmt.Println("空链表")
		return 0
	}

	i := 0
	for t != nil {
		i++
		t = t.Next
	}
	return i
}

func main() {
	fmt.Println(head)
	head = nil
	traverse(head)
	addNode(head, 1)
	addNode(head, -1)
	traverse(head)

	addNode(head, 7)
	addNode(head, 8)
	addNode(head, 9)
	traverse(head)

	//查找存在节点
	if lookupNode(head, 8) {
		fmt.Println("节点存在")
	} else {
		fmt.Println("节点不存在")
	}
}
