package database

import (
	"log"
	"time"

	"example.id/mygram/dto"
	"example.id/mygram/models"
)

func CreateSocialMedia(userID uint, socmedDto *dto.SocialMedia) (models.SocialMedia, error) {
	if db == nil {
		return models.SocialMedia{}, ErrDbNotStarted
	}
	newSocmed := models.SocialMedia{
		UserID:         userID,
		Name:           socmedDto.Name,
		SocialMediaUrl: socmedDto.SocialMediaUrl,
		Model: models.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	if err := db.Create(&newSocmed).Error; err != nil {
		return models.SocialMedia{}, err
	}
	log.Printf("Social Media Created: %+v\n", newSocmed)
	return newSocmed, nil
}

func GetAllSocialMedias() ([]models.SocialMedia, error) {
	socmeds := make([]models.SocialMedia, 1)
	if err := db.Model(&models.SocialMedia{}).Find(&socmeds).Error; err != nil {
		return nil, err
	}
	log.Println("All social medias is read from db.")
	return socmeds, nil
}

func GetSingleSocialMedia(socmedID uint) (models.SocialMedia, error) {
	socmed := models.SocialMedia{}
	if db == nil {
		return socmed, ErrDbNotStarted
	}
	err := db.Model(&models.SocialMedia{}).Take(&socmed, socmedID).Error
	return socmed, err
}

func UpdateSocialMedia(socmedID uint, socmedDto *dto.SocialMedia) (UpdatedAt time.Time, err error) {
	socmed, err := GetSingleSocialMedia(socmedID)
	if err != nil {
		return
	}
	socmed.Name = socmedDto.Name
	socmed.SocialMediaUrl = socmedDto.SocialMediaUrl
	socmed.UpdatedAt = time.Now()
	err = db.Save(&socmed).Error
	if err == nil {
		UpdatedAt = socmed.UpdatedAt
		log.Printf("Social Media Updated: %+v\n", socmed)
	} else {
		log.Printf("Failed Update Social Media %+v with %+v because of %q", socmed, socmedDto, err.Error())
	}
	return
}

func DeleteSocialMediaById(socmedID uint) error {
	socmed, err := GetSingleSocialMedia(socmedID)
	if err != nil {
		return err
	}
	if err := db.Delete(&socmed, socmedID).Error; err != nil {
		return err
	}
	log.Println("Social Media with ID", socmedID, "has been successfully deleted.")
	return nil
}
