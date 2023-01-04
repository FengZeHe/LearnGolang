package main

import (
	"errors"
	"fmt"
)

/*
使用数组实现队列
*/

type queue struct {
	maxsize    int
	buf        []int
	head, tail int
	overflow   bool //判断是否溢出
}

// 创建队列
func NewQueue(size int) *queue {
	return &queue{
		buf:     make([]int, size+1),
		maxsize: size,
	}
}

// 入队操作
func (q *queue) Push(elem int) (err error) {
	//如果队列已经满了
	if q.overflow {
		if q.tail-1 == q.head || q.tail == q.maxsize {
			err = errors.New("queue full ")
			return err
		}
	}

	q.tail++
	q.buf[q.tail-1] = elem
	//	队尾指针向后移一位
	if q.tail == q.maxsize {
		q.tail = 0
		q.overflow = true
	}
	return nil
}

// 出队操作
func (q *queue) Pop() (elem int, err error) {
	//	如果队列是空的
	if q.overflow != true {
		if q.tail == q.head {
			err = errors.New("queue empty")
			return -1, err
		}
	}
	//	出队操作
	elem = q.buf[q.head]
	q.head++

	if q.head == q.maxsize {
		q.head = 0
		q.overflow = false
	}
	return elem, nil
}

// 查看队列元素
func (q *queue) Show() {
	data := make([]int, 0, q.maxsize)
	if q.overflow {
		data = append(data, q.buf[:q.tail]...)
		data = append(data, q.buf[q.head:q.maxsize]...)
	} else {
		data = append(data, q.buf[q.head:q.tail]...)
	}

	fmt.Println("queue =>", data)
}

func main() {
	q := NewQueue(5)
	q.Push(1)
	q.Push(2)
	q.Push(3)

	q.Show()

	q.Push(4)
	q.Push(5)

	q.Show()

	q.Pop()
	q.Pop()
	q.Show()
}
