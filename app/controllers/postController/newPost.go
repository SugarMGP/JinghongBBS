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
}

func NewPost(c *gin.Context) {
	id := c.GetUint("user_id")
	var data PostData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	_, err = userService.GetUserByID(id)
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
		User:    id,
	})
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
