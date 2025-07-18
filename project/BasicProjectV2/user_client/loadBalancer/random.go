package loadBalancer

import (
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
	"math/rand"
	"sync"
)

type RandomPicker struct {
	subConns []balancer.SubConn
	mu       sync.Mutex
}

type RandomPickerBuilder struct{}

func (rp *RandomPicker) Pick(info balancer.PickInfo) (balancer.PickResult, error) {
	rp.mu.Lock()
	defer rp.mu.Unlock()

	if len(rp.subConns) == 0 {
		return balancer.PickResult{}, balancer.ErrNoSubConnAvailable
	}

	index := rand.Intn(len(rp.subConns))
	return balancer.PickResult{SubConn: rp.subConns[index]}, nil
}

func (rp *RandomPickerBuilder) Build(info base.PickerBuildInfo) balancer.Picker {
	var subConns []balancer.SubConn

	for subConn := range info.ReadySCs {
		subConns = append(subConns, subConn)
	}

	return &RandomPicker{
		subConns: subConns,
	}
}
