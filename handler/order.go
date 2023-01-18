package handler

import (
	"net/http"
	"strconv"
	"tokokecilkita-go/helper"
	"tokokecilkita-go/order"

	"github.com/gin-gonic/gin"
)

type orderHandler struct {
	service order.Service
}

func NewOrderHandler(service order.Service) *orderHandler {
	return &orderHandler{service}
}

func (h *orderHandler) Save(c *gin.Context) {
	var input order.OrderCreateInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Failed to create order", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	orders, err := h.service.Save(input)
	if err != nil {
		response := helper.APIResponse("Failed to save order", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success save order !", http.StatusOK, "success", order.FormatOrderV1(orders))
	c.JSON(http.StatusOK, response)
}

func (h *orderHandler) UpdateOrder(c *gin.Context) {
	var inputID order.FindById
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update order", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData order.OrderCreateInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Failed to update order", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	updatedOrder, err := h.service.UpdateOrder(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update order", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success create order", http.StatusOK, "success", order.FormatOrderV1(updatedOrder))
	c.JSON(http.StatusOK, response)
}

func (h *orderHandler) FindById(c *gin.Context) {
	var input order.FindById
	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Failed to create order", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	orderSingle, err := h.service.FindById(input.ID)
	if err != nil {
		response := helper.APIResponse("Failed to get order", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success get order !", http.StatusOK, "success", order.FormatOrderV1(orderSingle))
	c.JSON(http.StatusOK, response)
}

func (h *orderHandler) BasketFindByOrderId(c *gin.Context) {
	var input order.FindById
	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Failed to create order", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	orderProduct, err := h.service.BasketFindByOrderId(input.ID)
	if err != nil {
		response := helper.APIResponse("Failed to get order", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success get order !", http.StatusOK, "success", order.FormatOrderProducts(orderProduct))
	c.JSON(http.StatusOK, response)
}

func (h *orderHandler) SaveBasket(c *gin.Context) {
	var input order.OrderCreateV2Input
	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Failed to create order", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	orders, err := h.service.Save(input.Order)
	if err != nil {
		response := helper.APIResponse("Failed to save order", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	orderBasket, err := h.service.SaveBasket(input, orders.ID)
	if err != nil {
		response := helper.APIResponse("Failed to save basket", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success save order !", http.StatusOK, "success", order.FormatOrderBasketV1(orders, orderBasket))
	c.JSON(http.StatusOK, response)
}

func (h *orderHandler) FindAll(c *gin.Context) {
	q := c.Request.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	pageSize, _ := strconv.Atoi(q.Get("page_size"))
	s := q.Get("q")
	orders, err := h.service.FindAll(page, pageSize, s)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Failed to get product data", http.StatusOK, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of order", http.StatusOK, "success", order.FormatOrders(orders))
	c.JSON(http.StatusOK, response)
}

// func (h *orderHandler) UpdateBasket(c *gin.Context) {
// 	currentUser := c.MustGet("currentUser").(user.User)
// 	userId := currentUser.ID

// 	var input order.UpdateBasketInput
// 	err := c.ShouldBindJSON(&input)

// 	if err != nil {
// 		errors := helper.FormatValidationError(err)
// 		errorMessage := gin.H{"error": errors}
// 		response := helper.APIResponse("Failed to create order", http.StatusUnprocessableEntity, "error", errorMessage)
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	order, err := h.service.FindByStatus("ongoing", userId)
// 	if err != nil {
// 		response := helper.APIResponse("Failed to fetch order", http.StatusBadRequest, "error", nil)
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	orderBasket, err := h.service.SaveBasketV2(input, order.ID)
// 	if err != nil {
// 		response := helper.APIResponse("Failed to save basket", http.StatusBadRequest, "error", nil)
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	response := helper.APIResponse("Success save order !", http.StatusOK, "success", order.FormatOrderBasketV1(orders, orderBasket))
// 	c.JSON(http.StatusOK, response)
// }
