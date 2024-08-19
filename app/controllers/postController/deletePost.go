package postController

import (
	"BBS/app/services/postService"
	"BBS/app/utils"

	"github.com/gin-gonic/gin"
)

type DeleteData struct {
	Post uint `json:"post_id"`
	User uint `json:"user_id"`
}

func DeletePost(c *gin.Context) {
	var data DeleteData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	var user uint
	user, err = postService.GetUserByPostID(data.Post)
	if user != data.User {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	err = postService.DeletePost(data.Post)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
