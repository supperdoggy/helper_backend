package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/supperdoggy/helper/pkg/models"
	"github.com/supperdoggy/helper/pkg/utils"
	"go.uber.org/zap"
)

func (h *handler) CreateAdvert(c *gin.Context) {
	var (
		req  models.CreateAdvertRequest
		resp models.CreateAdvertResponse
		err  error
	)

	if err := c.Bind(&req); err != nil {
		h.logger.Error("error Bing", zap.Error(err))
		resp.Error = "error reading request"
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	if err := validation.ValidateStruct(&req, validation.Field(&req.UserID, validation.Required),
		validation.Field(&req.Body, validation.Required, validation.Length(10, 10000)),
		validation.Field(&req.Category, validation.Required), validation.Field(&req.Name, validation.Required),
		validation.Field(&req.Location, validation.Required), validation.Field(&req.Type, validation.Required)); err != nil {
		h.logger.Error("error validating request", zap.Error(err))
		resp.Error = "error validating request"
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	h.logger.Info("CreateAdvert", zap.Any("req", req))

	advert, err := h.service.CreateAdvert(c.Request.Context(), req.Name, req.Body, req.Type, req.Category, req.Location, req.UserID, req.Attachments)
	if err != nil {
		h.logger.Error("error creating advert", zap.Error(err))
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp.Advert = utils.MapDBAdvertToModelAdvert(advert)
	h.logger.Info("CreateAdvert", zap.Any("resp", resp))

	c.JSON(http.StatusOK, resp)
}

func (h *handler) DeleteAdvert(c *gin.Context) {
	var (
		req  models.DeleteAdvertRequest
		resp models.DeleteAdvertResponse
		err  error
	)

	if err := c.Bind(&req); err != nil {
		h.logger.Error("error Bing", zap.Error(err))
		resp.Error = "error reading request"
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	h.logger.Info("DeleteAdvert", zap.Any("req", req))

	err = h.service.DeleteAdvert(c.Request.Context(), req.ID)
	if err != nil {
		h.logger.Error("error deleting advert", zap.Error(err))
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	h.logger.Info("DeleteAdvert", zap.Any("resp", resp))

	c.JSON(http.StatusOK, resp)
}

func (h *handler) GetAdvert(c *gin.Context) {
	var (
		req  models.GetAdvertRequest
		resp models.GetAdvertResponse
		err  error
	)

	if err := c.Bind(&req); err != nil {
		h.logger.Error("error Bing", zap.Error(err))
		resp.Error = "error reading request"
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	h.logger.Info("GetAdvert", zap.Any("req", req))

	advert, err := h.service.GetAdvert(c.Request.Context(), req.ID)
	if err != nil {
		h.logger.Error("error getting advert", zap.Error(err))
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp.Advert = utils.MapDBAdvertToModelAdvert(advert)
	h.logger.Info("GetAdvert", zap.Any("resp", resp))

	c.JSON(http.StatusOK, resp)
}

// func (h *handler) GetAdverts(c *gin.Context) {
// 	var (
// 		req  models.GetAdvertsRequest
// 		resp models.GetAdvertsResponse
// 		err  error
// 	)

// 	if err := c.Bind(&req); err != nil {
// 		h.logger.Error("error Bing", zap.Error(err))
// 		resp.Error = "error reading request"
// 		c.JSON(http.StatusBadRequest, resp)
// 		return
// 	}

// 	h.logger.Info("GetAdverts", zap.Any("req", req))

// 	adverts, err := h.service.GetAdverts(c.Request.Context(), req.Limit, req.Offset)
// 	if err != nil {
// 		h.logger.Error("error getting adverts", zap.Error(err))
// 		resp.Error = err.Error()
// 		c.JSON(http.StatusBadRequest, resp)
// 		return
// 	}

// 	resp.Adverts = utils.MapDBAdvertsToModelAdverts(adverts, nil)
// 	h.logger.Info("GetAdverts", zap.Any("resp", resp))

// 	c.JSON(http.StatusOK, resp)
// }