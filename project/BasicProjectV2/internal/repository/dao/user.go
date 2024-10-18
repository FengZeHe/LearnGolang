package dao

import (
	"context"
	"database/sql"
	"github.com/basicprojectv2/internal/domain"
	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
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
	GetUserList(ctx context.Context, req domain.UserListRequest) ([]User, int, error)
	UpdateUserByID(ctx context.Context, u User) error
	UpsertUserAvatar(ctx context.Context, u domain.UserAvatar) error
}

func NewUserDAO(db *gorm.DB) UserDAO {
	return &GORMUserDAO{
		db: db,
	}
}

func (dao *GORMUserDAO) UpdateUserByID(ctx context.Context, u User) (err error) {
	var user User
	if err = dao.db.First(&user, u.ID).Error; err != nil {
		log.Println("User not found", err)
		return err
	}

	user.Email = u.Email
	user.Phone = u.Phone
	user.Role = u.Role
	user.Aboutme = u.Aboutme
	user.Birthday = u.Birthday
	user.Nickname = u.Nickname

	if err = dao.db.WithContext(ctx).Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func (dao *GORMUserDAO) GetUserList(ctx context.Context, req domain.UserListRequest) (ul []User, count int, err error) {
	// 在gorm中实现分页， Limit用户设置每页的记录数，offset用于跳过指定数量的记录
	// 计算offset
	offset := (req.PageIndex - 1) * req.PageSize
	if err = dao.db.WithContext(ctx).Limit(req.PageSize).Offset(offset).Find(&ul).Error; err != nil {
		log.Println("dao Get User List ERROR", err)
		return ul, count, err
	}
	return ul, len(ul), nil
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

func (dao *GORMUserDAO) UpsertUserAvatar(ctx context.Context, u domain.UserAvatar) (err error) {
	if err = dao.db.Table("user_avatar").Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}},                // 使用user_id作为唯一键判断冲突
		DoUpdates: clause.AssignmentColumns([]string{"avatar_file"}), // 发生冲突时更新avatar_file
	}).Create(&u).Error; err != nil {
		return err
	}
	return nil
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
