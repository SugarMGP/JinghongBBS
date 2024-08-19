package postController

import (
	"BBS/app/services/postService"
	"BBS/app/utils"

	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {
	postList, err := postService.GetAllPosts()
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	utils.JsonSuccessResponse(c, postList)
}
