package main

import (
	"context"
	"fmt"
	pb "github.com/basicprojectv2/proto/user_service"
	"github.com/basicprojectv2/user_client/interceptors/jwt"
	lg "github.com/basicprojectv2/user_client/interceptors/log"
	"github.com/basicprojectv2/user_client/loadBalancer"
	et "github.com/basicprojectv2/user_client/serviceDiscovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())

	// 注册负载均衡器
	balancer.Register(base.NewBalancerBuilder(
		"my_random",
		&loadBalancer.RandomPickerBuilder{}, //负载均衡构建
		base.Config{HealthCheck: true},      // 启用健康检查
	))
}

func main() {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3ODY5NjU0MDAsIlVzZXJJZCI6IjE5MzQ4ODI1NjQ3MDM1MjI4MTYifQ.UvpMjVj5AsYemY0X9HXvDr-BOlLGgniklZb45_qg8_g"
	ji := jwt.NewJWTInterceptor(token)
	li := lg.NewLogInterceptor()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// todo 初始化etcd客户端
	etcdEndpoints := []string{"http://localhost:2379"}
	etcdClient, err := et.InitEtcdClient(etcdEndpoints)
	if err != nil {
		log.Fatal("init etcd client error:", err)
	}
	defer etcdClient.Close()

	et.InitResolver(etcdClient)

	serviceName := "user_service" // 替换为实际的服务名
	targetURI := fmt.Sprintf("%s:///%s", et.Scheme, serviceName)
	log.Printf("Target URI: %s", targetURI)
	conn, err := grpc.Dial(
		targetURI,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithResolvers(&et.EtcdResolverBuilder{EtcdClient: etcdClient}),
		grpc.WithChainUnaryInterceptor(
			ji.JwtClientInterceptor(),
			li.LogUnaryServerInterceptor(),
		),
		//grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin": {}}]}`),
		//grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"pick_first": {}}]}`),
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"my_random": {}}]}`),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	//conn, err := grpc.Dial(
	//	"localhost:50051",
	//	grpc.WithTransportCredentials(insecure.NewCredentials()),
	//	grpc.WithChainUnaryInterceptor(
	//		ji.JwtClientInterceptor(),
	//		li.LogUnaryServerInterceptor()))
	//if err != nil {
	//	log.Fatalf("did not connect: %v", err)
	//}
	//
	//defer conn.Close()

	client := pb.NewUserServiceClient(conn)
	resp, err := client.Hi(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Println(resp.Msg)
}
