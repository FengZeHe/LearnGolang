package dao

import (
	"context"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var (
	ErrDuplicateEmail = errors.New("邮箱冲突")
	ErrRecordNotFound = gorm.ErrRecordNotFound
)

type GORMUserDAO struct {
	db *gorm.DB
}

type UserDAO interface {
	Insert(ctx context.Context, u User) error
	FindByEmail(ctx context.Context, email string) (User, error)
	FindByPhone(ctx context.Context, phone string) (User, error)
	FindById(ctx context.Context, id string) (User, error)
}

func NewUserDAO(db *gorm.DB) UserDAO {
	return &GORMUserDAO{
		db: db,
	}
}

func (dao *GORMUserDAO) Insert(ctx context.Context, u User) (err error) {
	err = dao.db.WithContext(ctx).Create(&u).Error
	if me, ok := err.(*mysql.MySQLError); ok {
		const duplicateErr uint16 = 1062
		if me.Number == duplicateErr {
			// 用户冲突，邮箱冲突
			return ErrDuplicateEmail
		}
	}
	return err
}

func (dao *GORMUserDAO) FindByEmail(ctx context.Context, email string) (u User, err error) {
	err = dao.db.WithContext(ctx).Table("users").Where("email = ?", email).First(&u).Error
	return u, err
}

func (dao *GORMUserDAO) FindByPhone(ctx context.Context, phone string) (u User, err error) {
	err = dao.db.WithContext(ctx).Table("users").Where("phone = ?", phone).First(&u).Error
	return u, err
}

func (dao *GORMUserDAO) FindById(ctx context.Context, id string) (u User, err error) {
	err = dao.db.WithContext(ctx).Table("users").Where("id = ?", id).First(&u).Error
	return u, err
}

type User struct {
	ID       string         `json:"id"`
	Email    sql.NullString `json:"email"`
	Password string         `json:"password"`
	Phone    sql.NullString `json:"phone"`
	Birthday int            `json:"birthday"`
	Nickname string         `json:"nickname"`
	Aboutme  string         `json:"aboutme"`
	Role     string         `json:"role"`
}
