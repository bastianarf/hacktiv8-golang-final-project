package models

import (
	"time"
)

type Model struct {
	ID        uint      `gorm:"PrimaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
