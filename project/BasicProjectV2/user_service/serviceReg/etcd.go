package serviceReg

import (
	"context"
	"errors"
	"fmt"
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
	Client     *clientv3.Client
	Config     *EtcdConfig
	leaseID    clientv3.LeaseID
	serviceKey string
}

func NewEtcdClient(config *EtcdConfig) (*EtcdClient, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   config.Endpoints,
		DialTimeout: config.DialTimeout,
	})
	if err != nil {
		return nil, errors.New("create etcd client error:" + err.Error())
	}
	return &EtcdClient{Client: cli, Config: config}, nil
}

func (e *EtcdClient) RegisterService(serviceName, serviceAddr string) error {
	// 创建租约
	ctx, cancel := context.WithTimeout(context.Background(), e.Config.DialTimeout)
	resp, err := e.Client.Grant(ctx, e.Config.LeaseTTL)
	cancel()

	if err != nil {
		return errors.New("create lease error" + err.Error())
	}

	// 注册服务
	serviceKey := fmt.Sprintf("services/%s/%s", serviceName, serviceAddr)
	ctx, cancel = context.WithTimeout(context.Background(), e.Config.DialTimeout)
	_, err = e.Client.Put(ctx, serviceKey, serviceAddr, clientv3.WithLease(resp.ID))
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
	log.Printf("register service %s %s success,TTL: %d", serviceName, serviceAddr, e.Config.LeaseTTL)
	return nil
}

// 保持心跳 ＆ 自动续约
func (e *EtcdClient) keepAlive() {
	ctx := context.Background()
	al, err := e.Client.KeepAlive(ctx, e.leaseID)
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

// 服务下线
func (e *EtcdClient) UnregisterService() error {
	// todo 1. 取消租约
	if e.leaseID != 0 {
		ctx, cancel := context.WithTimeout(context.Background(), e.Config.DialTimeout)
		_, err := e.Client.Revoke(ctx, e.leaseID)
		cancel()
		if err != nil {
			return errors.New("revoke lease error:" + err.Error())
		}
	}

	// todo 2. 删除节点
	ctx, cancel := context.WithTimeout(context.Background(), e.Config.DialTimeout)
	_, err := e.Client.Delete(ctx, e.serviceKey)
	cancel()
	if err != nil {
		return errors.New("delete service error:" + err.Error())
	}

	log.Println("unregister service success")
	return nil
}
