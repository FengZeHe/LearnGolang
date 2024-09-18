package domain

type AddArticle struct {
	ID         string `json:"id"`
	AuthorName string `json:"authorName"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Status     string `json:"status"`
}
