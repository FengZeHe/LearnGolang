package domain

type Menu struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"size:255;not null"`
	Path     string `gorm:"size:255;not null"`
	ParentID *uint  `gorm:"default:null"`
	OrderNo  int    `gorm:"not null"`
}
