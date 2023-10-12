package handler

import (
	"MTSHackathonBackEnd/internal/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

//	@Summary	Get services
//	@Tags		service
//	@ID			get-services
//	@Produce	json
//	@Success	200	{object}	[]entity.Service
//	@Router		/service [get]
func (h *Handler) getAllServices(c *gin.Context) {
	services, err := h.services.Service.GetAllServices()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Internal error")
		return
	}
	c.JSON(http.StatusOK, map[string][]entity.Service{
		"list": services,
	})
}
