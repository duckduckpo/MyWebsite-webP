package models

import "time"

type Trip struct {
	// CHANGE: string -> uint
	ID          uint      `json:"id" gorm:"primaryKey"` 
	Name        string    `json:"name"`
	InviteCode  string    `json:"invite_code"`
	Currency    string    `json:"currency"`
	Members     []User    `json:"members" gorm:"many2many:trip_members;"`
	CreatedAt   time.Time `json:"created_at"`
}