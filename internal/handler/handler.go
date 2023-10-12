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

	h.InitServiceRoutes(router)
	h.InitCategoryRoutes(router)
	h.InitCertificateRoutes(router)
	h.InitUserRoutes(router)
	return router
}

func (h *Handler) InitServiceRoutes(router *gin.Engine) {
	serviceGroup := router.Group("/service")
	{
		serviceGroup.GET("/", h.getAllServices)
	}
}

func (h *Handler) InitCategoryRoutes(router *gin.Engine) {
	categoryGroup := router.Group("/category")
	{
		categoryGroup.GET("/", h.getAllCategories)
	}
}

func (h *Handler) InitCertificateRoutes(router *gin.Engine) {
	certGroup := router.Group("/certificate")
	{
		certGroup.GET("/", h.getAllBoughtCertificates)
		certGroup.PUT("/link", h.acceptLink)
		certGroup.GET("/link", h.generateLink)
		certGroup.POST("/", h.createNewCertificate)
		certGroup.PUT("/phone-number", h.shareByPhoneNumber)
	}
}

func (h *Handler) InitUserRoutes(router *gin.Engine) {

}
