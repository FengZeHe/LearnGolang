package consistenthash

import (
	"crypto/sha1"
	"fmt"
	"sort"
)

type ConsistentHashLoadBalancer struct {
	replicas int            //每个服务器的虚拟节点数
	keys     []uint32       // 虚拟节点的哈希值
	hashMap  map[uint32]int // 虚拟节点哈希值 -> 服务器索引
	servers  []string       // 服务器列表
}

func NewConsistentHashLoadBalancer(replicas int, servers []string) *ConsistentHashLoadBalancer {
	chlb := &ConsistentHashLoadBalancer{
		replicas: replicas,
		keys:     []uint32{},
		hashMap:  make(map[uint32]int),
		servers:  servers,
	}

	// 初始化虚拟节点
	for i, server := range servers {
		for j := 0; j < replicas; j++ {
			virtualNode := fmt.Sprintf("%s-%d", server, j) // 服务器-编号
			// 计算虚拟节点的哈希值
			hash := chlb.calcHash(virtualNode)
			chlb.keys = append(chlb.keys, hash)
			chlb.hashMap[hash] = i
		}
	}
	sort.Slice(chlb.keys, func(i, j int) bool {
		return chlb.keys[i] < chlb.keys[j]
	})
	return chlb
}

func (chlb *ConsistentHashLoadBalancer) calcHash(key string) uint32 {
	hash := sha1.Sum([]byte(key))
	// 取前4字节作为uint32值
	return uint32(hash[0])<<24 | uint32(hash[1])<<16 | uint32(hash[2])<<8 | uint32(hash[3])
}

func (chlb *ConsistentHashLoadBalancer) SelectServer(clientIP string) string {
	if len(chlb.servers) == 0 {
		return ""
	}
	hash := chlb.calcHash(clientIP)

	idx := sort.Search(len(chlb.keys), func(i int) bool {
		return chlb.keys[i] >= hash
	})
	if idx == len(chlb.keys) {
		idx = 0
	}

	serverIdx := chlb.hashMap[chlb.keys[idx]]
	return chlb.servers[serverIdx]
}
