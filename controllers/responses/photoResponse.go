package responses

import (
	"time"

	"example.id/mygram/dto"
	"example.id/mygram/models"
)

type Photo struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserID   uint   `json:"user_id"`
}

type CreatePhoto struct {
	Photo
	CreatedAt time.Time `json:"created_at"`
}

type GetPhoto struct {
	models.Photo
	User dto.UserUpdate
}

type UpdatePhoto struct {
	Photo
	UpdatedAt time.Time `json:"updated_at"`
}

func (getPhoto *GetPhoto) Set(photo models.Photo) {
	getPhoto.ID = photo.ID
	getPhoto.CreatedAt = photo.CreatedAt
	getPhoto.UpdatedAt = photo.UpdatedAt
	getPhoto.Title = photo.Title
	getPhoto.Caption = photo.Caption
	getPhoto.PhotoUrl = photo.PhotoUrl
	getPhoto.UserID = photo.UserID
}
