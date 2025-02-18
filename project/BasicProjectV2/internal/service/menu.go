package service

import (
	"context"
	"github.com/basicprojectv2/internal/domain"
	"github.com/basicprojectv2/internal/repository"
)

type MenuService interface {
	GetMenuByUserID(ctx context.Context, userid string) ([]domain.Menu, error)
}

type menuService struct {
	repo repository.MenuRepository
}

func NewMenuService(repo repository.MenuRepository) MenuService {
	return &menuService{
		repo: repo,
	}
}

func (s *menuService) GetMenuByUserID(ctx context.Context, userid string) (menus []domain.Menu, err error) {
	menus, err = s.repo.GetMenuByUserID(ctx, userid)
	if err != nil {
		return nil, err
	}
	//处理用户返回菜单
	return menus, err
}
