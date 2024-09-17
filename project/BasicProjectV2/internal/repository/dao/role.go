package dao

import (
	"github.com/basicprojectv2/internal/domain"
	"gorm.io/gorm"
)

type GORMRoleDAO struct {
	db *gorm.DB
}

type RoleDAO interface {
	GetRoles() ([]domain.Role, error)
}

func NewRoleDAO(db *gorm.DB) GORMRoleDAO {
	return GORMRoleDAO{
		db: db,
	}
}

func (r *GORMRoleDAO) GetRoles() (roles []domain.Role, err error) {
	if err = r.db.Table("roles").Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, err
}
