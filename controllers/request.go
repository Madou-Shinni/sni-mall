package controllers

import "github.com/gin-gonic/gin"

const (
	ContextUserIdKey = "userId" // adminAuth.go中设置的上下文
)

// GetCurrentUserId 获取当前的用户id
func GetCurrentUserId(c *gin.Context) (userId int64) {
	uid, exists := c.Get(ContextUserIdKey)
	if !exists {
		// 上下文中不存在
		return
	}
	if userId, exists = uid.(int64); !exists {
		// 类型断言失败
		return
	}
	return
}
