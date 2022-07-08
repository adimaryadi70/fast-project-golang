package model

import (
	"time"
)

type Transaction struct {
	ID             uint    `json:"id" gorm:"primary_key"`
	Po_number      string  `json:"po_number"`
	Po_date        time.Time  `json:"po_date"`
	Po_price_total int     `json:"po_price_total"`
	Po_cost_total  int     `json:"po_cost_total"`
}

type TransactionDetails struct {
	ID uint `json:"id" gorm:"primary_key"`
	transaction_id int `json:"transaction_id"`
	item_id int `json:"item_id"`
	item_quantity string `json:"item_quantity"`
	item_price int `json:"item_price"`
	item_cost int `json:"item_cost"`
}