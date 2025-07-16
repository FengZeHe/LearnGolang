package main

import (
	"lbdemo/weightedroundrobin"
	"log"
	"sync"
)

func main() {
	//addrs := []string{
	//	"127.0.0.1:8080",
	//	"127.0.0.1:8081",
	//	"127.0.0.1:8082",
	//	"127.0.0.1:8083",
	//	"127.0.0.1:8084",
	//}
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

	wbalancer, wrrlbErr := weightedroundrobin.NewWeightedRoundRobinBalancer(map[string]int{
		"127.0.0.1:8080": 5,
		"127.0.0.1:8081": 3,
		"127.0.0.1:8082": 2,
	})
	if wrrlbErr != nil {
		log.Fatalf("NewWeightedRoundRobinBalancer err: %v", wrrlbErr)
	}

	totalReq := 100000
	// 总请求次数 10000
	record := make(map[string]int64)
	var mu sync.Mutex

	var wg sync.WaitGroup
	workers := 10
	perWorker := totalReq / workers
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < perWorker; j++ {
				//addr := wbalancer.SelectServerV2()
				addr := wbalancer.SelectServerV1()
				mu.Lock()
				record[addr]++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()

	log.Println("请求结果:")
	for k, v := range record {
		log.Println(k, v, float64(v)/float64(totalReq)*100, "%")
	}

}
