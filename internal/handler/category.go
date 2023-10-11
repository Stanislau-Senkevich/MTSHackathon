package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetAllCategories(c *gin.Context) {
	categories, err := h.services.GetAllCategories()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Internal error")
		return
	}
	c.JSON(http.StatusOK, categories)
}
