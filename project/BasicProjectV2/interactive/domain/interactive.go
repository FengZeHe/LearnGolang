package domain

type Interactive struct {
	ID           string `json:"id"`
	Aid          string `json:"aid"`
	ReadCount    int    `json:"readCount"`
	LikeCount    int    `json:"likeCount"`
	CollectCount int    `json:"collectCount"`
}

type ArticleStatus struct {
	Liked     bool `json:"liked"`
	Collected bool `json:"collected"`
}
