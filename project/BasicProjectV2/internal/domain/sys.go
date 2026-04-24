package domain

type SimplifyMenu struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"size:255;not null" json:"name"`
}

type GetRoleMenuListReq struct {
	RoleName string `json:"role_name"`
}

type GetRoleApiListReq struct {
	RoleName string `json:"role_name"`
}

type API struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Name    string `gorm:"size:255;" json:"name"`
	Url     string `gorm:"size:255;" json:"url"`
	Methods string `gorm:"size:255;" json:"methods"`
}

type UserProfile struct {
	ID       uint   `gorm:"primaryKey" json:"userID"`
	Email    string `gorm:"column:email" json:"email"`
	Role     string `gorm:"column:role" json:"role"`
	Phone    string `gorm:"column:phone" json:"phone"`
	Birthday string `gorm:"size:255;" json:"-"`
	NickName string `gorm:"column:nickname" json:"nickName"`
	AboutMe  string `gorm:"column:aboutme" json:"aboutMe"`
}

func (u UserProfile) TableName() string {
	return "users"
}

type UserAvatar struct {
	UserID     string ` json:"userID"`
	AvatarFile []byte ` json:"avatarFile"`
}

type UploadFile struct {
	UserID   string `json:"userID"`
	FileURL  string `json:"fileURL"`
	FileName string `json:"fileName"`
	FileType string `json:"fileType"`
	Ctime    string `json:"CTime"`
	File     []byte `json:"file"`
}

type UploadFileReq struct {
	UserID   string `json:"userID"`
	FileName string `json:"fileName"`
	FileType string `json:"fileType"`
	File     []byte `json:"file"`
}

type DownloadFileReq struct {
	UserID   string `json:"userID"`
	FileName string `json:"fileName"`
}

type UpdateCasbinPolicyReq struct {
	OldPolicy []string `gorm:"size:255;" json:"old_policy"`
	NewPolicy []string `gorm:"size:255;" json:"new_policy"`
}

type AddCasbinRulePolicyReq struct {
	NewPolicy []string `gorm:"size:255;" json:"new_policy"`
}

type RemoveCasbinPolicyReq struct {
	RemovePolicy []string `gorm:"size:255;" json:"remove_policy"`
}

type TransactionPolicyReq struct {
	OldPolicies [][]string `gorm:"size:255;" json:"old_policies"`
	NewPolicies [][]string `gorm:"size:255;" json:"new_policies"`
}
