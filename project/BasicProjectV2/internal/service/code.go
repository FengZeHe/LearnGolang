package service

import (
	"context"
	"fmt"
	"github.com/basicprojectv2/internal/repository"
	"github.com/basicprojectv2/internal/web/sms"
	"log"
	"math/rand"
)

type CodeService interface {
	SendCode(ctx context.Context, biz, phone string) error
	VerifyCode(ctx context.Context, biz, phone, inputCode string) (bool, error)
}

type codeService struct {
	repo repository.CodeRepository
	sms  sms.Service
}

func NewCodeService(repo repository.CodeRepository, smsSvc sms.Service) CodeService {
	return &codeService{
		repo: repo,
		sms:  smsSvc,
	}
}

// 发送验证码
func (svc *codeService) SendCode(ctx context.Context, biz, phone string) error {
	code := svc.GenCode()
	if err := svc.repo.SetCode(ctx, biz, phone, code); err != nil {
		log.Println(err)
		return err
	}
	return svc.sms.Send(ctx, "666666", []string{code}, phone)
}

// 验证验证码
func (svc *codeService) VerifyCode(ctx context.Context, biz, phone, inputCode string) (bool, error) {
	ok, err := svc.repo.VerifyCode(ctx, biz, phone, inputCode)
	return ok, err
}

// 生成验证码
func (svc *codeService) GenCode() string {
	code := rand.Intn(10000)
	return fmt.Sprintf("%04d", code)
}
