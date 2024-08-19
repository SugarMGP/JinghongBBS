package reportController

import "github.com/gin-gonic/gin"

type ReportData struct{
	UserID uint `json:"user_id"`
	PostID uint `json:"post_id"`
	Reason string `json:"reason"`
}

func NewReport(c *gin.Context){
	
}