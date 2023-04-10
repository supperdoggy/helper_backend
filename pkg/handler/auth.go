package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/supperdoggy/helper/pkg/models"
	"go.uber.org/zap"
)

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
