package repository

import (
	"context"
	"github.com/basicprojectv2/internal/domain"
	"github.com/basicprojectv2/internal/repository/dao"
)

type SysRepository interface {
	GetMenu(ctx context.Context, u domain.User) error
}

type sysRepository struct {
	dao dao.SysDAO
}

func (s sysRepository) GetMenu(ctx context.Context, u domain.User) error {
	//TODO implement me
	panic("implement me")
}

func NewSysRepository(dao dao.SysDAO) SysRepository {
	return &sysRepository{
		dao: dao,
	}
}
