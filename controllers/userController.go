package controllers

import (
	"errors"
	"net/http"

	"example.id/mygram/controllers/responses"
	"example.id/mygram/database"
	"example.id/mygram/dto"
	"example.id/mygram/utils/token"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

func RegisterUser(ctx *gin.Context) {
	var newUser dto.UserRegister
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		abortBadRequest(err, ctx)
		return
	}
	if err := validate.Struct(&newUser); err != nil {
		validationAbort(err, ctx)
		return
	}
	ID, err := database.CreateUser(&newUser)
	if err != nil {
		var perr *pgconn.PgError
		if ok := errors.As(err, &perr); ok {
			if perr.Code == uniqueViolationErr {
				ctx.AbortWithStatusJSON(http.StatusOK, responses.Message{
					Message: "The email or username is already registered. If it is yours, do login instead.",
				})
				return
			}
		}
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, responses.UserRegister{
		Age:      newUser.Age,
		Email:    newUser.Email,
		ID:       ID,
		Username: newUser.Username,
	})
}

func LoginUser(ctx *gin.Context) {
	var userLogin dto.UserLogin
	if err := ctx.ShouldBindJSON(&userLogin); err != nil {
		abortBadRequest(err, ctx)
		return
	}
	if err := validate.Struct(&userLogin); err != nil {
		validationAbort(err, ctx)
		return
	}
	jwt, err := database.GenerateToken(userLogin)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, database.ErrPasswordMismatch) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.ErrorMessage{
				ErrorMessage: "Email or password is incorrect.",
			})
			return
		}
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, responses.UserLogin{
		Token: jwt,
	})
}

func UpdateUser(ctx *gin.Context) {
	var userDto dto.UserUpdate
	if err := ctx.ShouldBindJSON(&userDto); err != nil {
		abortBadRequest(err, ctx)
		return
	}
	if err := validate.Struct(&userDto); err != nil {
		validationAbort(err, ctx)
		return
	}
	userID, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	user, err := database.UpdateUser(userID, &userDto)
	if err != nil {
		var perr *pgconn.PgError
		if ok := errors.As(err, &perr); ok {
			if perr.Code == uniqueViolationErr {
				ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
					messageStr: "The email or username is already registered.",
				})
				return
			}
		}
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, responses.UserUpdate{
		ID:        user.ID,
		Email:     user.Email,
		Username:  user.Username,
		Age:       user.Age,
		UpdatedAt: user.UpdatedAt,
	})
}

func DeleteUser(ctx *gin.Context) {
	userID, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if err := database.DeleteUserById(userID); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, responses.Message{
		Message: "Your account has been successfully deleted",
	})
}
