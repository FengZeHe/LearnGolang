package repository

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"github.com/basicprojectv2/internal/domain"
	"github.com/basicprojectv2/internal/repository/cache"
	"github.com/basicprojectv2/internal/repository/dao"
	"io"
	"log"
	"os"
	"path/filepath"
)

var (
	ErrDuplicateUser = dao.ErrDuplicateEmail
	ErrUserNotFound  = dao.ErrRecordNotFound
)

type UserRepository interface {
	Create(ctx context.Context, u domain.User) error
	FindByEmail(ctx context.Context, email string) (domain.User, error)
	FindByPhone(ctx context.Context, phone string) (domain.User, error)
	FindById(ctx context.Context, id string) (domain.User, error)
	GetUserList(ctx context.Context, req domain.UserListRequest) ([]domain.User, int, error)
	UpdateUser(ctx context.Context, req domain.User) error
	UpsertUserAvatar(ctx context.Context, req domain.UserAvatar) error
	UploadUserFile(ctx context.Context, req domain.UploadFile) error
}

type CacheUserRepository struct {
	dao   dao.UserDAO
	cache cache.UserCache
}

func NewCacheUserRepository(dao dao.UserDAO, c cache.UserCache) UserRepository {
	return &CacheUserRepository{
		dao:   dao,
		cache: c,
	}
}

func (repo *CacheUserRepository) UpdateUser(ctx context.Context, u domain.User) (err error) {
	if err = repo.dao.UpdateUserByID(ctx, repo.toEntity(u)); err != nil {
		return err
	}
	return nil
}

func (repo *CacheUserRepository) UploadUserFile(ctx context.Context, req domain.UploadFile) (err error) {
	// 将文件存储到指定路径
	uploadPath := fmt.Sprintf("/Users/feng/Desktop/%s", req.UserID)

	// 检查文件夹是否存在，如果不存在则创建
	if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
		err := os.MkdirAll(uploadPath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	// MySQL查询不重复的fileName
	UniqueFileName, err := repo.dao.CheckUniqueFileName(ctx, req)
	if err != nil {
		return err
	}
	UniqueFileName = fmt.Sprintf("%s.%s", UniqueFileName, "jpg")
	filePath := filepath.Join(uploadPath, UniqueFileName)

	// 创建保存文件
	outFile, err := os.Create(filePath)
	if err != nil {
		log.Println("output file ERROR", err)
		return err
	}

	defer outFile.Close()

	// 将文件写入到本地文件夹
	reader := bytes.NewReader(req.File)
	_, err = io.Copy(outFile, reader)
	if err != nil {
		log.Println("Failed to save file")
		return err
	}

	if err = repo.dao.InsertUserFile(ctx, req); err != nil {
		return err
	}
	return nil
}

func (repo *CacheUserRepository) GetUserList(ctx context.Context, req domain.UserListRequest) (ul []domain.User, count int, err error) {
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

func (repo *CacheUserRepository) Create(ctx context.Context, u domain.User) (err error) {
	eu := repo.toEntity(u)
	return repo.dao.Insert(ctx, eu)
}

func (repo *CacheUserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	u, err := repo.dao.FindByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return repo.toDomain(u), nil
}

func (repo *CacheUserRepository) FindByPhone(ctx context.Context, phone string) (domain.User, error) {
	u, err := repo.dao.FindByPhone(ctx, phone)
	if err != nil {
		return domain.User{}, err
	}
	return repo.toDomain(u), nil
}

func (repo *CacheUserRepository) FindById(ctx context.Context, id string) (domain.User, error) {
	// 首先找一下redis有没有
	du, err := repo.cache.Get(ctx, id)
	if err == nil {
		return du, nil
	}
	u, err := repo.dao.FindById(ctx, id)
	if err != nil {
		return domain.User{}, err
	}
	du = repo.toDomain(u)

	if err = repo.cache.Set(ctx, du); err != nil {
		log.Println("set cache error", err)
	}
	return du, nil
}

func (repo *CacheUserRepository) UpsertUserAvatar(ctx context.Context, req domain.UserAvatar) (err error) {
	if err = repo.dao.UpsertUserAvatar(ctx, req); err != nil {
		return err
	}
	return nil
}

// addReqToDomain 将dao.User转为 domain.User
func (repo *CacheUserRepository) toDomain(u dao.User) domain.User {
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
func (repo *CacheUserRepository) toEntity(u domain.User) dao.User {
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
