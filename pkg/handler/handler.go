package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/supperdoggy/helper/pkg/models"
	"github.com/supperdoggy/helper/pkg/service"
	"github.com/supperdoggy/helper/pkg/utils"
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

func (h *handler) CreateUser(c *gin.Context) {
	var (
		req  models.CreateUserRequest
		resp models.CreateUserResponse
		ctx  context.Context
	)
	if err := c.Bind(&req); err != nil {
		h.logger.Error("error Bing", zap.Error(err))
		resp.Error = "error reading request"
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	user, err := h.service.CreateUser(ctx, req.Password, req.Email, req.FullName)
	if err != nil {
		h.logger.Error("error CreateUser", zap.Error(err))
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp.ID = user.ID
	c.JSON(http.StatusOK, resp)
}

func (h *handler) DeleteUser(c *gin.Context) {
	var (
		req  models.DeleteUserRequest
		resp models.DeleteUserResponse
		ctx  context.Context
		err  error
	)
	if err := c.Bind(&req); err != nil {
		h.logger.Error("error Bing", zap.Error(err))
		resp.Error = "error reading request"
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	h.logger.Info("DeleteUser", zap.Any("req", req))

	id, err := h.service.DeleteUser(ctx, req.ID)
	if err != nil {
		h.logger.Error("error deleting user", zap.Error(err), zap.Any("req", req))
		resp.Error = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp.ID = id
	c.JSON(http.StatusOK, resp)
}

func (h *handler) UpdateUser(c *gin.Context) {
	var (
		req  models.UpdateUserRequest
		resp models.UpdateUserResponse
		ctx  context.Context
		err  error
	)
	if err := c.Bind(&req); err != nil {
		h.logger.Error("error Bing", zap.Error(err))
		resp.Error = "error reading request"
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	user, err := h.service.UpdateUser(ctx, req.ID, req.Password, req.Email)
	if err != nil {
		h.logger.Error("error UpdateUser", zap.Error(err), zap.Any("id", req.ID))
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp.User = utils.MapDBUserToResponseUser(*user)
	c.JSON(http.StatusBadRequest, resp)
}

func (h *handler) Login(c *gin.Context) {
	var (
		req  models.LoginReq
		resp models.LoginResp
		err  error
	)

	if err := c.Bind(&req); err != nil {
		h.logger.Error("error Bing", zap.Error(err))
		resp.Error = "error reading request"
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp.UserID, resp.Token, err = h.service.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		h.logger.Info("error Login", zap.Error(err))
		resp.Error = "Login error"
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *handler) CheckToken(c *gin.Context) {
	var (
		req  models.CheckTokenReq
		resp models.CheckTokenResp
		err  error
	)

	if err := c.Bind(&req); err != nil {
		h.logger.Error("error Bing", zap.Error(err))
		resp.Error = "error reading request"
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp.UserID, err = h.service.CheckToken(c.Request.Context(), req.Token)
	if err != nil {
		h.logger.Info("error Login", zap.Error(err))
		resp.Error = "CheckToken error"
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	resp.OK = true

	c.JSON(http.StatusOK, resp)
}

func (h *handler) Register(c *gin.Context) {
	var (
		req  models.RegisterReq
		resp models.RegisterResp
		err  error
	)

	if err := c.Bind(&req); err != nil {
		h.logger.Error("error Bing", zap.Error(err))
		resp.Error = "error reading request"
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	fullname := fmt.Sprintf("%s %s", req.FirstName, req.LastName)

	resp.UserID, resp.Token, err = h.service.Register(c.Request.Context(), req.Email, fullname, req.Password)
	if err != nil {
		h.logger.Error("error Bing", zap.Error(err))
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
}
