package hash

import "hash/crc32"

type HashLoadBalancer struct {
	servers []string
}

func NewHashLoadBalancer(servers []string) *HashLoadBalancer {
	return &HashLoadBalancer{servers: servers}
}

func (h *HashLoadBalancer) SelectServer(clientIP string) string {
	if len(h.servers) == 0 {
		return ""
	}
	hash := crc32.ChecksumIEEE([]byte(clientIP))
	// 对服务器数量取模
	index := int(hash) % len(h.servers)
	return h.servers[index]
}
