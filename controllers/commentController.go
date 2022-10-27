package controllers

import (
	"net/http"

	"example.id/mygram/controllers/responses"
	"example.id/mygram/database"
	"example.id/mygram/dto"
	"example.id/mygram/utils/token"
	"github.com/gin-gonic/gin"
)

func CreateComment(ctx *gin.Context) {
	var newComment dto.Comment
	if err := ctx.ShouldBindJSON(&newComment); err != nil {
		abortBadRequest(err, ctx)
		return
	}
	if err := validate.Struct(&newComment); err != nil {
		validationAbort(err, ctx)
		return
	}
	userID, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	comment, err := database.CreateComment(userID, &newComment)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, responses.CreateComment{
		ID:        comment.ID,
		Message:   comment.Message,
		PhotoID:   comment.PhotoID,
		UserID:    userID,
		CreatedAt: comment.CreatedAt,
	})
}
