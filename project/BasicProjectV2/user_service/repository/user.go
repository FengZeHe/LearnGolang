package repository

import (
	"context"
	"database/sql"
	"github.com/basicprojectv2/user_service/domain"
	"github.com/basicprojectv2/user_service/repository/dao"
	"github.com/pkg/errors"
	"log"
)

var (
	ErrDuplicateUser = dao.ErrDuplicateEmail
	ErrUserNotFound  = dao.ErrRecordNotFound
	ErrFileNotFound  = dao.ErrFileNotFound
	ErrReadFile      = errors.New("reading file error")
)

type UserRepository interface {
	Create(ctx context.Context, u domain.User) error
	FindByEmail(ctx context.Context, email string) (domain.User, error)
	FindByPhone(ctx context.Context, phone string) (domain.User, error)
	FindById(ctx context.Context, id string) (domain.User, error)
	GetUserList(ctx context.Context, req domain.UserListRequest) ([]domain.User, int, error)
	UpdateUser(ctx context.Context, req domain.User) error
}

type userRepository struct {
	dao dao.UserDAO
}

func (repo *userRepository) FindById(ctx context.Context, id string) (domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserRepository(dao dao.UserDAO) UserRepository {
	return &userRepository{
		dao: dao,
	}
}

func (repo *userRepository) UpdateUser(ctx context.Context, u domain.User) (err error) {
	if err = repo.dao.UpdateUserByID(ctx, repo.toEntity(u)); err != nil {
		return err
	}
	return nil
}

func (repo *userRepository) GetUserList(ctx context.Context, req domain.UserListRequest) (ul []domain.User, count int, err error) {
	list, c, err := repo.dao.GetUserList(ctx, req)
	count = int(c)
	if err != nil {
		log.Println("dao get user list error", err)
	}
	for _, u := range list {
		ul = append(ul, repo.toDomain(u))
	}
	return ul, count, err
}

func (repo *userRepository) Create(ctx context.Context, u domain.User) (err error) {
	eu := repo.toEntity(u)
	return repo.dao.Insert(ctx, eu)
}

func (repo *userRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	u, err := repo.dao.FindByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return repo.toDomain(u), nil
}

func (repo *userRepository) FindByPhone(ctx context.Context, phone string) (domain.User, error) {
	u, err := repo.dao.FindByPhone(ctx, phone)
	if err != nil {
		return domain.User{}, err
	}
	return repo.toDomain(u), nil
}

// addReqToDomain 将dao.User转为 domain.User
func (repo *userRepository) toDomain(u dao.User) domain.User {
	return domain.User{
		ID:       u.ID,
		Email:    u.Email.String,
		Phone:    u.Phone.String,
		Password: u.Password,
		Aboutme:  u.Aboutme,
		Nickname: u.Nickname,
		Birthday: u.Birthday,
		Role:     u.Role,
	}
}

// toEntity 将domain.User 转为 dao.User
func (repo *userRepository) toEntity(u domain.User) dao.User {
	return dao.User{
		ID: u.ID,
		Email: sql.NullString{
			String: u.Email,
			Valid:  u.Email != "",
		},
		Phone: sql.NullString{
			String: u.Phone,
			Valid:  u.Phone != "",
		},
		Password: u.Password,
		Birthday: u.Birthday,
		Aboutme:  u.Aboutme,
		Nickname: u.Nickname,
		Role:     u.Role,
	}
}
