package service

import (
	"context"
	"github.com/basicprojectv2/internal/domain"
	"github.com/basicprojectv2/internal/repository"
	"github.com/basicprojectv2/pkg/bcrypt"
	"github.com/basicprojectv2/pkg/snowflake"
	"github.com/pkg/errors"
	"strconv"
)

type UserService interface {
	Signup(ctx context.Context, u domain.User) error
	Login(ctx context.Context, email string, password string) (domain.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (svc *userService) Signup(ctx context.Context, u domain.User) (err error) {
	id := snowflake.GenId()
	hashStr, err := bcrypt.GetPwd(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashStr
	u.ID = strconv.Itoa(id)
	return svc.repo.Create(ctx, u)
}

func (svc *userService) Login(ctx context.Context, email string, password string) (u domain.User, err error) {
	u, err = svc.repo.FindByEmail(ctx, email)
	// 用户不存在
	if err != nil {
		return domain.User{}, err
	}
	// 对比密码
	if result := bcrypt.ComparePwd(u.Password, password); result != true {
		// 密码错误
		err = errors.New("passwordError")
		return domain.User{}, err
	}
	return u, nil

}
