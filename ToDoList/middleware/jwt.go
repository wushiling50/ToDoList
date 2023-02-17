package middleware

import (
	"main/ToDoList/pkg/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int = 200

		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {
			claim, err := utils.ParseToken(token)
			if err != nil {
				code = 403 // 无权限，token是无权限的，是假的
			} else if time.Now().Unix() > claim.ExpiresAt {
				code = 401 //Token无效
			}
		}

		if code != 200 {
			c.JSON(200, gin.H{
				"Status": code,
				"msg":    "Tolen解析错误",
			})
			c.Abort()
			return
		}
		c.Next()
	}

}
