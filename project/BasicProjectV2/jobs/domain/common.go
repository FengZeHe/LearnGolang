package domain

type PageReq struct {
	PageIndex int `form:"pageIndex"`
	PageSize  int `form:"PageSize"`
}

type PageResp struct {
	PageIndex int         `json:"pageIndex"`
	PageSize  int         `json:"PageSize"`
	Total     int         `json:"total"`
	Data      interface{} `json:"data"`
}
