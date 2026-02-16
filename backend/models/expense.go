package models

import "time"

type Expense struct {
	// CHANGE: string -> uint
	ID          uint           `json:"id" gorm:"primaryKey"`
	TripID      uint           `json:"trip_id"` // Must match Trip.ID type
	Description string         `json:"description"`
	Amount      float64        `json:"amount"`
	PaidBy      uint           `json:"paid_by"` // Must match User.ID type
	
	LocationName string        `json:"location_name"`
	Latitude     float64       `json:"latitude"`
	Longitude    float64       `json:"longitude"`

	Splits      []ExpenseSplit `json:"splits" gorm:"foreignKey:ExpenseID"`
	Date        time.Time      `json:"date"`
}