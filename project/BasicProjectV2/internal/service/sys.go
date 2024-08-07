package service

import (
	"context"
	"github.com/basicprojectv2/internal/domain"
	"github.com/basicprojectv2/internal/repository"
)

type SysService interface {
	GetMenu(ctx context.Context, u domain.User) error
}

type sysService struct {
	repo repository.SysRepository
}

func (s sysService) GetMenu(ctx context.Context, u domain.User) error {
	//TODO implement me
	panic("implement me")
}

func NewSysService(repo repository.SysRepository) SysService {
	return &sysService{repo: repo}
}
