package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

const userkey = "session_id"

func SetSession() gin.HandlerFunc {
	store := cookie.NewStore([]byte(userkey))
	return sessions.Sessions("mysession", store)
}

// User Auth Session Middle
func AuthSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 拿到我在 userApi := router.Group("user/api", session.SetSession())所暫存在Default變數的session
		session := sessions.Default(c)
		// 拿到我在session所暫存的ID
		sessionID := session.Get(userkey)
		// 如果我所暫存的變數不存在代表該人沒有登入過
		if sessionID == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "此頁面需要登入",
			})
			return
		}
	}
}

// Save Session for User
func SaveSession(c *gin.Context, userID int) {
	// 拿到我在 userApi := router.Group("user/api", session.SetSession())所暫存在Default變數的session
	session := sessions.Default(c)
	// 用session來暫存userID
	session.Set(userkey, userID)
	// 將當下登入的sesison存入伺服器端
	session.Save()
}

// Clear Session for User
func ClearSession(c *gin.Context) {
	// 拿到我在 userApi := router.Group("user/api", session.SetSession())所暫存在Default變數的session
	session := sessions.Default(c)
	// 將目前session清除
	session.Clear()
	// 將當下登出的sesison存入伺服器端
	session.Save()
}

// Get Session for User
func GetSession(c *gin.Context) int {
	session := sessions.Default(c)
	sessionID := session.Get(userkey)
	if sessionID == nil {
		return 0
	}
	return sessionID.(int)
}
