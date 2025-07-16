package roundrobin

import (
	"errors"
	"sync"
)

type RoundrobinLoadBalancer struct {
	addrs []string // 服务器地址列表
	index int
	mu    sync.RWMutex
}

func NewRoundrobinLoadBalancer(addrs []string) (*RoundrobinLoadBalancer, error) {
	if len(addrs) == 0 {
		return nil, errors.New("addr list is empty ")
	}
	return &RoundrobinLoadBalancer{addrs: addrs, index: 0}, nil
}

// 获取下一个访问的地址
func (lb *RoundrobinLoadBalancer) Next() (addr string) {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	currentAddr := lb.addrs[lb.index]
	// 更新索引
	lb.index = (lb.index + 1) % len(lb.addrs)
	return currentAddr
}

// 添加新地址
func (lb *RoundrobinLoadBalancer) Add(addr string) {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	for _, v := range lb.addrs {
		if v == addr {
			return
		}
	}
	lb.addrs = append(lb.addrs, addr)
}

// 删除无效地址
func (lb *RoundrobinLoadBalancer) Remove(addr string) {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	for i, v := range lb.addrs {
		if v == addr {
			lb.addrs = append(lb.addrs[:i], lb.addrs[i+1:]...)
		}

		// 调整索引  删掉的是结尾，索引越界
		if lb.index >= len(lb.addrs) {
			lb.index = 0
		}
		break
	}
}
