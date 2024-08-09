package dao

import (
	"context"
	"github.com/basicprojectv2/internal/domain"
	"gorm.io/gorm"
	"log"
)

type GORMSysDAO struct {
	db *gorm.DB
}

type SysDAO interface {
	FindByEmail(ctx context.Context) error
	FindMenusByRole(ctx context.Context, role string) (menuItems []domain.Menu, err error)
	FindUserByID(ctx context.Context, id string) (user domain.User, err error)
}

func NewSysDAO(db *gorm.DB) SysDAO {
	return &GORMSysDAO{
		db: db,
	}
}

func (dao *GORMSysDAO) FindUserByID(ctx context.Context, id string) (user domain.User, err error) {
	err = dao.db.Table("users").Where("id = ?", id).Find(&user).Error
	if err != nil {
		log.Println("DAO find user by ID error", err)
		return user, err
	}
	return user, nil
}

func (dao *GORMSysDAO) FindMenusByRole(ctx context.Context, role string) (menuItems []domain.Menu, err error) {
	err = dao.db.Table("menu").
		Select("menu.id, menu.name, menu.path, menu.parentid,menu.orderno").
		Joins("JOIN casbin_rule ON casbin_rule.v1 = menu.path").
		Where("casbin_rule.v0 = ?", role).
		Order("menu.orderno").Scan(&menuItems).Error
	if err != nil {
		log.Println("dao get menu error", err)
	}

	return menuItems, nil
}

func (dao *GORMSysDAO) FindByEmail(ctx context.Context) error {
	return nil
}
