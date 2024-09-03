package userController

import (
	"BBS/app/models"
	"BBS/app/services/userService"
	"BBS/app/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginData struct {
	User     string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	UserType uint   `json:"user_type"`
	Token    string `json:"token"`
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
	err = userService.VerifyPassword(data.Password, user.Password)
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			utils.JsonErrorResponse(c, 200507, "密码错误")
		} else {
			utils.JsonInternalServerErrorResponse(c)
		}
		return
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	response := LoginResponse{
		UserType: user.UserType,
		Token:    token,
	}

	// 返回用户信息
	utils.JsonSuccessResponse(c, response)
}
