package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/supperdoggy/helper/config"
	"github.com/supperdoggy/helper/pkg/handler"
	"github.com/supperdoggy/helper/pkg/service"
	"github.com/supperdoggy/helper/pkg/storage"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	cfg, err := config.NewConfig()
	if err != nil {
		logger.Fatal("error getting config", zap.Error(err))
	}

	mongo, err := storage.NewMongoClient(ctx, cfg.MongoUrl, logger)
	if err != nil {
		logger.Fatal("error connecting to mongo", zap.Error(err))
	}

	services := service.NewService(logger, mongo)
	handlers := handler.NewHandler(logger, services)

	r := gin.Default()

	// endpoints

	api := r.Group("/api/v1")

	apiUser := api.Group("/user")
	apiUser.POST("/create", handlers.CreateUser)
	apiUser.DELETE("/delete", handlers.DeleteUser)
	apiUser.PATCH("/update", handlers.UpdateUser)

	auth := api.Group("/auth")

	auth.POST("/register", handlers.Register)
	auth.POST("/login", handlers.Login)
	auth.POST("/check_token", handlers.CheckToken)
	// auth.POST("/")

	if err := r.Run(fmt.Sprintf("localhost:%d", cfg.Port)); err != nil {
		logger.Fatal("error r.Run", zap.Error(err))
	}
}
