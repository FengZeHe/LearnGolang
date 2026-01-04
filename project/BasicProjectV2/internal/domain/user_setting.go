package domain

type UserSetting struct {
	ID        int    `json:"id" column:"id"`
	UserID    string `json:"user_id"  column:"user_id"`
	ThemeMode string `json:"theme_mode" column:"theme_mode"`
	CreatedAt string `json:"-" column:"created_at"`
	UpdatedAt string `json:"-" column:"updated_at"`
}

type UserSettingReq struct {
	ThemeMode string `json:"themeMode" column:"theme_mode"`
}

func (u UserSettingReq) TableName() string {
	return "user_setting"
}

func (u UserSetting) TableName() string {
	return "user_setting"
}
