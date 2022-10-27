package models

type Photo struct {
	Model
	Title    string `gorm:"not null" json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `gorm:"not null" json:"photo_url"`
	UserID   uint   `json:"user_id"`
}
