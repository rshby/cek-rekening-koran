package coresystem

import "time"

type CoreSystem struct {
	Id              int        `json:"id,omitempty"`
	AccountNumber   string     `json:"account_number,omitempty"`
	AccountName     string     `json:"account_name,omitempty"`
	Remark          string     `json:"remark,omitempty"`
	Dorc            string     `json:"dorc,omitempty"`
	Amount          float64    `json:"amount,omitempty"`
	TransactionDate *time.Time `json:"transaction_date,omitempty"`
	TransactionTime *time.Time `json:"transaction_time,omitempty"`
}
