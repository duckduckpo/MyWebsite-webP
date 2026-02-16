package models

type ExpenseSplit struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	// CHANGE: string -> uint
	ExpenseID uint    `json:"expense_id"` 
	UserID    uint    `json:"user_id"` 
	Amount    float64 `json:"amount"`
}