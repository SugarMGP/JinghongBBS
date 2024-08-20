package postService

import (
	"BBS/app/models"
	"BBS/config/database"
)

func NewPost(post models.Post) error {
	result := database.DB.Create(&post)
	return result.Error
}

func GetAllPosts() ([]models.Post, error) {
	var posts []models.Post
	result := database.DB.Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	return posts, nil
}

func GetPostByID(id uint) (*models.Post, error) {
	var post models.Post
	result := database.DB.Where("id = ?", id).First(&post)
	if result.Error != nil {
		return nil, result.Error
	}
	return &post, nil
}

func DeletePost(id uint) error {
	result := database.DB.Where("id = ?", id).Delete(&models.Post{})
	return result.Error
}

func EditPost(id uint, content string) error {
	result := database.DB.Where("id = ?", id).First(&models.Post{}).Update("content", content)
	return result.Error
}
