package database

import (
	"bloggify-api/models"
	"errors"
)

func GetPostsFromDb(title string) ([]models.Blog, error) {
	var posts []models.Blog

	result := db.Where("title LIKE ?", "%"+title+"%").Find(&posts)

	return posts, result.Error
}

func GetPostsOfAUserFromDb(username string) ([]models.Blog, error) {
	var posts []models.Blog

	result := db.Where("creator = ?", username).Find(&posts)

	return posts, result.Error
}

func InsertAPostIntoDB(Blog *models.Blog) error {
	result := db.Create(Blog)
	return result.Error
}

func DeleteAPostByTitleFromDB(Blog *models.Blog, title string, creator string) error {
	db.Where("title = ?", title).First(&Blog)

	if Blog.ID == 0 {
		return errors.New("blog with this title doesn't exist")
	}

	if Blog.Creator != creator {
		return errors.New("this Blog doesn't belong to you")
	}

	db.Delete(&Blog)
	return nil
}

func UpdateAPostByTitleFromDb(toUpdateData *models.Blog, title string) (models.Blog, error) {
	var updatedPost models.Blog
	db.Where("title = ?", title).First(&updatedPost)

	if updatedPost.ID == 0 {
		return updatedPost, errors.New("blog with this title doesn't exist")
	}

	db.Model(&updatedPost).Updates(models.Blog{Title: toUpdateData.Title, Description: toUpdateData.Description})
	return updatedPost, nil
}
