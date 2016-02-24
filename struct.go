package main

import (
	"github.com/shopspring/decimal"
)

type User struct {
	Id      int             `json:"id"`
	Name    string          `json:"name"`
	Balance decimal.Decimal `json:"balance"`
}

type BalanceUpdate struct {
	User   int             `json:"user"`
	Amount decimal.Decimal `json:"amount"`
}

type BalanceTransfer struct {
	From   int             `json:"from"`
	To     int             `json:"to"`
	Amount decimal.Decimal `json:"amount"`
}
