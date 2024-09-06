package repository

import (
	"context"
	"github.com/basicprojectv2/internal/domain"
	"github.com/basicprojectv2/internal/repository/dao"
	"log"
)

type SysRepository interface {
	GetApiByUserID(ctx context.Context, id string) ([]domain.API, error)
	GetMenuByUserID(ctx context.Context, id string) ([]domain.Menu, error)
	GetMenuByRole(ctx context.Context, role string) ([]domain.Menu, error)
	GetAPIByRole(ctx context.Context, role string) ([]domain.API, error)
	GetMenu(ctx context.Context) ([]domain.Menu, error)
	GetRole(ctx context.Context) ([]domain.Role, error)
	GetAPI(ctx context.Context) ([]domain.API, error)
}

type sysRepository struct {
	dao dao.SysDAO
}

func (s sysRepository) GetApiByUserID(ctx context.Context, id string) (apis []domain.API, err error) {
	user, err := s.dao.FindUserByID(ctx, id)
	// todo 通过userID查询到Role
	apis, err = s.dao.FindApisByRole(ctx, user.Role)
	if err != nil {
		log.Println("repo Get Menus By Role Error", err)
		return nil, err
	}
	return apis, nil
}

func (s sysRepository) GetAPI(ctx context.Context) (al []domain.API, err error) {
	al, err = s.dao.GetAPI(ctx)
	return al, nil
}

func (s sysRepository) GetRole(ctx context.Context) (rl []domain.Role, err error) {
	rl, err = s.dao.GetRole(ctx)
	if err != nil {
		return rl, err
	}
	return rl, nil
}

func (s sysRepository) GetMenu(ctx context.Context) (sm []domain.Menu, err error) {
	sm, err = s.dao.GetMenu(ctx)
	if err != nil {
		return sm, err
	}
	return sm, nil
}

func (s sysRepository) GetMenuByUserID(ctx context.Context, id string) ([]domain.Menu, error) {
	user, err := s.dao.FindUserByID(ctx, id)
	// todo 通过userID查询到Role
	menus, err := s.dao.FindMenusByRole(ctx, user.Role)
	if err != nil {
		log.Println("repo Get Menus By Role Error", err)
	}
	return menus, err
}

func (s sysRepository) GetMenuByRole(ctx context.Context, role string) ([]domain.Menu, error) {
	menus, err := s.dao.FindMenusByRole(ctx, role)
	if err != nil {
		log.Println("repo Get Menus By Role Error", err)
	}
	return menus, err
}

func (s sysRepository) GetAPIByRole(ctx context.Context, role string) ([]domain.API, error) {
	apis, err := s.dao.FindApisByRole(ctx, role)
	if err != nil {
		log.Println("repo Get Apis By Role Error", err)
	}
	return apis, err
}

func NewSysRepository(dao dao.SysDAO) SysRepository {
	return &sysRepository{
		dao: dao,
	}
}
