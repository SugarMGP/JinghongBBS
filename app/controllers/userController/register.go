package userController

import (
	"BBS/app/models"
	"BBS/app/services/userService"
	"BBS/app/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RegisterData struct {
	Username string `json:"username" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	UserType uint   `json:"user_type" binding:"required"`
}

func Register(c *gin.Context) {
	var data RegisterData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	// 判断用户类型
	if data.UserType != 1 && data.UserType != 2 {
		utils.JsonErrorResponse(c, 200504, "用户类型错误")
		return
	}

	// 判断账号是否为纯数字
	_, err = strconv.Atoi(data.Username)
	if err != nil {
		utils.JsonErrorResponse(c, 200502, "用户名必须为纯数字")
		return
	}

	// 判断密码长度
	length := len([]rune(data.Password))
	if length < 8 || length > 16 {
		utils.JsonErrorResponse(c, 200503, "密码长度必须在8-16位")
		return
	}

	// 判断用户是否已经注册
	_, err = userService.GetUserByUsername(data.Username)
	if err == nil {
		utils.JsonErrorResponse(c, 200505, "用户名已存在")
		return
	} else if err != gorm.ErrRecordNotFound {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	// 注册用户
	err = userService.Register(models.User{
		Username: data.Username,
		Name:     data.Name,
		Password: data.Password,
		UserType: data.UserType,
	})
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
