package random

import (
	"math/rand"
	"time"
)

type RandomLoadBalancer struct {
	servers []string
	rng     *rand.Rand
}

func NewRandomLoadBalancer(servers []string) *RandomLoadBalancer {
	// 初始化seed
	s := rand.NewSource(time.Now().UnixNano())
	return &RandomLoadBalancer{servers: servers, rng: rand.New(s)}
}

// 随机算法
func (r *RandomLoadBalancer) SelectServer() string {
	if len(r.servers) == 0 {
		return ""
	}

	index := r.rng.Intn(len(r.servers))
	return r.servers[index]
}
