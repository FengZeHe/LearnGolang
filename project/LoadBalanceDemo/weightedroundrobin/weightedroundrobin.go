package weightedroundrobin

import (
	"errors"
	"math"
	"sync"
	"sync/atomic"
)

type WeightedServer struct {
	Addr          string
	Weight        int // 权重
	CurrentWeight int
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
		addrs: make([]*WeightedServer, 0, len(servers)),
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

// 加权轮询
func (wrrb *WeightedRoundRobinBalancer) SelectServer() string {
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

func (wrrb *WeightedRoundRobinBalancer) SelectServerV1() string {
	wrrb.mu.Lock()
	defer wrrb.mu.Unlock()

	var (
		selectedServer *WeightedServer
		totalWeight    int
	)

	// 1. 计算总权重，并累加当前权重
	for _, server := range wrrb.addrs {
		totalWeight += server.Weight
		server.CurrentWeight += server.Weight
	}

	// 2. 选择当前权重最大的服务器
	for _, server := range wrrb.addrs {
		if selectedServer == nil || server.CurrentWeight > selectedServer.CurrentWeight {
			selectedServer = server
		}
	}

	// 3. 选中服务器的当前权重减去总权重
	if selectedServer != nil {
		selectedServer.CurrentWeight -= totalWeight
		return selectedServer.Addr
	}

	return ""
}

// 平滑加权轮询
func (wrrb *WeightedRoundRobinBalancer) SelectServerV2() string {
	wrrb.mu.Lock()
	defer wrrb.mu.Unlock()

	if len(wrrb.addrs) == 0 {
		return ""
	}

	totalWeight := 0
	var bestServer *WeightedServer
	maxCurrentWeight := math.MinInt64

	// 计算总权重
	for _, server := range wrrb.addrs {
		totalWeight += server.Weight
		server.CurrentWeight += server.Weight

		// 找出当前权重最大的服务器
		if server.CurrentWeight > maxCurrentWeight {
			maxCurrentWeight = server.CurrentWeight
			bestServer = server
		}
	}

	if bestServer != nil {
		bestServer.CurrentWeight -= totalWeight
		return bestServer.Addr
	}

	return wrrb.addrs[0].Addr

}
