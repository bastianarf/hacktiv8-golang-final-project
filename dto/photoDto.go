package dto

type Photo struct {
	Title    string `validate:"required"`
	Caption  string
	PhotoUrl string `validate:"required,url" json:"photo_url"`
}
