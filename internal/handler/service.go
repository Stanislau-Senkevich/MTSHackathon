package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetAllServices(c *gin.Context) {
	services, err := h.services.GetAllServices()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Internal error")
		return
	}
	c.JSON(http.StatusOK, services)
}
