package weightedroundrobin

import (
	"errors"
	"sync"
	"sync/atomic"
)

type WeightedServer struct {
	Addr          string
	Weight        int // 权重
	CurrentWeight int
}

type Server struct {
}

type WeightedRoundRobinBalancer struct {
	addrs []*WeightedServer
	mu    sync.RWMutex
	index uint64
}

func NewWeightedRoundRobinBalancer(servers map[string]int) (*WeightedRoundRobinBalancer, error) {
	if len(servers) == 0 {
		return nil, errors.New("servers is empty")
	}
	wrrb := &WeightedRoundRobinBalancer{
		addrs: make([]*WeightedServer, len(servers)),
	}

	// 服务器列表
	for addr, weight := range servers {
		if weight <= 0 {
			weight = 1 // 权重最小是1
		}
		wrrb.addrs = append(wrrb.addrs, &WeightedServer{Addr: addr, Weight: weight, CurrentWeight: 0})
	}
	return wrrb, nil
}

func (wrrb *WeightedRoundRobinBalancer) Next() string {
	wrrb.mu.RLock()
	defer wrrb.mu.RUnlock()

	return ""
}

// 加权轮询
func (wrrb *WeightedRoundRobinBalancer) selectServerV1() string {
	totalWeight := 0
	for _, server := range wrrb.addrs {
		totalWeight += server.Weight
	}

	// 递增索引
	currentIndex := atomic.AddUint64(&wrrb.index, 1) - 1

	// 根据索引计算应选择的服务器
	weightIndex := int(currentIndex % uint64(totalWeight))

	// 找到对应权重区间
	for _, server := range wrrb.addrs {
		weightIndex -= server.Weight
		if weightIndex < 0 {
			return server.Addr
		}
	}
	return wrrb.addrs[0].Addr
}

// 平滑加权轮询
func (wrrb *WeightedRoundRobinBalancer) SelectServerV2() string {
	totalWeight := 0
	bestServer := wrrb.addrs[0]
	maxCurrentWeight := wrrb.addrs[0].CurrentWeight // 先使用第一台服务器的当前权重

	for _, server := range wrrb.addrs {
		totalWeight += server.Weight // 总权重
		server.CurrentWeight += server.Weight

		if server.CurrentWeight > maxCurrentWeight {
			maxCurrentWeight = server.CurrentWeight
			bestServer = server
		}
	}
	bestServer.CurrentWeight += maxCurrentWeight
	return bestServer.Addr
}
