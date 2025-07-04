package main

import (
	"context"
	pb "github.com/basicprojectv2/proto/user_service"
	"github.com/basicprojectv2/user_service/interceptors/jwt"
	"github.com/basicprojectv2/user_service/serviceReg"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
	"net/http"
	"time"
)

type App struct {
	grpcServer *grpc.Server
	gwServer   *http.Server
	userSvc    *UserService
	healthSvc  grpc_health_v1.HealthServer
	etcdClient *serviceReg.EtcdClient
}

func NewApp(userSvc *UserService, etcdClient *serviceReg.EtcdConfig) *App {
	jwtInterceptor := jwt.NewJWTInterceptor([]string{
		// 免检路径
		/*
			    在 gRPC 中，每个服务方法都有一个唯一的全限定路径（Full Method Name），
				格式为：/包名.服务名/方法名。这个路径用于客户端与服务器之间的通信，也是拦截器中配置免检路径的依据。
						package声明/ proto中对应的service关键字  / Userlogin
		*/
		"/grpc.health.v1.Health/Check",
		"/grpc.health.v1.Health/Watch",
		"/user_service.UserService/UserLogin",
	})

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(jwtInterceptor.UnaryInterceptor()))

	// 注册健康检查服务
	healthServer := health.NewServer()

	// 包装器-health check 服务器
	wrappedHealthServer := &healthServerWrapper{healthServer: healthServer}
	healthServer.SetServingStatus("user_service.UserService", grpc_health_v1.HealthCheckResponse_SERVING)

	ec, _ := serviceReg.NewEtcdClient(etcdClient)

	return &App{
		grpcServer: grpcServer,
		userSvc:    userSvc,
		healthSvc:  wrappedHealthServer,
		etcdClient: ec,
	}
}

func (a *App) Start() error {
	pb.RegisterUserServiceServer(a.grpcServer, a.userSvc)

	grpc_health_v1.RegisterHealthServer(a.grpcServer, a.healthSvc.(*healthServerWrapper))

	//healthServer := a.healthSvc.(*health.Server)
	//healthServer.SetServingStatus("user_service.UserService", grpc_health_v1.HealthCheckResponse_SERVING)
	//grpc_health_v1.RegisterHealthServer(a.grpcServer, healthServer)

	log.Println("已注册服务", a.grpcServer.GetServiceInfo())

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}
	log.Println("start gRPC Server on 50051")

	// 注册到etcd
	serviceName := "user_service"
	serviceAddr := "localhost:50051"

	if err = a.etcdClient.RegisterService(serviceName, serviceAddr); err != nil {
		log.Printf("register service %s to ETCD error: %v", serviceName, err)
	}

	// 10s后服务下线
	t1 := time.NewTimer(10 * time.Second)
	go func() {
		<-t1.C
		log.Println("时间到喽")
		if err := a.etcdClient.UnregisterService(); err != nil {
			log.Println("unregister service error:", err)
		} else {
			log.Println("unregister service success")
		}
		// 退出程序
		//os.Exit(0)
	}()

	go func() {
		if err := a.grpcServer.Serve(lis); err != nil {
			log.Println(err)
		}
	}()

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Println("HTTP server listening at :9091")
		log.Fatal(http.ListenAndServe(":9091", nil))
	}()

	conn, err := grpc.NewClient(
		"0.0.0.0:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Failed to dial server:", err)
	}
	defer conn.Close()

	gwmux := runtime.NewServeMux()
	err = pb.RegisterUserServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatal("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}
	// 8090端口提供gRPC-Gateway服务
	log.Println("Serving gRPC-Gateway on in 8090...")
	log.Fatalln(gwServer.ListenAndServe())
	return a.gwServer.ListenAndServe()
}

var (
	grpcHealthCheckStatus = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "grpc_health_check_status",
			Help: "gRPC健康检查 1正常 0不健康",
		},
		[]string{"user_service"})
)

func init() {
	prometheus.MustRegister(grpcHealthCheckStatus)
}

type healthServerWrapper struct {
	healthServer *health.Server
}

func (h *healthServerWrapper) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	// 执行健康检查
	response, err := h.healthServer.Check(ctx, req)
	// 根据健康检查结果更新 Prometheus 指标
	if err == nil && response.Status == grpc_health_v1.HealthCheckResponse_SERVING {
		grpcHealthCheckStatus.WithLabelValues(req.Service).Set(1)
	} else {
		grpcHealthCheckStatus.WithLabelValues(req.Service).Set(0)
	}
	return response, err
}

func (h *healthServerWrapper) Watch(req *grpc_health_v1.HealthCheckRequest, stream grpc_health_v1.Health_WatchServer) error {
	// 执行健康检查状态观察
	err := h.healthServer.Watch(req, stream)
	// 根据健康检查状态更新 Prometheus 指标
	// 这里可以根据实际需求进一步处理
	return err
}
