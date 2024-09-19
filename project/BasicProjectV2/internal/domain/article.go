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
	AuthorName string `json:"authorName"`
	AuthorID   string `json:"authorId"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	DeletedAt  string `json:"deleted_at"`
}

type ArticleResponse struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	AuthorName string `json:"authorName"`
	CreatedAt  string `json:"created_at"`
}
