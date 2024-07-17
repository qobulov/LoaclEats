package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	pb "api_getaway/genproto/orders"
)

// @Summary Get Kitchen Statistics
// @Description Get Kitchen Statistics
// @Tags Extra
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Kitchen ID"
// @Param start_date query string true "Start Date"
// @Param end_date query string true "End Date"
// @Success 200 {object} orders.GetKitchenStatisticsResponse
// @Failure 400 {object} models.Error
// @Router /api/v1/extra/kitchens/{id}/statistics [get]
func (h *Handler) GetKitchenStatistics(c *gin.Context) {
	req := &pb.GetKitchenStatisticsRequest{
		KitchenId: c.Param("id"),
		StartDate: c.Query("start_date"),
		EndDate:   c.Query("end_date"),
	}

	if req.KitchenId == "" || req.StartDate == "" || req.EndDate == "" {
		h.Logger.Error("Missing parameters")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing parameters",
		})
		return
	}

	res, err := h.OrderClient.GetKitchenStatistics(context.Background(), req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("GetKitchenStatistics request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Get User Activity
// @Description Get User Activity
// @Tags Extra
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "User ID"
// @Param start_date query string true "Start Date"
// @Param end_date query string true "End Date"
// @Success 200 {object} orders.GetUserActivityResponse
// @Failure 400 {object} models.Error
// @Router /api/v1/extra/users/{id}/activity [get]
func (h *Handler) GetUserActivity(c *gin.Context) {
	req := &pb.GetUserActivityRequest{
		UserId:    c.Param("id"),
		StartDate: c.Query("start_date"),
		EndDate:   c.Query("end_date"),
	}

	if req.UserId == "" || req.StartDate == "" || req.EndDate == "" {
		h.Logger.Error("Missing parameters")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing parameters",
		})
		return
	}

	res, err := h.OrderClient.GetUserActivity(context.Background(), req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("GetUserActivity request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Update Working Hours
// @Description Update Working Hours
// @Tags Extra
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param working_hours body orders.UpdateWorkingHoursRequest true "Update Working Hours"
// @Success 200 {object} orders.UpdateWorkingHoursResponse
// @Failure 400 {object} models.Error
// @Router /api/v1/extra/kitchens/update/working-hours [put]
func (h *Handler) UpdateWorkingHours(c *gin.Context) {
	req := &pb.UpdateWorkingHoursRequest{}

	if err := c.ShouldBindJSON(req); err != nil {
		h.Logger.Error(fmt.Sprintf("UpdateWorkingHours request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}


	res, err := h.OrderClient.UpdateWorkingHours(context.Background(), req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("UpdateWorkingHours request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary create Kitchen Working Hours
// @Description create Kitchen Working Hours
// @Tags Extra
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param working_hours body orders.CreateWorkingHoursRequest true "Create Working Hours"
// @Success 200 {object} orders.CreateWorkingHoursResponse
// @Failure 400 {object} models.Error
// @Router /api/v1/extra/kitchens/working-hours [post]
func (h *Handler) CreateKitchenWorkingHours(c *gin.Context) {
	req := &pb.CreateWorkingHoursRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		h.Logger.Error(fmt.Sprintf("CreateWorkingHours request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := h.OrderClient.CreaterWorkingHours(context.Background(), req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("CreateWorkingHours request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}