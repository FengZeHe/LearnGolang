package main

import (
	"fmt"
	"lbdemo/roundrobin"
	"log"
)

func main() {
	addrs := []string{
		"127.0.0.1:8080",
		"127.0.0.1:8081",
		"127.0.0.1:8082",
		"127.0.0.1:8083",
		"127.0.0.1:8084",
	}

	rrlb, rrlbErr := roundrobin.NewRoundrobinLoadBalancer(addrs)
	if rrlbErr != nil {
		log.Fatalf("NewRoundrobinLoadBalancer err: %v", rrlbErr)
	}

	for i := 0; i < 10; i++ {
		addr := rrlb.Next()
		fmt.Println("轮询访问地址->", addr)
	}

	//wbalancer, wrrlbErr := weightedroundrobin.NewWeightedRoundRobinBalancer(map[string]int{
	//	"127.0.0.1:8080": 5,
	//	"127.0.0.1:8081": 3,
	//	"127.0.0.1:8082": 2,
	//})
	//if wrrlbErr != nil {
	//	log.Fatalf("NewWeightedRoundRobinBalancer err: %v", wrrlbErr)
	//}
	//for i := 0; i < 10; i++ {
	//	addr := wbalancer.Next()
	//	log.Println("加权轮询访问地址->", addr)
	//}

}
