package dao

import (
	"context"
	"github.com/basicprojectv2/internal/domain"
	"gorm.io/gorm"
	"log"
)

type GORMMenuDAO struct {
	db *gorm.DB
}

type MenuDAO interface {
	GetMenuList() ([]Menu, error)
	GetMenuListByRole(role string) ([]Menu, error)
	FindUserByID(ctx context.Context, id string) (domain.User, error)
	FindMenusByRole(ctx context.Context, role string) ([]domain.Menu, error)
}

func NewMenuDAO(db *gorm.DB) GORMMenuDAO {
	return GORMMenuDAO{
		db: db,
	}
}

func (dao *GORMMenuDAO) GetMenuList() (menus []Menu, err error) {
	if err = dao.db.Table("menu").Find(&menus).Error; err != nil {
		return menus, err
	}
	return menus, nil
}

func (dao *GORMMenuDAO) GetMenuListByRole(role string) (menus []Menu, err error) {
	err = dao.db.Table("menu").
		Select("menu.id, menu.name, menu.path, menu.parentid,menu.orderno,menu.methods").
		Joins("JOIN casbin_rule ON casbin_rule.v1 = menu.path AND casbin_rule.v2 = menu.methods").
		Where("casbin_rule.v0 = ?", role).
		Order("menu.orderno").Scan(&menus).Error
	if err != nil {
		log.Println("dao get menu error", err)
	}
	log.Println("--->", menus)
	return menus, err
}

func (dao *GORMMenuDAO) FindUserByID(ctx context.Context, id string) (user domain.User, err error) {
	err = dao.db.Table("users").Where("id = ?", id).Find(&user).Error
	if err != nil {
		log.Println("DAO find user by ID error", err)
		return user, err
	}
	return user, nil
}

func (dao *GORMMenuDAO) FindMenusByRole(ctx context.Context, role string) (menus []domain.Menu, err error) {
	err = dao.db.Table("menu").
		Select("menu.id, menu.name, menu.path, menu.parentid,menu.orderno,menu.methods").
		Joins("JOIN casbin_rule ON casbin_rule.v1 = menu.path AND casbin_rule.v2 = menu.methods").
		Where("casbin_rule.v0 = ?", role).
		Order("menu.orderno").Scan(&menus).Error
	if err != nil {
		log.Println("dao get menu error", err)
	}
	return menus, nil
}

type Menu struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Path     string `json:"path"`
	ParentId string `json:"parent_id" gorm:"column:parentid"`
	OrderNo  string `json:"order_no" gorm:"column:orderno"`
}
