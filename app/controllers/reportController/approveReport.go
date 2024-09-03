package reportController

import (
	"BBS/app/services/postService"
	"BBS/app/services/reportService"
	"BBS/app/services/userService"
	"BBS/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ApproveData struct {
	PostID   uint `json:"post_id"`
	Approval uint `json:"approval"`
}

func ApproveReport(c *gin.Context) {
	id := c.GetUint("user_id")
	var data ApproveData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

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

	// 设置举报状态
	err = reportService.SetReportStatus(data.PostID, data.Approval)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.JsonErrorResponse(c, 200506, "举报不存在")
		} else {
			utils.JsonInternalServerErrorResponse(c)
		}
		return
	}

	// 若举报被批准则删除帖子
	if data.Approval == 1 {
		err = postService.DeletePost(data.PostID)
		if err != nil {
			utils.JsonInternalServerErrorResponse(c)
			return
		}
	}

	utils.JsonSuccessResponse(c, nil)
}
