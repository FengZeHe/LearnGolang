package repository

import (
	"context"
	"github.com/basicprojectv2/internal/domain"
	"github.com/basicprojectv2/internal/repository/cache"
	"github.com/basicprojectv2/internal/repository/dao"
)

type UserRepository interface {
	Create(ctx context.Context, u domain.User) error
	FindByEmail(ctx context.Context, email string) (domain.User, error)
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

func (repo *CacheUserRepository) Create(ctx context.Context, u domain.User) (err error) {
	return repo.dao.Insert(ctx, u)
}

func (repo *CacheUserRepository) FindByEmail(ctx context.Context, email string) (u domain.User, err error) {
	return repo.dao.FindByEmail(ctx, email)
}
