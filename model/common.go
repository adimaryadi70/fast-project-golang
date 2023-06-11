package model

type PagingModel struct {
	Page     string `json:"page"`
	PageSize string `json:"pageSize"`
	TotalPages string `json:"totalPages"`
	Data     interface{} `json:"data"`
}