package ioc

import (
	"github.com/basicprojectv2/internal/web/sms"
	"github.com/basicprojectv2/internal/web/sms/localsms"
)

func InitSMSService() sms.Service {
	return localsms.NewService()
}
