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
}

func NewUserSettingRepository(dao dao.UserSettingDAO) UserSettingRepository {
	return &userSettingRepository{userSettingDAO: dao}
}

func (u *userSettingRepository) HandleUserSetting(uid string, us domain.UserSettingReq, ctx context.Context) (err error) {
	return u.userSettingDAO.HandleUserSetting(uid, us, ctx)
}
