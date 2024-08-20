package dao

import (
	"gorm.io/gorm"
	"log"
)

type GORMMenuDAO struct {
	db *gorm.DB
}

type MenuDAO interface {
	GetMenuList() ([]Menu, error)
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
	log.Println("dao", menus)
	return menus, nil
}

type Menu struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Path     string `json:"path"`
	ParentId string `json:"parent_id" gorm:"column:parentid"`
	OrderNo  string `json:"order_no" gorm:"column:orderno"`
}
