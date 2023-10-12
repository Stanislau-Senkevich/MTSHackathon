package handler

import (
	"MTSHackathonBackEnd/internal/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Input struct {
	CertId      string `json:"cert_id"`
	UserId      string `json:"user_id"`
	PhoneNumber string `json:"phone_number" bson:"phone_number"`
}

// @Summary	Get all bought certificates
// @Tags		certificate
// @ID			get-bought-certificate
// @Produce	json
// @Param		input	body		Input	true	"input info"
// @Success	200		{object}	[]entity.Certificate
// @Router		/certificate [get]
func (h *Handler) getAllBoughtCertificates(c *gin.Context) {
	var input Input
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	certs, err := h.services.Certificate.GetAllBoughtCertificates(input.UserId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, certs)
}

// @Summary	Accept link
// @Tags		certificate
// @ID			accept-link
// @Produce	json
// @Param		input	body		Input	true	"input info"
// @Success	200		{object}	Input
// @Router		/certificate/link [put]
func (h *Handler) acceptLink(c *gin.Context) {
	var input Input
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	err = h.services.ChangeOwnerOfCertificate(input.CertId, input.UserId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Bad request")
		return
	}
	c.JSON(http.StatusOK, input)
}

// @Summary	Generate link
// @Tags		certificate
// @ID			generate-link
// @Produce	json
// @Param		input	body		entity.Certificate	true	"certificate info"
// @Success	200		{object}	map[string]string
// @Router		/certificate/link [get]
func (h *Handler) generateLink(c *gin.Context) {
	var cert entity.Certificate
	err := c.ShouldBindJSON(&cert)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Invalid id")
	}
	link := h.services.GenerateLink(cert.Id)
	c.JSON(http.StatusOK, map[string]string{
		"link": link,
	})
}

// @Summary	Create new certificate
// @Tags		certificate
// @ID			create-new-certificate
// @Produce	json
// @Param		input	body		entity.Certificate	true	"certificate info"
// @Success	200		{object}	map[string]string
// @Router		/certificate [post]
func (h *Handler) createNewCertificate(c *gin.Context) {
	var certificate entity.Certificate
	err := c.ShouldBindJSON(&certificate)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.GenerateUniqueId()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Internal error due generating id")
		return
	}
	certificate.Id = id

	err = h.services.CreateNewCertificate(certificate)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]string{
		"cert_id": certificate.Id,
	})
}

// @Summary	Share phone number
// @Tags		certificate
// @ID			share-number
// @Produce	json
// @Param		input	body		Input	true	"input info"
// @Success	200		{object}	Input
// @Router		/certificate/phone-number [get]
func (h *Handler) shareByPhoneNumber(c *gin.Context) {
	var input Input
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.GetUserIdByPhoneNumber(input.PhoneNumber)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}
	err = h.services.ChangeOwnerOfCertificate(input.CertId, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, input)
}
