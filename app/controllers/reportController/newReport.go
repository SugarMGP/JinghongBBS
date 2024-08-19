package reportController

import (
	"BBS/app/models"
	"BBS/app/services/postService"
	"BBS/app/services/reportService"
	"BBS/app/services/userService"
	"BBS/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	_, err = reportService.GetReportByID(data.PostID)
	if err == nil {
		utils.JsonErrorResponse(c, 200505, "该帖子已被举报")
		return
	} else if err != gorm.ErrRecordNotFound {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	// 获取帖子内容
	post, err := postService.GetPostByID(data.PostID)
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
		Content:  post.Content,
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
