package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (h *handler) GetAdvertAttachments(c *gin.Context) {
	advertID := c.Param("id")

	attachments, err := h.service.GetAdvertAttachments(c, advertID)
	if err != nil {
		h.logger.Error("failed to get advert attachments", zap.Error(err), zap.Any("id", advertID))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, attachments)
}
