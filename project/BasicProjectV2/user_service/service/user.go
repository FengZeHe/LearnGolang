package service

import (
	"context"
	"errors"
	"github.com/basicprojectv2/user_service/domain"
	"github.com/basicprojectv2/user_service/pkg/bcrypt"
	"github.com/basicprojectv2/user_service/pkg/snowflake"
	"github.com/basicprojectv2/user_service/repository"

	//"github.com/pkg/errors"
	"strconv"
)

var (
	ErrDuplicateEmail        = repository.ErrDuplicateUser
	ErrInvalidUserOrPassword = errors.New("用户不存在或者密码不对")
	ErrGenPasswd             = bcrypt.ErrGenPasswd
)

type UserService interface {
	Signup(ctx context.Context, u domain.User) error
	Login(ctx context.Context, email string, password string) (domain.User, error)
	FindOrCreate(ctx context.Context, phone string, id string) (domain.User, error)
	FindById(ctx context.Context, id string) (domain.User, error)
	GetUserList(ctx context.Context, req domain.UserListRequest) (domain.UserListResponse, error)
	UpdateUser(ctx context.Context, req domain.User) error
}
type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) UpdateUser(ctx context.Context, req domain.User) error {

	err := s.repo.UpdateUser(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

// 返回用户列表
func (svc *userService) GetUserList(ctx context.Context, req domain.UserListRequest) (ul domain.UserListResponse, err error) {
	l, c, err := svc.repo.GetUserList(ctx, req)
	ul.Users = l
	ul.Count = c
	if err != nil {
		return ul, err
	}
	return ul, nil
}

func (svc *userService) Signup(ctx context.Context, u domain.User) (err error) {
	id := snowflake.GenId()
	hashStr, err := bcrypt.GetPwd(u.Password)
	if err == bcrypt.ErrGenPasswd {
		return ErrGenPasswd
	}
	u.Password = hashStr
	u.ID = strconv.Itoa(id)
	return svc.repo.Create(ctx, u)
}

func (svc *userService) Login(ctx context.Context, email string, password string) (u domain.User, err error) {
	u, err = svc.repo.FindByEmail(ctx, email)
	// 用户不存在
	if err == repository.ErrUserNotFound {
		return domain.User{}, ErrInvalidUserOrPassword
	}
	// 对比密码
	if result := bcrypt.ComparePwd(u.Password, password); result != true {
		// 密码错误
		//err = errors.New("passwordError")
		err = ErrInvalidUserOrPassword
		return domain.User{}, err
	}
	return u, nil
}

func (svc *userService) FindOrCreate(ctx context.Context, phone string, id string) (u domain.User, err error) {
	// 先找一个是否有该用户，如果没有的话创建一个
	u, err = svc.repo.FindByPhone(ctx, phone)
	if err != repository.ErrUserNotFound {
		return u, err
	}

	err = svc.repo.Create(ctx, domain.User{Phone: phone, ID: id})
	if err != nil {
		// 创建失败
		return domain.User{}, err
	}
	return svc.repo.FindByPhone(ctx, phone)
}

func (svc *userService) FindById(ctx context.Context,
	id string) (domain.User, error) {
	return svc.repo.FindById(ctx, id)
}
