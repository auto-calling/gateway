package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"

	"github.com/auto-calling/gateway/config"
	"github.com/auto-calling/gateway/handler"
	"github.com/auto-calling/gateway/middleware"
)

func main() {

	// To initialize MongoDB's handler
	DB, err := config.ConnectionDB()
	if err != nil {
		logrus.Error("Can't connection to mongodb")
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := DB.Connect(ctx); err != nil {
		log.Error(err.Error())
	}
	defer DB.Disconnect(ctx)

	logrus.SetFormatter(&logrus.JSONFormatter{})

	router := gin.New()
	// Ignore logging with /api/ping api
	router.GET("/api/ping", handler.Ping)
	router.Use(gin.Recovery())
	router.Use(middleware.JSONLogMiddleware())
	router.Use(middleware.TokenValidate())

	router.POST("/api/v1/make/event", handler.MakeEvent)

	if err := router.Run(":80"); err != nil {
		log.Error(err.Error())
	}
}
