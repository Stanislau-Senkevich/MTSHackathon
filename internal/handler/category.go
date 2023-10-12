package handler

import (
	"MTSHackathonBackEnd/internal/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary	Get categories
// @Tags		category
// @ID			get-categories
// @Produce	json
// @Success	200	{object} []entity.Category
// @Router		/category [get]
func (h *Handler) getAllCategories(c *gin.Context) {
	categories, err := h.services.GetAllCategories()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Internal error")
		return
	}
	c.JSON(http.StatusOK, map[string][]entity.Category{
		"list": categories,
	})
}
