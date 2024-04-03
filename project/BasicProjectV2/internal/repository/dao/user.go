package dao

import (
	"context"
	"github.com/basicprojectv2/internal/domain"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"log"
)

var (
	ErrDuplicateEmail = errors.New("邮箱冲突")
	ErrRecordNotFound = gorm.ErrRecordNotFound
)

type UserDAO interface {
	Insert(ctx context.Context, u domain.User) error
	FindByEmail(ctx context.Context, email string) (domain.User, error)
	FindByPhone(ctx context.Context, phone string) (domain.User, error)
}

type GORMUserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) UserDAO {
	return &GORMUserDAO{
		db: db,
	}
}

func (dao *GORMUserDAO) Insert(ctx context.Context, u domain.User) (err error) {
	if err = dao.db.WithContext(ctx).Create(&u).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (dao *GORMUserDAO) FindByEmail(ctx context.Context, email string) (u domain.User, err error) {
	err = dao.db.WithContext(ctx).Table("users").Where("email = ?", email).First(&u).Error
	return u, err
}

func (dao *GORMUserDAO) FindByPhone(ctx context.Context, phone string) (u domain.User, err error) {
	err = dao.db.WithContext(ctx).Table("users").Where("phone = ?", phone).First(&u).Error
	return u, err
}
