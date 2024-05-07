package server

import (
	"log"
	"scalexmedia-assignment/pkg/api/handlers"
	"scalexmedia-assignment/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(adminHandler *handlers.AdminHandler, userHandler *handlers.UserHandler) *ServerHTTP {
	r := gin.New()
	r.Use(gin.Logger())

	adminGroup := r.Group("/admin")
	adminGroup.POST("/login", adminHandler.AdminLogin)
	adminGroup.Use(middleware.AdminAuthMiddleware())
	{
		adminGroup.GET("/home", adminHandler.ShowBooks)
		adminGroup.POST("/addBook", adminHandler.AddBook)
		adminGroup.DELETE("/deleteBook", adminHandler.DeleteBook)
	}

	userGroup := r.Group("/user")
	userGroup.POST("/login", userHandler.UserLogin)
	userGroup.Use(middleware.UserAuthMiddleware())
	{
		userGroup.GET("/home", adminHandler.ShowBooks)
	}

	return &ServerHTTP{engine: r}
}

func (s *ServerHTTP) Start() {
	log.Printf("Starting Server on 3000")
	err := s.engine.Run(":3000")
	if err != nil {
		log.Printf("error while starting the server")
	}
}
