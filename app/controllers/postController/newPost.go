package postController

import (
	"BBS/app/models"
	"BBS/app/services/postService"
	"BBS/app/services/userService"
	"BBS/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostData struct {
	Content string `json:"content" binding:"required"`
	User    uint   `json:"user_id" binding:"required"`
}

func NewPost(c *gin.Context) {
	var data PostData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	_, err = userService.GetUserByID(data.User)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.JsonErrorResponse(c, 200506, "用户不存在")
		} else {
			utils.JsonInternalServerErrorResponse(c)
		}
		return
	}

	// 新建帖子
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
