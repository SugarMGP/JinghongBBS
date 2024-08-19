package userController

import (
	"BBS/app/models"
	"BBS/app/services/userService"
	"BBS/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginData struct {
	User     string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var data LoginData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	// 判断用户是否存在并获取用户信息
	var user *models.User
	user, err = userService.GetUserByUsername(data.User)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.JsonErrorResponse(c, 200506, "用户不存在")
		} else {
			utils.JsonInternalServerErrorResponse(c)
		}
		return
	}

	// 判断密码是否正确
	if data.Password != user.Password {
		utils.JsonErrorResponse(c, 200507, "密码错误")
		return
	}

	// 返回用户信息
	utils.JsonSuccessResponse(c, user)
}
