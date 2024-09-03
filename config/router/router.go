package router

import (
	"BBS/app/controllers/postController"
	"BBS/app/controllers/reportController"
	"BBS/app/controllers/userController"
	"BBS/app/midwares"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/user/login", userController.Login)
		api.POST("/user/reg", userController.Register)
		api.GET("/student/post", postController.GetAllPosts)

		api.POST("/student/post", midwares.JWTAuth, postController.NewPost)
		api.DELETE("/student/post", midwares.JWTAuth, postController.DeletePost)
		api.PUT("/student/post", midwares.JWTAuth, postController.EditPost)

		api.POST("/student/report-post", midwares.JWTAuth, reportController.NewReport)
		api.GET("/student/report-post", midwares.JWTAuth, reportController.GetReport)
		api.GET("/admin/report", midwares.JWTAuth, reportController.GetAllReportsUnhandled)
		api.POST("/admin/report", midwares.JWTAuth, reportController.ApproveReport)
	}
}
