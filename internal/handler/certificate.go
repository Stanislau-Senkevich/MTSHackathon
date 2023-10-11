package handler

import (
	"MTSHackathonBackEnd/internal/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetAllBoughtCertificates(c *gin.Context) {
	id, ok := c.Get("user_id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Invalid user id")
		return
	}
	certs, err := h.services.Certificate.GetAllBoughtCertificates(id.(string))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, certs)
}

func (h *Handler) ChangeOwnerOfCertificate(c *gin.Context) {
	id, ok := c.Get("id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Invalid certificate id")
		return
	}
	newUserId, ok := c.Get("new_user_id")
	if !ok || !h.services.CheckUser(newUserId.(string)) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Invalid user id")
		return
	}
	err := h.services.ChangeOwnerOfCertificate(id.(string), newUserId.(string))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Bad request")
		return
	}
	c.JSON(http.StatusOK, map[string]string{
		"id": id.(string),
	})
}

func (h *Handler) GenerateLink(c *gin.Context) {
	id, ok := c.Get("id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Invalid certificate id")
		return
	}
	link := h.services.GenerateLink(id.(string))
	c.JSON(http.StatusOK, map[string]string{
		"link": link,
	})
}

func (h *Handler) CreateNewCertificate(c *gin.Context) {
	cert, ok := c.Get("certificate")
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Invalid certificate")
		return
	}
	certificate := cert.(entity.Certificate)
	err := h.services.CreateNewCertificate(certificate)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]string{
		"id": certificate.Id,
	})
}

func (h *Handler) ShareByPhoneNumber(c *gin.Context) {
	certId, ok := c.Get("id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Invalid id")
		return
	}
	number, ok := c.Get("phone_number")
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "No phone number")
		return
	}
	id, err := h.services.GetUserIdByPhoneNumber(number.(string))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}
	err = h.services.ChangeOwnerOfCertificate(certId.(string), id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, map[string]string{
		"cert_id": certId.(string),
	})
}
