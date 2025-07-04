package ioc

import (
	"github.com/basicprojectv2/user_service/serviceReg"
	"time"
)

func NewEtcdConfig() *serviceReg.EtcdConfig {
	return &serviceReg.EtcdConfig{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
		LeaseTTL:    60 * 60 * 24,
	}
}
