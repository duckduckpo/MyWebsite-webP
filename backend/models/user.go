package models

import (
	"time"
)

// User Model
type User struct {
	// CHANGE: ID is now uint (Auto-incrementing Integer)
	ID           uint      `json:"id" gorm:"primaryKey"` 
	Name         string    `json:"name"`
	Email        string    `json:"email" gorm:"unique"`  
	Password     []byte    `json:"-"`                    
	AvatarURL    string    `json:"avatar_url"`
	HomeCurrency string    `json:"home_currency"`
	CreatedAt    time.Time `json:"created_at"`
}