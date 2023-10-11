package handler

import (
	"MTSHackathonBackEnd/internal/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}

func (h *Handler) InitServiceRoutes(router *gin.Engine) {
	serviceGroup := router.Group("/service")
	{
		serviceGroup.GET("/", h.GetAllServices)
	}
}

func (h *Handler) InitCategoryRoutes(router *gin.Engine) {
	categoryGroup := router.Group("/category")
	{
		categoryGroup.GET("/", h.GetAllCategories)
	}
}

func (h *Handler) InitCertificateRoutes(router *gin.Engine) {
	certGroup := router.Group("/certificate")
	{
		certGroup.GET("/", h.GetAllBoughtCertificates)
		//certGroup.PUT("/", h.ChangeOwnerOfCertificate)
		certGroup.GET("/link", h.GenerateLink)
		certGroup.POST("/", h.CreateNewCertificate)
		certGroup.PUT("/phone-number", h.ShareByPhoneNumber)
	}
}
