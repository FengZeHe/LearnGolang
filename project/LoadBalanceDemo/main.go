package main

import (
	"fmt"
	"lbdemo/consistenthash"
	"lbdemo/hash"
	"lbdemo/random"
	"lbdemo/weightedrandom"
	"log"
)

func main() {
	// todo 轮询
	addrs := []string{
		"127.0.0.1:8080",
		"127.0.0.1:8081",
		"127.0.0.1:8082",
		"127.0.0.1:8083",
		"127.0.0.1:8084",
	}
	//
	//rrlb, rrlbErr := roundrobin.NewRoundrobinLoadBalancer(addrs)
	//if rrlbErr != nil {
	//	log.Fatalf("NewRoundrobinLoadBalancer err: %v", rrlbErr)
	//}
	//
	//for i := 0; i < 10; i++ {
	//	addr := rrlb.Next()
	//	fmt.Println("轮询访问地址->", addr)
	//}

	// todo 加权轮询
	//wbalancer, wrrlbErr := weightedroundrobin.NewWeightedRoundRobinBalancer(map[string]int{
	//	"A": 5,
	//	"B": 3,
	//	"C": 2,
	//})
	//if wrrlbErr != nil {
	//	log.Fatalf("NewWeightedRoundRobinBalancer err: %v", wrrlbErr)
	//}
	//
	//totalReq := 100000
	//// 总请求次数 10000
	//record := make(map[string]int64)
	//var mu sync.Mutex
	//
	//var wg sync.WaitGroup
	//workers := 10
	//perWorker := totalReq / workers
	//wg.Add(workers)
	//for i := 0; i < workers; i++ {
	//	go func() {
	//		defer wg.Done()
	//		for j := 0; j < perWorker; j++ {
	//			//addr := wbalancer.SelectServerV2()
	//			addr := wbalancer.SelectServerV1()
	//			fmt.Printf(addr)
	//			mu.Lock()
	//			record[addr]++
	//			mu.Unlock()
	//		}
	//	}()
	//}
	//wg.Wait()
	//
	//log.Println("请求结果:")
	//for k, v := range record {
	//	log.Println(k, v, float64(v)/float64(totalReq)*100, "%")
	//}

	randomAddrs := []string{
		"127.0.0.1:8080",
		"127.0.0.1:8081",
		"127.0.0.1:8082",
		"127.0.0.1:8083",
		"127.0.0.1:8084",
	}

	rlb := random.NewRandomLoadBalancer(randomAddrs)
	for i := 0; i < 10; i++ {
		server := rlb.SelectServer()
		log.Println("随机算法选出：", server)
	}

	wrlbAddr := []weightedrandom.WRServer{
		{"A", 5},
		{"B", 3},
		{"C", 1},
	}
	wrlb := weightedrandom.NewWeightedRandomLoadBalancer(wrlbAddr)

	wrSelectCount := make(map[string]int)
	for i := 1; i < 20; i++ {
		server := wrlb.SelectServer()
		wrSelectCount[server]++
	}
	log.Println("随机加权选中结果统计:", wrSelectCount)

	/*
		哈希/一致性哈希
	*/
	// 测试客户端IP列表
	clients := []string{
		"192.168.1.1",
		"192.168.1.2",
		"192.168.1.3",
		"192.168.1.4",
		"192.168.1.5",
		"192.168.1.6",
		"192.168.1.7",
		"192.168.1.8",
		"192.168.1.9",
		"192.168.1.10",
	}
	hlbCount := make(map[string]int)
	hlb := hash.NewHashLoadBalancer(addrs)
	for _, client := range clients {
		server := hlb.SelectServer(client)
		fmt.Printf("hash client: %s -> server: %s \n", client, server)
		hlbCount[server]++
	}
	log.Println("hlbCount->", hlbCount)

	chlb := consistenthash.NewConsistentHashLoadBalancer(100, addrs)
	for _, client := range clients {
		server := chlb.SelectServer(client)
		fmt.Printf("cosistent hash client: %s -> server: %s \n", client, server)
	}

}
