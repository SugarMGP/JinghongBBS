package postController

import (
	"BBS/app/services/postService"
	"BBS/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type EditData struct {
	PostID  uint   `json:"post_id"`
	UserID  uint   `json:"user_id"`
	Content string `json:"content"`
}

func EditPost(c *gin.Context) {
	var data EditData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	// 检查请求用户是否为发帖人
	post, err := postService.GetPostByID(data.PostID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.JsonErrorResponse(c, 200506, "帖子不存在")
		} else {
			utils.JsonInternalServerErrorResponse(c)
		}
		return
	}
	if post.User != data.UserID {
		utils.JsonErrorResponse(c, 200502, "请求的用户与发帖人不符")
		return
	}

	// 编辑帖子
	err = postService.EditPost(data.PostID, data.Content)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
