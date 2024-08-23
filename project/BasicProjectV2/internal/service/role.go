package service

import (
	"github.com/basicprojectv2/internal/domain"
	"github.com/basicprojectv2/internal/repository"
	"github.com/casbin/casbin/v2"
)

type roleService struct {
	repo     repository.RoleRepository
	enforcer *casbin.Enforcer
}

type RoleService interface {
	GetRoles() ([]domain.Role, error)
}

func NewRoleService(repo repository.RoleRepository, enforcer *casbin.Enforcer) RoleService {
	return &roleService{
		repo:     repo,
		enforcer: enforcer,
	}
}

func (r *roleService) GetRoles() (roles []domain.Role, err error) {
	roles, err = r.repo.GetRoles()
	if err != nil {
		return nil, err
	}
	return roles, err
}
