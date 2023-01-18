package handler

import (
	"net/http"
	"tokokecilkita-go/helper"
	"tokokecilkita-go/organization"

	"github.com/gin-gonic/gin"
)

type organizationHandler struct {
	service organization.Service
}

func NewOrganizationHandler(service organization.Service) *organizationHandler {
	return &organizationHandler{service}
}

func (h *organizationHandler) Save(c *gin.Context) {
	var input organization.OrganizationCreateInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Failed to create task", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	organizations, err := h.service.Save(input)
	if err != nil {
		response := helper.APIResponse("Failed to save organization", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success save organization !", http.StatusOK, "success", organization.FormatOrganization(organizations))
	c.JSON(http.StatusOK, response)
}

func (h *organizationHandler) FindById(c *gin.Context) {
	var input organization.FindById
	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Failed to create task", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	organizationSingle, err := h.service.FindById(input.ID)
	if err != nil {
		response := helper.APIResponse("Failed to get organization", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success get organization !", http.StatusOK, "success", organization.FormatOrganization(organizationSingle))
	c.JSON(http.StatusOK, response)
}
