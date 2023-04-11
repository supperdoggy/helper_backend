package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/supperdoggy/helper/pkg/models"
	"github.com/supperdoggy/helper/pkg/utils"
	"go.uber.org/zap"
)

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

func (h *handler) GetUser(c *gin.Context) {
	var (
		resp models.GetUserResponse
		ctx  context.Context
		err  error
	)
	userID := c.Param("id")

	user, err := h.service.GetUser(ctx, userID)
	if err != nil {
		h.logger.Error("error GetUser", zap.Error(err), zap.Any("id", userID))
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp.User = utils.MapDBUserToResponseUser(*user)
	c.JSON(http.StatusOK, resp)
}
