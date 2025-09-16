package models

import "time"

type Transaction struct {
	ID        uint      `gorm:"primaryKey" json:"transaction_id"`
	UserID    string    `json:"user_id"`
	Points    int       `json:"points"`
	DeviceID  string    `json:"device_id"`
	Timestamp time.Time `json:"timestamp"`
}
