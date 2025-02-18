package repository

import (
	"github.com/basicprojectv2/internal/domain"
	"github.com/basicprojectv2/internal/repository/dao"
)

type roleRepository struct {
	dao dao.GORMRoleDAO
}

type RoleRepository interface {
	GetRoles() ([]domain.Role, error)
}

func NewRoleRepository(dao dao.GORMRoleDAO) RoleRepository {
	return &roleRepository{dao: dao}
}

func (r *roleRepository) GetRoles() (roles []domain.Role, err error) {
	roles, err = r.dao.GetRoles()
	if err != nil {
		return nil, err
	}
	return roles, err
}
