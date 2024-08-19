package postController

import (
	"BBS/app/models"
	"BBS/app/services/postService"
	"BBS/app/utils"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	PostList []models.Post `json:"post_list"`
}

func GetPosts(c *gin.Context) {
	list, err := postService.GetAllPosts()
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	var data ResponseData
	data.PostList = list
	utils.JsonSuccessResponse(c, data)
}