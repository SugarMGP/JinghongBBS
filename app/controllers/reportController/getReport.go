package reportController

import (
	"BBS/app/models"
	"BBS/app/services/reportService"
	"BBS/app/services/userService"
	"BBS/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ResponseData struct {
	ReportList []models.Report `json:"report_list"`
}

func GetReport(c *gin.Context) {
	var err error
	id := c.GetUint("user_id")

	// 获取举报列表
	list, err := reportService.GetReports(id)
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

func GetAllReportsUnhandled(c *gin.Context) {
	id := c.GetUint("user_id")

	// 获取用户信息并检查是否为管理员
	user, err := userService.GetUserByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.JsonErrorResponse(c, 200506, "用户不存在")
		} else {
			utils.JsonInternalServerErrorResponse(c)
		}
		return
	}
	if user.UserType != 2 {
		utils.JsonErrorResponse(c, 200502, "用户不是管理员")
		return
	}

	// 获取未处理的举报列表
	list, err := reportService.GetAllReportsUnhandled()
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	var data ResponseData
	data.ReportList = list
	utils.JsonSuccessResponse(c, data)
}
