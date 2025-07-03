package ioc

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

type EtcdConfig struct {
	Endpoints   []string      // etcd地址
	DialTimeout time.Duration // 连接超时时间
	LeaseTTL    int64         // 租约时间 s
}

// etcd client
type EtcdClient struct {
	client     *clientv3.Client
	config     *EtcdConfig
	leaseID    clientv3.LeaseID
	serviceKey string
}

func NewEtcdConfig() *EtcdConfig {
	return &EtcdConfig{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
		LeaseTTL:    10,
	}
}

func NewEtcdClient(config *EtcdConfig) (*EtcdClient, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   config.Endpoints,
		DialTimeout: config.DialTimeout,
	})
	if err != nil {
		return nil, errors.New("create etcd client error:" + err.Error())
	}
	return &EtcdClient{client: cli, config: config}, nil
}

func (e *EtcdClient) RegisterService(serviceName, serviceAddr string) error {
	// 创建租约
	ctx, cancel := context.WithTimeout(context.Background(), e.config.DialTimeout)
	resp, err := e.client.Grant(ctx, e.config.LeaseTTL)
	cancel()

	if err != nil {
		return errors.New("create lease error" + err.Error())
	}

	// 注册服务
	serviceKey := fmt.Sprintf("services/%s/%s", serviceName, serviceAddr)
	ctx, cancel = context.WithTimeout(context.Background(), e.config.DialTimeout)
	_, err = e.client.Put(ctx, serviceKey, serviceAddr, clientv3.WithLease(resp.ID))
	cancel()
	if err != nil {
		return errors.New("register service error:" + err.Error())
	}

	// 保存租约ID和 serviceKey
	e.leaseID = resp.ID
	e.serviceKey = serviceKey

	// 保持心跳
	go e.keepAlive()
	log.Println("serviceKey", serviceKey)
	log.Printf("register service %s %s success,TTL: %d", serviceName, serviceAddr, e.config.LeaseTTL)
	return nil
}

// 保持心跳 ＆ 自动续约
func (e *EtcdClient) keepAlive() {
	ctx := context.Background()
	al, err := e.client.KeepAlive(ctx, e.leaseID)
	if err != nil {
		log.Println("start keep alive error:" + err.Error())
		return
	}

	for {
		select {
		case k, ok := <-al:
			if !ok {
				log.Println("keep alive closed,重新注册服务")
				return
			}
			log.Printf("keep alive TTL %d", k.TTL)
		}
	}
}
