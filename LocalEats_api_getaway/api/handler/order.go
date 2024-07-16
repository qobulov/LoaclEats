package handler

import (
	pb "api_getaway/genproto/orders"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create Order
// @Description Create Order
// @Tags Orders
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param order body orders.CreateOrderRequest true "Create Order"
// @Success 200 {object} orders.CreateOrderResponse
// @Failure 400 {object} models.Error
// @Router /api/v1/orders [post]
func (h *Handler) CreateOrder(c *gin.Context) {
	req := pb.CreateOrderRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("CreateOrder request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if h.OrderClient == nil {
		h.Logger.Error("OrderClient is not initialized")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "OrderClient is not initialized",
		})
		return
	}

	res, err := h.OrderClient.CreateOrder(c.Request.Context(), &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("CreateOrder request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Println("Response Object:", res)
	c.JSON(http.StatusOK, res)
}


// @Summary Get User Orders
// @Description Get User Orders
// @Tags Orders
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Get User Order"
// @Success 200 {object} orders.GetUserOrdersResponse
// @Failure 400 {object} models.Error
// @Router /api/v1/orders/{id} [get]
func (h *Handler) ListUserOrders(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Requested ID:", id)

	req := pb.GetUserOrdersRequest{
		UserId: id,
	}

	res, err := h.OrderClient.GetUserOrders(c, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("GetOrder request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

// Summary Update Order Status
// @Description Update Order Status
// @Tags Orders
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param order body orders.UpdateOrderStatusRequest true "Update Order Status"
// @Success 200 {object} orders.Status
// @Failure 400 {object} models.Error
// @Router /api/v1/orders/status [put]
func (h *Handler) UpdateOrderStatus(c *gin.Context) {
	req := pb.UpdateOrderStatusRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("UpdateOrderStatus request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := h.OrderClient.UpdateOrderStatus(c, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("UpdateOrderStatus request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if !res.Status{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to update order status",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message":"Order status updated successfully"})
	}

}

// @Summary Get Order By Id
// @Description Get Order By Id
// @Tags Orders
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Get Order By Id"
// @Success 200 {object} orders.GetOrderByIDResponse
// @Failure 400 {object} models.Error
// @Router /api/v1/orders/{id} [get]
func (h *Handler) GetOrderByID(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Requested ID:", id)

	req := pb.GetOrderByIDRequest{
		Id: id,
	}

	res, err := h.OrderClient.GetOrderByID(c, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("GetOrderById request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary List Kitchen Orders
// @Description List Kitchen Orders
// @Tags Orders
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "List Kitchen Orders"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} orders.GetKitchenOrdersResponse
// @Failure 400 {object} models.Error
// @Router /api/v1/orders/kitchen/{id} [get]
func (h *Handler) ListKitchenOrders(c *gin.Context) {

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		limit = 10
	}
	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		offset = 0
	}

	id := c.Param("id")
	fmt.Println("Requested ID:", id)

	req := pb.GetKitchenOrdersRequest{
		KitchenId: id,
		Limit:     int32(limit),
		Offset:    int32(offset),
	}

	res, err := h.OrderClient.GetKitchenOrders(c, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("ListKitchenOrders request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}
