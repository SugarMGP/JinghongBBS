package router

import (
	"BBS/app/controllers/postController"
	"BBS/app/controllers/reportController"
	"BBS/app/controllers/userController"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/user/login", userController.Login)
		api.POST("/user/reg", userController.Register)

		api.POST("/student/post", postController.NewPost)
		api.GET("/student/post", postController.GetAllPosts)
		api.DELETE("/student/post", postController.DeletePost)
		api.PUT("/student/post", postController.EditPost)

		api.POST("/student/report-post", reportController.NewReport)
		api.GET("/student/report-post", reportController.GetReport)
		api.GET("/admin/report", reportController.GetAllReportsUnhandled)
	}
}
