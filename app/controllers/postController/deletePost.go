package postController

import (
	"BBS/app/services/postService"
	"BBS/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DeleteData struct {
	PostID uint `json:"post_id"`
	UserID uint `json:"user_id"`
}

func DeletePost(c *gin.Context) {
	var data DeleteData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	var user uint
	user, err = postService.GetUserByPostID(data.PostID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.JsonErrorResponse(c, 200506, "帖子不存在")
			return
		} else {
			utils.JsonInternalServerErrorResponse(c)
			return
		}
	}
	if user != data.UserID {
		utils.JsonErrorResponse(c, 200502, "请求的用户与发帖人不符")
		return
	}

	err = postService.DeletePost(data.PostID)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
