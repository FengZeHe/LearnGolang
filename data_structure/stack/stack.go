package main

import (
	"fmt"
	"sync"
)

type Item interface {
}

type ItemStack struct {
	items []Item
	lock  sync.RWMutex
}

func NewStack() *ItemStack {
	s := &ItemStack{}
	s.items = []Item{}
	return s
}

func (s *ItemStack) Print() {
	fmt.Println(s.items)
}

func (s *ItemStack) Push(t Item) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.items = append(s.items, t)
}

func (s *ItemStack) Pop() Item {
	s.lock.Lock()
	defer s.lock.Unlock()
	if len(s.items) == 0 {
		return nil
	}

	item := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]
	return item
}
