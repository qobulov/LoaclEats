package handler

import (
	pb "api_getaway/genproto/orders"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary CreatePayment
// @Description CreatePayment
// @Tags payment
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param payment body orders.CreatePaymentRequest true "Payment"
// @Success 200 {object} orders.Payment
// @Router /api/v1/payments [post]
func (h *Handler) CreatePayment(c *gin.Context) {
	var req pb.CreatePaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Logger.Error(fmt.Sprintf("CreatePayment request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if len(req.CardNumber) != 16 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid card number",
		})
		return
	}
	res, err := h.OrderClient.CreatePayment(c, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("CreatePayment request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary GetPayment
// @Description GetPayment
// @Tags payment
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param orderID path string true "Payment"
// @Success 200 {object} orders.Payment
// @Router /api/v1/payments/{orderID} [get]
func (h *Handler) GetPaymentById(c *gin.Context) {
	id := c.Param("id")
	req := pb.GetPaymentsRequest{
		OrderId: id,
	}
	res, err := h.OrderClient.GetPayments(c, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("GetPayment request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

