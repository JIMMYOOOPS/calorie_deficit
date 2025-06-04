package user

import (
    "time"
)

// TODO: Define your User model(s) here
type User struct {
    // Add fields for the User model
    ID   int    `json:"id"`
    CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}
