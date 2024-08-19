package reportController

import (
	"BBS/app/models"
	"BBS/app/services/reportService"
	"BBS/app/services/userService"
	"BBS/app/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ResponseData struct {
	ReportList []models.Report `json:"report_list"`
}

func GetReport(c *gin.Context) {
	var userID int
	var err error

	userID, err = strconv.Atoi(c.Query("user_id"))
	if err != nil || userID < 0 {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	list, err := reportService.GetReports(uint(userID))
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	// 修改 Username 为空来屏蔽输出
	for i := range list {
		list[i].Username = ""
	}

	var data ResponseData
	data.ReportList = list
	utils.JsonSuccessResponse(c, data)
}

func GetAllReports(c *gin.Context) {
	var userID int
	var err error

	userID, err = strconv.Atoi(c.Query("user_id"))
	if err != nil || userID < 0 {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	var user *models.User
	user, err = userService.GetUserByID(uint(userID))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.JsonErrorResponse(c, 200506, "用户不存在")
			return
		} else {
			utils.JsonInternalServerErrorResponse(c)
			return
		}
	}
	if user.UserType != 2 {
		utils.JsonErrorResponse(c, 200502, "用户不是管理员")
		return
	}

	list, err := reportService.GetAllReports()
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	var data ResponseData
	data.ReportList = list
	utils.JsonSuccessResponse(c, data)
}
