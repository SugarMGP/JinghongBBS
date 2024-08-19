package postController

import (
	"BBS/app/models"
	"BBS/app/services/postService"
	"BBS/app/utils"

	"github.com/gin-gonic/gin"
)

type PostData struct {
	Content string `json:"content" binding:"required"`
	User    uint   `json:"user_id" binding:"required"`
}

func Post(c *gin.Context) {
	var data PostData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	err = postService.NewPost(models.Post{
		Content: data.Content,
		User:    data.User,
	})
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
