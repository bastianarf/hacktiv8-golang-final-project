package database

import (
	"log"
	"time"

	"example.id/mygram/dto"
	"example.id/mygram/models"
)

func CreateComment(userID uint, commentDto *dto.Comment) (models.Comment, error) {
	if db == nil {
		return models.Comment{}, ErrDbNotStarted
	}
	newComment := models.Comment{
		UserID:  userID,
		PhotoID: commentDto.PhotoID,
		Message: commentDto.Message,
		Model: models.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	if err := db.Create(&newComment).Error; err != nil {
		return models.Comment{}, err
	}
	log.Printf("Photo Created: %+v\n", newComment)
	return newComment, nil
}

func GetAllComments() ([]models.Comment, error) {
	comments := make([]models.Comment, 1)
	if err := db.Model(&models.Comment{}).Find(&comments).Error; err != nil {
		return nil, err
	}
	log.Println("All comments is read from db.")
	return comments, nil
}

func GetSingleComment(commentID uint) (models.Comment, error) {
	comment := models.Comment{}
	if db == nil {
		return comment, ErrDbNotStarted
	}
	err := db.Model(&models.Comment{}).Take(&comment, commentID).Error
	return comment, err
}

func UpdateComment(commentID uint, messageDto *dto.CommentMessage) (UpdatedAt time.Time, err error) {
	comment, err := GetSingleComment(commentID)
	if err != nil {
		return
	}
	comment.Message = messageDto.Message
	comment.UpdatedAt = time.Now()
	err = db.Save(&comment).Error
	if err == nil {
		UpdatedAt = comment.UpdatedAt
		log.Printf("Comment Updated: %+v\n", comment)
	} else {
		log.Printf("Failed Update Comment %+v with %+v because of %q\n", comment, messageDto, err.Error())
	}
	return
}

func DeleteCommentById(commentID uint) error {
	comment, err := GetSingleComment(commentID)
	if err != nil {
		return err
	}
	if err := db.Delete(&comment, commentID).Error; err != nil {
		return err
	}
	log.Println("Comment with ID", commentID, "has been successfully deleted.")
	return nil
}
