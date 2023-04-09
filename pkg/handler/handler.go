package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/supperdoggy/helper/pkg/service"
	"go.uber.org/zap"
)

type IHandler interface {
	// users
	CreateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	UpdateUser(c *gin.Context)

	// auth
	Login(c *gin.Context)
	Register(c *gin.Context)
	CheckToken(c *gin.Context)

	// adverts
	CreateAdvert(c *gin.Context)
	DeleteAdvert(c *gin.Context)
	GetAdvert(c *gin.Context)
}

type handler struct {
	logger  *zap.Logger
	service service.IService
}

func NewHandler(l *zap.Logger, s service.IService) IHandler {
	return &handler{
		logger:  l,
		service: s,
	}
}
