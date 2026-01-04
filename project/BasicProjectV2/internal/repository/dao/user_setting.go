package dao

import (
	"context"
	"log"
	"time"

	"github.com/basicprojectv2/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GORMUserSetting struct {
	db *gorm.DB
}

type UserSettingDAO interface {
	HandleUserSetting(uid string, us domain.UserSettingReq, ctx context.Context) (err error)
	GetUserSetting(uid string, ctx context.Context) (*domain.UserSettingReq, error)
	ResetUserSetting(uid string, ctx context.Context) error
}

func NewUserSettingDAO(db *gorm.DB) UserSettingDAO {
	return &GORMUserSetting{db: db}
}

func (u *GORMUserSetting) HandleUserSetting(uid string, us domain.UserSettingReq, ctx context.Context) (err error) {
	log.Println("Handle User Setting", us, uid)
	/*
		todo upsert操作
	*/
	now := time.Now().Format("2006-01-02 15:04:05")
	return u.db.Transaction(func(tx *gorm.DB) error {
		var usRec domain.UserSettingReq
		err = tx.Model(domain.UserSettingReq{}).Where("id = ?", uid).First(&usRec).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}

		if err = tx.Model(domain.UserSetting{}).Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "user_id"}},
			DoUpdates: clause.Assignments(map[string]any{
				"theme_mode": us.ThemeMode,
				"updated_at": now,
			}),
		}).Create(&domain.UserSetting{
			UserID:    uid,
			ThemeMode: us.ThemeMode,
			CreatedAt: now,
		}).Error; err != nil {
			return err
		}
		return nil
	})

}

func (u *GORMUserSetting) GetUserSetting(uid string, ctx context.Context) (*domain.UserSettingReq, error) {
	return &domain.UserSettingReq{}, nil
}

func (u *GORMUserSetting) ResetUserSetting(uid string, ctx context.Context) error {
	return nil
}
