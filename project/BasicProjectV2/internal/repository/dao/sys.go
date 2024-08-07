package dao

import (
	"context"
	"gorm.io/gorm"
)

type GORMSysDAO struct {
	db *gorm.DB
}

type SysDAO interface {
	FindByEmail(ctx context.Context) error
}

func NewSysDAO(db *gorm.DB) SysDAO {
	return &GORMSysDAO{
		db: db,
	}
}

func (dao *GORMSysDAO) FindByEmail(ctx context.Context) error {
	return nil
}
