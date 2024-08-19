package reportController

import (
	"BBS/app/models"
	"BBS/app/services/postService"
	"BBS/app/services/reportService"
	"BBS/app/services/userService"
	"BBS/app/utils"

	"github.com/gin-gonic/gin"
)

type ReportData struct {
	UserID uint   `json:"user_id"`
	PostID uint   `json:"post_id"`
	Reason string `json:"reason"`
}

func NewReport(c *gin.Context) {
	var data ReportData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	// 获取帖子内容
	content, err := postService.GetPostContentByID(data.PostID)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	user, err := userService.GetUserByID(data.UserID)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	err = reportService.NewReport(models.Report{
		User:     data.UserID,
		Post:     data.PostID,
		Content:  content,
		Reason:   data.Reason,
		Status:   0,
		Username: user.Username,
	})
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
