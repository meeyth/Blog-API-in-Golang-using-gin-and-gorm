package database

import (
	"bloggify-api/models"
	"bloggify-api/utils"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func GetUsersFromDb(username string) ([]models.User, error) {
	var users []models.User

	result := db.Where("user_name LIKE ?", "%"+username+"%").Find(&users)

	return users, result.Error
}

func GetAUsersPostsFromDb(username string) ([]models.Blog, error) {
	var posts []models.Blog

	result := db.Where("creator = ?", username).Find(&posts)

	return posts, result.Error

}

func DeleteAccountFromDb(u *models.User, issuer string) error {
	db.Where("id = ?", issuer).Find(&u)
	result := db.Delete(&u)
	return result.Error
}

func UpdateAccountFromDb(body *ReqBody, issuer string) (models.User, error) {
	var user models.User

	// Check duplicate username

	db.Where("user_name = ?", body.UserName).Find(&user)
	if user.ID != 0 {
		return models.User{}, errors.New("username already exists, try another one")
	}

	db.Where("id = ?", issuer).First(&user)

	birthday, err := utils.CreateDate(body.Birthday)
	if err != nil {
		log.Fatal(err)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 14)

	if err != nil {
		log.Fatal(err)
	}

	// Updating posts table's creator field for this user.

	db.Model(&models.Blog{}).Where("creator = ?", user.UserName).Update("creator", body.UserName)

	// Then updating the actual user.

	result := db.Model(&user).Updates(models.User{UserName: body.UserName, Password: hashedPassword, Birthday: birthday})

	return user, result.Error
}

var GetAccountDetailsFromDb = func(u *models.User, issuer string) { db.Where("id = ?", issuer).First(&u) }
