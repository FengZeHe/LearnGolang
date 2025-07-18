package weightedrandom

import (
	"math/rand"
	"time"
)

type WeightedRandomLoadBalancer struct {
	servers []WRServer
	rng     *rand.Rand
}

type WRServer struct {
	Addr   string
	Weight int
}

func NewWeightedRandomLoadBalancer(servers []WRServer) *WeightedRandomLoadBalancer {
	src := rand.NewSource(time.Now().UnixNano())
	return &WeightedRandomLoadBalancer{servers: servers, rng: rand.New(src)}
}

func (w *WeightedRandomLoadBalancer) SelectServer() string {
	if len(w.servers) == 0 {
		return ""
	}
	totalWeight := 0
	for _, server := range w.servers {
		totalWeight += server.Weight
	}

	if totalWeight == 0 {
		return ""
	}

	randomValue := w.rng.Intn(totalWeight) + 1

	currentWeight := 0
	for _, server := range w.servers {
		currentWeight += server.Weight
		if randomValue <= currentWeight {
			return server.Addr
		}
	}
	return ""
}
