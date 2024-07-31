package models

// AccessControlPolicy 表示访问控制策略
type AccessControlPolicy struct {
	ID    uint   `gorm:"primaryKey"`        // 主键
	PType string `gorm:"type:varchar(191)"` //策略类型 p
	V0    string `gorm:"type:varchar(191)"` // 第一个参数 sub
	V1    string `gorm:"type:varchar(191)"` // 第二个参数 obj
	V2    string `gorm:"type:varchar(191)"` // 第三个参数 act
	V3    string `gorm:"type:varchar(191)"` // 第4个参数
	V4    string `gorm:"type:varchar(191)"` // 第5个参数
	V5    string `gorm:"type:varchar(191)"` // 第6个参数
}

// RoleLink 表示角色继承关系
type RoleLink struct {
	ID    uint   `gorm:"primaryKey"`        // 主键
	PType string `gorm:"type:varchar(191)"` //策略类型
	V0    string `gorm:"type:varchar(191)"` // 第一个参数 sub
	V1    string `gorm:"type:varchar(191)"` // 第二个参数 obj
}

func (AccessControlPolicy) TableName() string {
	return "casbin_rules"
}

func (RoleLink) TableName() string {
	return "casbin_rules_links"
}
