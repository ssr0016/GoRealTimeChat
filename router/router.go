package router

import (
	"server/internal/user"
	"server/internal/websocket"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler, websocketHandler *websocket.Handler) {
	r = gin.Default()

	r.POST("/signup", userHandler.CreateUser)
	r.POST("/login", userHandler.Login)
	r.GET("/logout", userHandler.Logout)

	r.POST("/websocket/createRoom", websocketHandler.CreateRoom)
	r.GET("/websocket/JoinRoom/:roomId", websocketHandler.JoinRoom)
	r.GET("/websocket/getRooms", websocketHandler.GetRooms)
	r.GET("/websocket/getClients/:roomId", websocketHandler.GetClients)
}

func Start(addr string) error {
	return r.Run(addr)
}
