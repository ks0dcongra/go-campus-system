package middleware

import (
	"net/http"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

const userkey = "session_id"

func SetSession() gin.HandlerFunc{
	store := cookie.NewStore([]byte(userkey))
	return sessions.Sessions("mysession", store)
}

// User Auth Session Middle
func AuthSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		session:= sessions.Default(c)
		
		sessionID := session.Get(userkey)
		if sessionID == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message" : "此頁面需要登入",
			})
			return
		}
	}
}

// Save Session for User
func SaveSession(c *gin.Context, userID int){
	session := sessions.Default(c)
	session.Set(userkey,userID)
	session.Save()
}

// Clear Session for User
func ClearSession(c *gin.Context){
	session := sessions.Default(c)
	session.Clear()
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
