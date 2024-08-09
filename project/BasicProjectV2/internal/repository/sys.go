package repository

import (
	"context"
	"github.com/basicprojectv2/internal/domain"
	"github.com/basicprojectv2/internal/repository/dao"
	"log"
)

type SysRepository interface {
	GetMenu(ctx context.Context, id string) ([]domain.Menu, error)
}

type sysRepository struct {
	dao dao.SysDAO
}

func (s sysRepository) GetMenu(ctx context.Context, id string) ([]domain.Menu, error) {
	user, err := s.dao.FindUserByID(ctx, id)
	// todo 通过userID查询到Role
	menus, err := s.dao.FindMenusByRole(ctx, user.Role)
	if err != nil {
		log.Println("repo Get Menus By Role Error", err)
	}
	return menus, err
}

func NewSysRepository(dao dao.SysDAO) SysRepository {
	return &sysRepository{
		dao: dao,
	}
}
