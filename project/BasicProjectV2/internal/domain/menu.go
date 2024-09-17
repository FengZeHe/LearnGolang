package domain

type Menu struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Path     string `json:"path"`
	Methods  string `json:"methods"`
	ParentID string `json:"parent_id"`
	OrderNo  string `json:"order_no"`
}
