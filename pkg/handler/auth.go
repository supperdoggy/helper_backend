package handler

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
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

func (h *handler) Middleware(c *gin.Context) {
	// read from the header the token
	token := c.GetHeader("Authorization")
	if token == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// check the token
	userID, err := h.service.CheckToken(c.Request.Context(), token)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// set the userID to the context
	c.Set("userID", userID)
	c.Next()
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

	if err := validation.ValidateStruct(&req,
		validation.Field(&req.FirstName, validation.Required, validation.
			Match(regexp.MustCompile("^[a-zA-ZА-Яа-я]+$"))),
		validation.Field(&req.LastName, validation.Required, validation.
			Match(regexp.MustCompile("^[a-zA-ZА-Яа-я]+$"))),
		validation.Field(&req.Email, validation.Required, is.Email),
		validation.Field(&req.Password, validation.Required, validation.Length(6, 100)),
	); err != nil {
		h.logger.Error("error Bing", zap.Error(err))
		resp.Error = err.Error()
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
