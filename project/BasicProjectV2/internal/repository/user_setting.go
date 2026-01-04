package repository

import (
	"context"

	"github.com/basicprojectv2/internal/domain"
	"github.com/basicprojectv2/internal/repository/dao"
)

type userSettingRepository struct {
	userSettingDAO dao.UserSettingDAO
}

type UserSettingRepository interface {
	HandleUserSetting(uid string, us domain.UserSettingReq, ctx context.Context) (err error)
	GetUserSetting(uid string, ctx context.Context) (domain.UserSetting, error)
}

func NewUserSettingRepository(dao dao.UserSettingDAO) UserSettingRepository {
	return &userSettingRepository{userSettingDAO: dao}
}

func (u *userSettingRepository) HandleUserSetting(uid string, us domain.UserSettingReq, ctx context.Context) (err error) {
	return u.userSettingDAO.HandleUserSetting(uid, us, ctx)
}

func (u *userSettingRepository) GetUserSetting(uid string, ctx context.Context) (domain.UserSetting, error) {
	return u.userSettingDAO.GetUserSetting(uid, ctx)
}
