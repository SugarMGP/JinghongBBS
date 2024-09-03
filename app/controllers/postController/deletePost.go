package postController

import (
	"BBS/app/services/postService"
	"BBS/app/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeletePost(c *gin.Context) {
	var err error
	var postID int
	id := c.GetUint("user_id")

	postID, err = strconv.Atoi(c.Query("post_id"))
	if err != nil || postID < 0 {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	post, err := postService.GetPostByID(uint(postID))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.JsonErrorResponse(c, 200506, "帖子不存在")
		} else {
			utils.JsonInternalServerErrorResponse(c)
		}
		return
	}
	if post.User != id {
		utils.JsonErrorResponse(c, 200502, "请求的用户与发帖人不符")
		return
	}

	err = postService.DeletePost(uint(postID))
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
