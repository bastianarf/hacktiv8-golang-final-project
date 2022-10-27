package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"example.id/mygram/controllers/responses"
	"example.id/mygram/database"
	"example.id/mygram/dto"
	"example.id/mygram/utils/token"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreatePhoto(ctx *gin.Context) {
	var newPhoto dto.Photo
	if err := ctx.ShouldBindJSON(&newPhoto); err != nil {
		abortBadRequest(err, ctx)
		return
	}
	if err := validate.Struct(&newPhoto); err != nil {
		validationAbort(err, ctx)
		return
	}
	userID, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ID, err := database.CreatePhoto(userID, &newPhoto)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, responses.CreatePhoto{
		Photo: responses.Photo{
			ID:       ID,
			Title:    newPhoto.Title,
			Caption:  newPhoto.Caption,
			PhotoUrl: newPhoto.PhotoUrl,
			UserID:   userID,
		},
		CreatedAt: time.Now(),
	})
}

func GetAllPhotos(ctx *gin.Context) {
	photos, err := database.GetAllPhotos()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	photosResponse := make([]responses.GetPhoto, len(photos))
	userDtos := make(map[uint]dto.UserUpdate)
	for i, photo := range photos {
		photosResponse[i].Set(photo)
		userDto, ok := userDtos[photo.UserID]
		if !ok {
			userDto, err = database.GetUsernameAndEmail(photo.UserID)
			if err != nil {
				ctx.AbortWithError(http.StatusInternalServerError, err)
				return
			}
			userDtos[photo.UserID] = userDto
		}
		photosResponse[i].User = userDto
	}
	ctx.JSON(http.StatusOK, photosResponse)
}

func UpdatePhoto(ctx *gin.Context) {
	photoID := ctx.Param("photoId")
	parsedID, err := strconv.ParseUint(photoID, 10, 0)
	if err != nil {
		abortBadRequest(err, ctx)
		return
	}
	var photoDto dto.Photo
	if err := ctx.ShouldBindJSON(&photoDto); err != nil {
		abortBadRequest(err, ctx)
		return
	}
	if err := validate.Struct(&photoDto); err != nil {
		validationAbort(err, ctx)
		return
	}
	userID, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	updatedAt, err := database.UpdatePhoto(uint(parsedID), userID, &photoDto)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, responses.ErrorMessage{
				ErrorMessage: fmt.Sprintf("Photo with ID %d is not found.", parsedID),
			})
			return
		}
		if errors.Is(err, database.ErrIllegalUpdate) {
			ctx.AbortWithStatusJSON(http.StatusForbidden, responses.ErrorMessage{
				ErrorMessage: err.Error(),
			})
			return
		}
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, responses.UpdatePhoto{
		Photo: responses.Photo{
			ID:       uint(parsedID),
			Title:    photoDto.Title,
			Caption:  photoDto.Caption,
			PhotoUrl: photoDto.PhotoUrl,
			UserID:   userID,
		},
		UpdatedAt: updatedAt,
	})
}

func DeletePhoto(ctx *gin.Context) {
	photoID := ctx.Param("photoId")
	parsedID, err := strconv.ParseUint(photoID, 10, 0)
	if err != nil {
		abortBadRequest(err, ctx)
		return
	}
	userID, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if err := database.DeletePhoto(uint(parsedID), userID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, responses.ErrorMessage{
				ErrorMessage: fmt.Sprintf("Photo with ID %d is not found.", parsedID),
			})
			return
		}
		if errors.Is(err, database.ErrIllegalUpdate) {
			ctx.AbortWithStatusJSON(http.StatusForbidden, responses.ErrorMessage{
				ErrorMessage: err.Error(),
			})
			return
		}
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, responses.Message{
		Message: "Your photo has been successfully deleted",
	})
}
