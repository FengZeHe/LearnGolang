package serviceDiscovery

import (
	"context"
	"errors"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/attributes"
	"google.golang.org/grpc/resolver"
	"log"
	"strings"
	"time"
)

const (
	Scheme = "etcd"
)

type etcdResolver struct {
	target     resolver.Target
	conn       resolver.ClientConn
	etcdClient *clientv3.Client
	ctx        context.Context
	cancel     context.CancelFunc
}

// 初始化etcdResolver
func newEtcdResolver(target resolver.Target, conn resolver.ClientConn, etcdClient *clientv3.Client) resolver.Resolver {
	ctx, cancel := context.WithCancel(context.Background())
	r := &etcdResolver{
		target:     target,
		conn:       conn,
		etcdClient: etcdClient,
		ctx:        ctx,
		cancel:     cancel,
	}
	r.ResolveNow(resolver.ResolveNowOptions{})
	return r
}

func (r *etcdResolver) ResolveNow(options resolver.ResolveNowOptions) {
	go func() {
		if err := r.watchServices(); err != nil {
			log.Printf("Failed to watch services: %v", err)
		}
	}()
}

// Close 关闭解析器
func (r *etcdResolver) Close() {
	r.cancel()
}

// watchServices 监听 etcd 中的服务变化
func (r *etcdResolver) watchServices() error {
	serviceName := r.target.URL.Path

	if len(serviceName) > 0 && serviceName[0] == '/' {
		serviceName = serviceName[1:]
	}

	if serviceName == "" {
		return fmt.Errorf("empty service name in target: %+v", r.target)
	}

	serviceKeyPrefix := fmt.Sprintf("services/%s/", serviceName)

	// 首次获取服务列表
	resp, err := r.etcdClient.Get(r.ctx, serviceKeyPrefix, clientv3.WithPrefix())
	if err != nil {
		return err
	}

	addrs := make([]resolver.Address, 0)
	for _, kv := range resp.Kvs {
		keyParts := strings.Split(string(kv.Key), "/")
		if len(keyParts) < 3 {
			log.Printf("Invalid key format: %s", string(kv.Key))
			continue
		}

		// 最后一部分是服务地址
		addr := keyParts[len(keyParts)-1]
		if addr == "" {
			log.Printf("Empty address in key: %s", string(kv.Key))
			continue
		}

		addrs = append(addrs, resolver.Address{
			Addr:       addr,
			Attributes: attributes.New("service", serviceName),
			ServerName: serviceName,
		})
	}

	log.Printf("Resolved addresses service==> %s Addr ==> %v", serviceName, addrs)

	// 更新客户端连接状态
	if len(addrs) == 0 {
		return fmt.Errorf("no addresses found for service %s", serviceName)
	}

	if err := r.conn.UpdateState(resolver.State{Addresses: addrs}); err != nil {
		return err
	}

	// 监听服务变化
	rch := r.etcdClient.Watch(r.ctx, serviceKeyPrefix, clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			// 从键中提取服务地址
			keyParts := strings.Split(string(ev.Kv.Key), "/")
			if len(keyParts) < 3 {
				log.Printf("Invalid key format in watch: %s", string(ev.Kv.Key))
				continue
			}

			addr := keyParts[len(keyParts)-1]
			if addr == "" {
				log.Printf("Empty address in watch key: %s", string(ev.Kv.Key))
				continue
			}

			switch ev.Type {
			case clientv3.EventTypePut:
				if !containsAddr(addrs, addr) {
					addrs = append(addrs, resolver.Address{
						Addr:       addr,
						Attributes: attributes.New("service", serviceName),
					})
					log.Printf("Added new address: %s", addr)
				}
			case clientv3.EventTypeDelete:
				addrs = removeAddr(addrs, addr)
				log.Printf("Removed address: %s", addr)
			}
		}

		// 更新客户端连接状态
		if err := r.conn.UpdateState(resolver.State{Addresses: addrs}); err != nil {
			return err
		}
	}

	return nil
}

// containsAddr 检查地址是否已存在
func containsAddr(addrs []resolver.Address, addr string) bool {
	for _, a := range addrs {
		if a.Addr == addr {
			return true
		}
	}
	return false
}

func removeAddr(addrs []resolver.Address, addr string) []resolver.Address {
	for i, a := range addrs {
		if a.Addr == addr {
			return append(addrs[:i], addrs[i+1:]...)
		}
	}
	return addrs
}

type EtcdResolverBuilder struct {
	EtcdClient *clientv3.Client
}

// Build 创建解析器实例
func (b *EtcdResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	return newEtcdResolver(target, cc, b.EtcdClient), nil
}

func (b *EtcdResolverBuilder) Scheme() string {
	return Scheme
}

// initEtcdClient 初始化 etcd 客户端
func InitEtcdClient(endpoints []string) (*clientv3.Client, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return nil, errors.New("create etcd client error:" + err.Error())
	}
	return cli, nil
}

func InitResolver(etcdClient *clientv3.Client) {
	builder := &EtcdResolverBuilder{
		EtcdClient: etcdClient,
	}
	resolver.Register(builder)
}
