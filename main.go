package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/supperdoggy/helper/config"
	"github.com/supperdoggy/helper/pkg/clients/email"
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

	emailClient := email.NewClient(*logger, cfg.EmailCheckService)
	services := service.NewService(logger, mongo, emailClient)
	handlers := handler.NewHandler(logger, services)

	r := gin.Default()

	// endpoints

	api := r.Group("/api/v1")

	// auth
	auth := api.Group("/auth")
	auth.POST("/register", handlers.Register)
	auth.POST("/login", handlers.Login)
	auth.POST("/check_token", handlers.CheckToken)
	auth.POST("/new_email_code", handlers.NewEmailCode)
	auth.POST("/check_email_code", handlers.CheckEmailCode)

	// users
	apiUser := api.Group("/user")
	apiUser.Use(handlers.Middleware)
	{
		apiUser.POST("/create", handlers.CreateUser)
		apiUser.DELETE("/delete", handlers.DeleteUser)
		apiUser.PATCH("/update", handlers.UpdateUser)
		apiUser.GET("/get/:id", handlers.GetUser)
	}

	// adverts
	adverts := api.Group("/adverts")
	adverts.Use(handlers.Middleware)
	{
		adverts.POST("/create", handlers.CreateAdvert)
		adverts.DELETE("/delete", handlers.DeleteAdvert)
		adverts.POST("/get", handlers.GetAdvert)
		adverts.POST("/filter", handlers.GetAdverts)
		adverts.GET("/attachments/:id", handlers.GetAdvertAttachments)
	}

	if err := r.Run(fmt.Sprintf("localhost:%d", cfg.Port)); err != nil {
		logger.Fatal("error r.Run", zap.Error(err))
	}
}
