package service

import (
	"context"
	"github.com/basicprojectv2/internal/domain"
	"github.com/basicprojectv2/internal/repository"
)

type SysService interface {
	GetMenu(ctx context.Context, userid string) (menus []domain.Menu, err error)
}

type sysService struct {
	repo repository.SysRepository
}

func NewSysService(repo repository.SysRepository) SysService {
	return &sysService{repo: repo}
}

func (s *sysService) GetMenu(ctx context.Context, userid string) (menus []domain.Menu, err error) {

	menus, err = s.repo.GetMenu(ctx, userid)
	if err != nil {
		return nil, err
	}
	//处理用户返回菜单
	return menus, err
}
