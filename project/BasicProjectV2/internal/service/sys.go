package service

import (
	"context"
	"github.com/basicprojectv2/internal/domain"
	"github.com/basicprojectv2/internal/repository"
)

type SysService interface {
	GetMenuByUserID(ctx context.Context, userid string) (menus []domain.Menu, err error)
	GetMenu(ctx context.Context) ([]domain.SimplifyMenu, error)
	GetRole(ctx context.Context) ([]domain.Role, error)
	GetAPI(ctx context.Context) ([]domain.API, error)
}

type sysService struct {
	repo repository.SysRepository
}

func NewSysService(repo repository.SysRepository) SysService {
	return &sysService{repo: repo}
}

func (s *sysService) GetRole(ctx context.Context) ([]domain.Role, error) {
	rl, err := s.repo.GetRole(ctx)
	return rl, err
}

func (s *sysService) GetAPI(ctx context.Context) ([]domain.API, error) {
	rl, err := s.repo.GetAPI(ctx)
	return rl, err
}

func (s *sysService) GetMenu(ctx context.Context) (menus []domain.SimplifyMenu, err error) {
	menus, err = s.repo.GetMenu(ctx)
	return menus, err
}

func (s *sysService) GetMenuByUserID(ctx context.Context, userid string) (menus []domain.Menu, err error) {
	menus, err = s.repo.GetMenuByUserID(ctx, userid)
	if err != nil {
		return nil, err
	}
	//处理用户返回菜单
	return menus, err
}
