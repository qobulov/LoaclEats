package handler

import (
	pb "api_getaway/genproto/users"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/swaggo/swag"

	"github.com/gin-gonic/gin"
)

// @Summary Get Kitchen by id
// @Description Get Kitchen
// @Tags Kitchens
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Get Kitchen by id"
// @Success 200 {object} users.Kitchen
// @Failure 400 {object} models.Error
// @Router /api/v1/kitchens/{id} [get]
func (h *Handler) GetKitchenById(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Requested ID:", id)

	req := pb.GetKitchenByIDRequest{
		Id: id,
	}

	res, err := h.UserClient.GetKitchenByID(c, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("GetKitchenById request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary List Kitchens
// @Description List Kitchens
// @Tags Kitchens
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} []users.GetKitchensResponse
// @Failure 400 {object} models.Error
// @Router /api/v1/kitchens [get]
func (h *Handler) ListKitchens(c *gin.Context) {

	const defaultLimit = 10
	const defaultOffset = 0

	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limitStr == "" {
		limit = defaultLimit
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offsetStr == "" {
		offset = defaultOffset
	}

	req := pb.GetKitchensRequest{
		Limit:  int32(limit),
		Offset: int32(offset),
	}
	res, err := h.UserClient.GetKitchens(c, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("ListKitchens request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Create Kitchen
// @Description Create Kitchen
// @Tags Kitchens
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param kitchen body users.CreateKitchenRequest true "Create Kitchen"
// @Success 200 {object} users.Kitchen
// @Failure 400 {object} models.Error
// @Router /api/v1/kitchens [post]
func (h *Handler) CreateKitchen(c *gin.Context) {
	req := pb.CreateKitchenRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("CreateKitchen request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	res, err := h.UserClient.CreateKitchen(c, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("CreateKitchen request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println("Response Object:", res)
	if !res.Status {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create kitchen",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Kitchen created successfully",
		})
	}
}

// @Summary Update Kitchen
// @Description Update Kitchen
// @Tags Kitchens
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param kitchen body users.UpdateKitchenRequest true "Update Kitchen"
// @Success 200 {object} users.Kitchen
// @Failure 400 {object} models.Error
// @Router /api/v1/kitchens/{id} [put]
func (h *Handler) UpdateKitchen(c *gin.Context) {
	req := pb.UpdateKitchenRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("UpdateKitchen request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	res, err := h.UserClient.UpdateKitchen(c, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("UpdateKitchen request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if !res.Status {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to update kitchen",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Kitchen updated successfully",
		})
	}
}

// @Summary Delete Kitchen
// @Description Delete Kitchen
// @Tags Kitchens
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Delete Kitchen"
// @Success 200 {object} users.Kitchen
// @Failure 400 {object} models.Error
// @Router /api/v1/kitchens/{id} [delete]
func (h *Handler) DeleteKitchen(c *gin.Context) {
	id := c.Param("id")
	req := pb.GetKitchenByIDRequest{
		Id: id,
	}
	res, err := h.UserClient.DeleteKitchen(c, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("DeleteKitchen request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if !res.Status {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "DeleteKitchen request error",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Kitchen Deleted successfully"})
	}
}
// @Summary Search Kitchens
// @Description Search Kitchens
// @Tags Kitchens
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param name query string false "Search Kitchens"
// @Param cuisine_type query string false "Cuisine Type"
// @Param rating query number false "Rating"
// @Param offset query int false "Offset number"
// @Param limit query int false "Limit per offset"
// @Success 200 {object} users.SearchKitchensResponse
// @Failure 400 {object} models.Error
// @Router /api/v1/kitchens/search [get]
func (h *Handler) SearchKitchens(c *gin.Context) {
	name := c.Query("name")
	cuisineType := c.Query("cuisine_type")
	ratingStr := c.Query("rating")
	offsetStr := c.Query("offset")
	limitStr := c.Query("limit")

	var rating float64
	var offset, limit int
	var err error

	const defaultLimit = 10
	const defaultOffset = 0

	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit value"})
			return
		}
	} else {
		limit = defaultLimit
	}

	if offsetStr != "" {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset value"})
			return
		}
	} else {
		offset = defaultOffset
	}

	if ratingStr != "" {
		rating, err = strconv.ParseFloat(ratingStr, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid rating value"})
			return
		}
	}
	
	req := &pb.SearchKitchensRequest{
		Name:        name,
		CuisineType: cuisineType,
		Rating:      float32(rating),
		Offset:      int32(offset),
		Limit:       int32(limit),
	}

	res, err := h.UserClient.SearchKitchens(c, req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("SearchKitchens request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
