package models

import "time"

// Account 账号
type Account struct {
	ID       uint64    `json:"id" gorm:"primaryKey"`
	User     string    `json:"user"`
	Password string    `json:"password"`
	Mail     string    `json:"mail" gorm:"size:50"`
	CreateAt time.Time `json:"create_at"`
	Locked   bool      `json:"locked"`
}
