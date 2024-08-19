package reportController

import (
	"BBS/app/models"
	"BBS/app/services/reportService"
	"BBS/app/utils"
	"strconv"

	"github.com/gin-gonic/gin"
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

	var data ResponseData
	data.ReportList = list
	utils.JsonSuccessResponse(c, data)
}
