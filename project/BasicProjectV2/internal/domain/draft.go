package domain

type Draft struct {
	ID         string `json:"id"`
	AuthorName string `json:"authorName"`
	AuthorID   string `json:"authorID"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Status     string `json:"status"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
	DeletedAt  string `json:"deletedAt"`
}

type AddDraftReq struct {
	AuthorName string `json:"authorName"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Status     string `json:"status"`
}

type GetDraftReq struct {
	DraftID  string `json:"draftID"`
	AuthorID string `json:"authorID"`
}

type UpdateDraftReq struct {
	DraftID  string `json:"draftID"`
	AuthorID string `json:"authorID"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Status   string `json:"status"`
}

type DeleteDraftReq struct {
	DraftID string `json:"draftID"`
}
