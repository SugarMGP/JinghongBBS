package router

import (
	"BBS/app/controllers/postController"
	"BBS/app/controllers/userController"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	const pre = "/api"

	api := r.Group(pre)
	{
		api.POST("/user/login", userController.Login)
		api.POST("/user/reg", userController.Register)

		api.POST("/student/post", postController.NewPost)
		api.GET("/student/post", postController.GetPosts)
		api.DELETE("/student/post", postController.DeletePost)
	}
}
