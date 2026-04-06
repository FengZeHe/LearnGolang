package domain

type PageReq struct {
	PageIndex int `form:"pageIndex"`
	pageSize  int `form:"pageSize"`
}

type PageResp struct {
	PageIndex int         `json:"pageIndex"`
	PageSize  int         `json:"pageSize"`
	Total     int         `json:"total"`
	Data      interface{} `json:"data"`
}
