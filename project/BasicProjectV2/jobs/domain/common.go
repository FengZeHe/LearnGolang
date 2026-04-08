package domain

type PageReq struct {
	PageIndex int `form:"pageIndex"`
	PageSize  int `form:"pageSize"`
}

type PageResp struct {
	PageIndex int         `json:"pageIndex"`
	PageSize  int         `json:"PageSize"`
	Total     int         `json:"total"`
	Data      interface{} `json:"data"`
}
