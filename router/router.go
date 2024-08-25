package router

import (
	"chats-api/internal/user"
	"chats-api/internal/ws"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func InitRouter(userHandler *user.Handler, wsHandler *ws.Handler) {
	router = gin.Default()

	router.POST("/api/users/signup", userHandler.CreateUser)
	router.POST("/api/users/signin", userHandler.Login)
	router.GET("/api/users/logout", userHandler.Logout)

	router.POST("/ws/create-room", wsHandler.CreateRoom)
	router.GET("/ws/join-room/:roomId", wsHandler.JoinRoom)
	router.GET("/ws/get-rooms", wsHandler.GetRooms)
	router.GET("/ws/get-clients/:roomId", wsHandler.GetClients)
}

func Start(addr string)	error {
	return router.Run(addr)
}