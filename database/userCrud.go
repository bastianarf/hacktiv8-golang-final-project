package database

import (
	"log"
	"time"

	"example.id/mygram/dto"
	"example.id/mygram/models"
	"example.id/mygram/utils/token"
	"golang.org/x/crypto/bcrypt"
)

var ErrPasswordMismatch = bcrypt.ErrMismatchedHashAndPassword

func CreateUser(userRegister *dto.UserRegister) (ID uint, err error) {
	if db == nil {
		err = ErrDbNotStarted
		return
	}
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(userRegister.Password), 4)
	if err != nil {
		return
	}
	newUser := models.User{
		Username: userRegister.Username,
		Email:    userRegister.Email,
		Password: string(passwordBytes),
		Age:      userRegister.Age,
		Model: models.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	err = db.Create(&newUser).Error
	if err != nil {
		return
	}
	ID = newUser.ID
	log.Println("User Created:", newUser)
	return
}

func UpdateUser(id uint, userDto *dto.UserUpdate) (models.User, error) {
	user, err := getUserWithoutPreload(id)
	if err != nil {
		return user, err
	}
	if userDto.Email != "" {
		user.Email = userDto.Email
	}
	if userDto.Username != "" {
		user.Username = userDto.Username
	}
	user.UpdatedAt = time.Now()
	err = db.Save(&user).Error
	log.Printf("User Updated: %+v\n", user)
	return user, err
}

func DeleteUserById(id uint) error {
	user, err := getUserWithoutPreload(id)
	if err != nil {
		return err
	}
	if err := db.Delete(&user, id).Error; err != nil {
		return err
	}
	log.Println("User with id", id, "has been successfully deleted.")
	return nil
}

func GenerateToken(userDto dto.UserLogin) (jwt string, err error) {
	user, err := getUserByEmail(userDto.Email)
	if err != nil {
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDto.Password))
	if err != nil {
		return
	}
	jwt, err = token.GenerateToken(user.ID)
	return
}

func GetUsernameAndEmail(id uint) (dto.UserUpdate, error) {
	userDto := dto.UserUpdate{}
	if db == nil {
		return userDto, ErrDbNotStarted
	}
	user := models.User{}
	if err := db.Select("username", "email").Take(&user, id).Error; err != nil {
		return userDto, err
	}
	userDto.Username = user.Username
	userDto.Email = user.Email
	log.Println("Get user dto:", userDto)
	return userDto, nil
}

func getUserByEmail(email string) (models.User, error) {
	user := models.User{}
	if db == nil {
		return user, ErrDbNotStarted
	}
	err := db.Model(&models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return user, err
	}
	log.Println("Get user by email:", user)
	return user, nil
}

func getUserWithoutPreload(id uint) (models.User, error) {
	user := models.User{}
	if db == nil {
		return user, ErrDbNotStarted
	}
	err := db.Model(&models.User{}).Take(&user, id).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
