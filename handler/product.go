package handler

import (
	"net/http"
	"strconv"
	"tokokecilkita-go/helper"
	"tokokecilkita-go/product"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	service product.Service
}

func NewProductHandler(service product.Service) *productHandler {
	return &productHandler{service}
}

func (h *productHandler) Save(c *gin.Context) {
	var input product.ProductCreateInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Failed to create product", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	products, err := h.service.Save(input)
	if err != nil {
		response := helper.APIResponse("Failed to save product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success save product !", http.StatusOK, "success", product.FormatProductV1(products))
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) SaveImage(c *gin.Context) {
	productId, _ := strconv.Atoi(c.PostForm("productId"))
	file, err := c.FormFile("image")
	if err != nil {
		response := helper.APIResponse("Failed to save image", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	path := "images/" + file.Filename
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		response := helper.APIResponse("Failed to save image", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	productsUpdated, err := h.service.SaveImage(productId, path)
	if err != nil {
		response := helper.APIResponse("Failed to save product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success save product !", http.StatusOK, "success", product.FormatProductV1(productsUpdated))
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) UpdateProduct(c *gin.Context) {
	var inputID product.FindById
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData product.ProductCreateInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Failed to update product", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	updatedProduct, err := h.service.UpdateProduct(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success create product", http.StatusOK, "success", product.FormatProductV1(updatedProduct))
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) FindById(c *gin.Context) {
	var input product.FindById
	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Failed to create product", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	productSingle, err := h.service.FindById(input.ID)
	if err != nil {
		response := helper.APIResponse("Failed to get product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success get product !", http.StatusOK, "success", product.FormatProductV1(productSingle))
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) FindAll(c *gin.Context) {
	q := c.Request.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	pageSize, _ := strconv.Atoi(q.Get("page_size"))
	s := q.Get("q")
	products, err := h.service.FindAll(page, pageSize, s)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Failed to get product data", http.StatusOK, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of products", http.StatusOK, "success", product.FormatProducts(products))
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) Delete(c *gin.Context) {
	var input product.FindById
	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Failed to delete product", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	productSingle, err := h.service.Delete(input.ID)
	if err != nil {
		response := helper.APIResponse("Failed to delete product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success delete product !", http.StatusOK, "success", product.FormatProductV1(productSingle))
	c.JSON(http.StatusOK, response)
}
