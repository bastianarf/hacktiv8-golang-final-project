package models

type Comment struct {
	Model
	UserID  uint
	PhotoID uint
	Message string `gorm:"not null;type:varchar(8192)"`
}
