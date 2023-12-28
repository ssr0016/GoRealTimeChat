package main

import (
	"log"
	"server/db"
	"server/internal/user"
	"server/internal/websocket"
	"server/router"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialize database: %s", err)
	}

	userRep := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)

	hub := websocket.NewHub()
	websocketHandler := websocket.NewHandler(hub)
	go hub.Run() //separategoroutine

	router.InitRouter(userHandler, websocketHandler)
	router.Start("0.0.0.0:8080")
}
