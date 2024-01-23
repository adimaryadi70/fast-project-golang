package model

type PagingModel struct {
	Page     string `json:"page"`
	PageSize string `json:"pageSize"`
	TotalPages string `json:"totalPages"`
	Data     interface{} `json:"data"`
}

type PayloadInquiry struct {
	KodeAgen      string `json:"kode_agen"`
	Token         string `json:"token"`
	TransactionId string `json:"transaction_id"`
}
type PayloadFlagging struct {
	Amount        float64 `json:"amount"`
	Status        string  `json:"status"`
	Token         string  `json:"token"`
	TransactionId string  `json:"transaction_id"`
}