package handler

import (
	pb "api_getaway/genproto/orders"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create Review
// @Description Create Review
// @Tags Reviews
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param review body orders.CreateReviewRequest true "Create Review"
// @Success 200 {object} orders.Review
// @Failure 400 {object} models.Error
// @Router /api/v1/reviews [post]
func (h *Handler) CreateReview(c *gin.Context) {
	req := pb.CreateReviewRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("CreateReview request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	res, err := h.OrderClient.CreateReview(c, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("CreateReview request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Get Review
// @Description Get Review
// @Tags Reviews
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param KitchenId path string true "Get Review about Kitchen"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} orders.GetReviewsResponse
// @Failure 400 {object} models.Error
// @Router /api/v1/reviews/{KitchenId} [get]
func (h *Handler) GetKitchenReviewsById(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Requested ID:", id)
	limitstr := c.Query("limit")
	offsetstr := c.Query("offset")
	if limitstr == "" {
		limitstr = "10"
	}

	if offsetstr == "" {
		offsetstr = "0"
	}

	limit, err := strconv.Atoi(limitstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "limit must be integer",
		})
		return
	}
	offset, err := strconv.Atoi(offsetstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "offset must be integer",
		})
		return
	}
	
	req := pb.GetReviewsRequest{
		KitchenId: id,
		Limit:     int32(limit),
		Offset:    int32(offset),
	}

	res, err := h.OrderClient.GetReviews(c, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("GetReview request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Update Review
// @Description Update Review
// @Tags Reviews
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param review body orders.UpdateReviewRequest true "Update Review"
// @Success 200 {object} orders.Status
// @Failure 400 {object} models.Error
// @Router /api/v1/reviews [put]
func (h *Handler) UpdateReview(c *gin.Context) {
	req := pb.UpdateReviewRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("UpdateReview request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	res, err := h.OrderClient.UpdateReview(c, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("UpdateReview request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !res.Status {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "UpdateReview request error",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Review Updated successfully",
		})
	}
}

// @Summary Delete Review
// @Description Delete Review
// @Tags Reviews
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Delete Review"
// @Success 200 {object} orders.Status
// @Failure 400 {object} models.Error
// @Router /api/v1/reviews/{id} [delete]
func (h *Handler) DeleteReview(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Requested ID:", id)
	req := pb.Review{
		Id: id,
	}
	res, err := h.OrderClient.DeleteReview(c, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("DeleteReview request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if !res.Status {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "DeleteReview request error",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Review Deleted successfully",
		})
	}
}
