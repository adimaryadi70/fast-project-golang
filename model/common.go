package model

type PagingModel struct {
	Page     string `json:"page"`
	PageSize string `json:"pageSize"`
	TotalRows string `json:"totalRows"`
	Data     interface{} `json:"data"`
}