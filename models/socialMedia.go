package models

type SocialMedia struct {
	Model
	Name           string `gorm:"not null;type:varchar(8192)"`
	SocialMediaUrl string `gorm:"not null;type:varchar(8192)"`
	UserID         uint
}
