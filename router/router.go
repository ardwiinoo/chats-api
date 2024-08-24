package router

import (
	"chats-api/internal/user"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func InitRouter(userHandler *user.Handler) {
	router = gin.Default()

	router.POST("/api/users/signup", userHandler.CreateUser)
	router.POST("/api/users/signin", userHandler.Login)
	router.GET("/api/users/logout", userHandler.Logout)
}

func Start(addr string)	error {
	return router.Run(addr)
}