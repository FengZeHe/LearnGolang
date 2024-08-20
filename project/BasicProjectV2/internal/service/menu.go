package service

import (
	"context"
	"github.com/basicprojectv2/internal/domain"
	"github.com/basicprojectv2/internal/repository"
)

type MenuService interface {
	GetMenus(ctx context.Context) ([]domain.Menu, error)
}

type menuService struct {
	repo repository.MenuRepository
}

func NewMenuService(repo repository.MenuRepository) MenuService {
	return &menuService{
		repo: repo,
	}
}

func (m *menuService) GetMenus(ctx context.Context) (menu []domain.Menu, err error) {
	menu, err = m.repo.GetList()
	if err != nil {
		return menu, err
	}
	return menu, nil
}
