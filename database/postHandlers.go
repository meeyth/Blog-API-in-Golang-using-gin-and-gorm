package database

import (
	"errors"
	"social-media/models"
)

func GetAllPostsFromDb() ([]models.Post, error) {
	var posts []models.Post
	result := db.Find(&posts)

	return posts, result.Error
}

func GetPostByTitleFromDb(title string) ([]models.Post, error) {
	var posts []models.Post
	db.Where("title LIKE ?", "%"+title+"%").Find(&posts)

	if len(posts) == 0 {
		return nil, errors.New("posts with this title don't exist")
	}
	return posts, nil
}

func InsertAPostIntoDB(post *models.Post) error {
	result := db.Create(post)
	return result.Error
}

func DeleteAPostByTitleFromDB(post *models.Post, title string, creator string) error {
	db.Where("title = ?", title).First(&post)

	if post.ID == 0 {
		return errors.New("post with this title doesn't exist")
	}

	if post.Creator != creator {
		return errors.New("this post doesn't belong to you")
	}

	db.Delete(&post)
	return nil
}

func UpdateAPostByTitleFromDb(toUpdateData *models.Post, title string) (models.Post, error) {
	var updatedPost models.Post
	db.Where("title = ?", title).First(&updatedPost)

	if updatedPost.ID == 0 {
		return updatedPost, errors.New("post with this title doesn't exist")
	}

	db.Model(&updatedPost).Updates(models.Post{Title: toUpdateData.Title, Description: toUpdateData.Description, LikeCount: toUpdateData.LikeCount})
	return updatedPost, nil
}
