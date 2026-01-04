package service

import (
	"context"

	"github.com/basicprojectv2/internal/domain"
	"github.com/basicprojectv2/internal/repository"
)

type userSettingService struct {
	repo repository.UserSettingRepository
}

type UserSettingService interface {
	// todo 获取/修改/重置
	HandleUserSetting(uid string, req domain.UserSettingReq, ctx context.Context) error
}

func NewUserSettingService(userSettingRepo repository.UserSettingRepository) UserSettingService {
	return &userSettingService{repo: userSettingRepo}
}

func (us *userSettingService) HandleUserSetting(uid string, req domain.UserSettingReq, ctx context.Context) error {
	return us.repo.HandleUserSetting(uid, req, ctx)
}
