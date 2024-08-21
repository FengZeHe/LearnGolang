package service

import (
	"github.com/basicprojectv2/internal/domain"
	"github.com/basicprojectv2/internal/repository"
)

type roleService struct {
	repo repository.RoleRepository
}

type RoleService interface {
	GetRoles() ([]domain.Role, error)
}

func NewRoleService(repo repository.RoleRepository) RoleService {
	return &roleService{
		repo: repo,
	}
}

func (r *roleService) GetRoles() (roles []domain.Role, err error) {
	roles, err = r.repo.GetRoles()
	if err != nil {
		return nil, err
	}
	return roles, err
}
