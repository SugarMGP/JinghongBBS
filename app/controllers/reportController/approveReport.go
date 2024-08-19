package reportController

import (
	"BBS/app/services/postService"
	"BBS/app/services/reportService"
	"BBS/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ApproveData struct {
	UserID   uint `json:"user_id"`
	PostID   uint `json:"post_id"`
	Approval uint `json:"approval"`
}

func ApproveReport(c *gin.Context) {
	var data ApproveData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	report, err := reportService.GetReportByID(data.UserID, data.PostID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.JsonErrorResponse(c, 200506, "举报不存在")
			return
		} else {
			utils.JsonInternalServerErrorResponse(c)
			return
		}
	}
	report.Status = data.Approval

	if data.Approval == 1 {
		err = postService.DeletePost(data.PostID)
		if err != nil {
			utils.JsonInternalServerErrorResponse(c)
			return
		}
	}

	utils.JsonSuccessResponse(c, nil)
}
