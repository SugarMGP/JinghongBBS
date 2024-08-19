package reportController

import (
	"BBS/app/models"
	"BBS/app/services/reportService"
	"BBS/app/utils"

	"github.com/gin-gonic/gin"
)

type GetData struct {
	UserID uint `json:"user_id"`
}

type ResponseData struct {
	ReportList []models.Report `json:"report_list"`
}

func GetReport(c *gin.Context) {
	var getData GetData
	err := c.ShouldBindJSON(&getData)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	list, err := reportService.GetReports(getData.UserID)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	var responseData ResponseData
	responseData.ReportList = list
	utils.JsonSuccessResponse(c, responseData)
}
