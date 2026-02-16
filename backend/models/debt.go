package models

type Debt struct {
	// CHANGE: string -> uint
	From   uint    `json:"from"`
	To     uint    `json:"to"`
	Amount float64 `json:"amount"`
}