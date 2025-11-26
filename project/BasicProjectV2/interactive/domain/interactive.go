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
type LikeReq struct {
	Aid  string `json:"aid"`
	Like int    `json:"like"`
}

type LikeRecord struct {
	ID    string `json:"id" gorm:"primary_key;autoIncrement;"`
	Aid   string `json:"aid" gorm:"aid"`
	Uid   string `json:"uid" gorm:"uid"`
	Like  int    `json:"like" gorm:"like"`
	CTime string `json:"-" gorm:"column:ctime"`
	UTime string `json:"-" gorm:"column:utime"`
}

type CollectReq struct {
	Aid     string `json:"aid"`
	Collect int    `json:"collect"`
}

type CollectRecord struct {
	ID        string `json:"id" gorm:"primary_key;autoIncrement;"`
	Aid       string `json:"aid" grom:"aid"`
	Uid       string `json:"uid" gorm:"uid"`
	Collected int    `json:"collected" grom:"collected"`
	Ctime     string `json:"-" gorm:"column:ctime"`
	Utime     string `json:"-" gorm:"column:utime"`
}

type InteractiveStatus struct {
	Collected int `json:"collected"`
	Liked     int `json:"liked"`
}

type InteractiveResp struct {
	Collected    int `json:"collected"`
	Liked        int `json:"liked"`
	ReadCount    int `json:"readCount"`
	LikeCount    int `json:"likeCount"`
	CollectCount int `json:"collectCount"`
}
