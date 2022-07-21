package model

import "time"

// Model 基类
type Model struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at "`
	// * 代表 null
	DeletedAt *time.Time `json:"deleted_at"`
}
