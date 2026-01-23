package domain

type AddArticle struct {
	ID         string `json:"id"`
	AuthorName string `json:"authorName"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Status     string `json:"status"`
}

type Article struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Status     string `json:"status"`
	Read       int    `json:"read"`
	AuthorName string `json:"authorName"`
	AuthorID   string `json:"authorId"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	DeletedAt  string `json:"-"`
}

type ArticlesDAOResponse struct {
	PageIndex  int       `json:"pageIndex"`
	PageCount  int       `json:"pageCount"`
	TotalCount int64     `json:"totalCount"`
	Articles   []Article `json:"articles"`
}

type ArticleResponse struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Read       int    `json:"read"`
	AuthorName string `json:"authorName"`
	CreatedAt  string `json:"created_at"`
}

type ArticleRepoResponse struct {
	PageIndex  int               `json:"pageIndex"`
	PageCount  int               `json:"pageCount"`
	TotalCount int               `json:"totalCount"`
	Articles   []ArticleResponse `json:"articles"`
}

type QueryArticlesReq struct {
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
}

type QueryAuthorArticlesReq struct {
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
}
type QueryArticlesByIDReq struct {
	ID string `json:"id"`
}

type AddArticleCount struct {
	ID string `json:"id"`
}

type ArticleWithInteractive struct {
	ID           string `json:"id" `
	Title        string `json:"title"`
	ReadCount    int    `json:"readCount" `
	LikeCount    int    `json:"likeCount" `
	CollectCount int    `json:"collectCount" `
	CreatedAt    string `json:"createdAt"`
}

type ArticleWithScores struct {
	ID    string  `json:"id"`
	Title string  `json:"title"`
	Score float64 `json:"score"`
}
