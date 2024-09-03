package midwares

import (
	"BBS/app/utils"

	"github.com/gin-gonic/gin"
)

func JWTAuth(c *gin.Context) {
	// 通过 header 中的 token 来认证
	token := c.Request.Header.Get("token")
	if token == "" {
		utils.JsonErrorResponse(c, 200502, "无权限访问")
		c.Abort()
		return
	}

	id, err := utils.ExtractToken(token)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		c.Abort()
		return
	}

	// 将 user_id 重新写入 gin.Context 对象中
	c.Set("user_id", id)
}
