package database

import "social-media/models"

func InsertAUserIntoDb(u *models.User) error {
	result := db.Create(u)

	return result.Error
}

var GetLoggedInUserFromDb = func(u *models.User, userName string) { db.Where("user_name = ?", userName).First(&u) }
