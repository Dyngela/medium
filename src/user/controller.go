package user

import "github.com/gin-gonic/gin"

func UserController(router *gin.Engine) {
	unauthenticated := router.Group("/api/v1/user/")
	{
		unauthenticated.POST("login", login)
		unauthenticated.PUT("register", register)
	}
}
