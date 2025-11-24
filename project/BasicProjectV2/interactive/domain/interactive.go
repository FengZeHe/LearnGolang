package domain

type Interactive struct {
	ID           string `json:"id" gorm:"primary_key;autoIncrement;"`
	Aid          string `json:"aid"`
	ReadCount    int    `json:"readCount" gorm:"column:read_count"`
	LikeCount    int    `json:"likeCount" gorm:"column:like_count"`
	CollectCount int    `json:"collectCount" gorm:"column:collect_count"`
	CTime        string `json:"-" gorm:"column:ctime"`
	UTime        string `json:"-" gorm:"column:utime"`
}

type ArticleStatus struct {
	Liked     bool `json:"liked"`
	Collected bool `json:"collected"`
}

type AddReadCountReq struct {
	Aid string `json:"aid"`
}
