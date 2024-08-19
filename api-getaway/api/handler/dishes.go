package handler

import (
	pb "api_getaway/genproto/orders"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create Dishes
// @Description Create Dishes
// @Tags Dishes
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param dishes body orders.CreateDishRequest true "Create Dishes"
// @Success 200 {object} orders.Dish
// @Failure 400 {object} models.Error
// @Router /api/v1/dishes [post]
func (h *Handler) CreateDishes(c *gin.Context) {
	req := pb.CreateDishRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("CreateDishes request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	res, err := h.OrderClient.CreateDish(c, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("CreateDishes request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Get Kitchen Dishes
// @Description Get Kitchen Dishes
// @Tags Dishes
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Get Kitchen Dishes"
// @Param limit query string false "Get Kitchen Dishes"
// @Param offset query string false "Get Kitchen Dishes"
// @Success 200 {object} orders.GetDishesResponse
// @Failure 400 {object} models.Error
// @Router /api/v1/dishes/kitchen/{id} [get]
func (h *Handler) ListDishesByKitchen(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Requested ID:", id)

	limit := c.DefaultQuery("limit", "10")
	offset := c.DefaultQuery("offset", "0")
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("GetDishes request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("GetDishes request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	req := pb.GetDishesRequest{
		KitchenId: id,
		Limit:     int32(limitInt),
		Offset:    int32(offsetInt),
	}
	res, err := h.OrderClient.GetDishes(c, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("GetDishes request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Update Dish
// @Description Update Dish
// @Tags Dishes
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param dish body orders.UpdateDishRequest true "Update Dish"
// @Success 200 {object} orders.Dish
// @Failure 400 {object} models.Error
// @Router /api/v1/dishes/update [put]
func (h *Handler) UpdateDish(c *gin.Context) {
	req := pb.UpdateDishRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("UpdateDish request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	res, err := h.OrderClient.UpdateDish(c, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("UpdateDish request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Delete Dish
// @Description Delete Dish
// @Tags Dishes
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Delete Dish"
// @Success 200 {object} orders.Dish
// @Failure 400 {object} models.Error
// @Router /api/v1/dishes/{id} [delete]
func (h *Handler) DeleteDish(c *gin.Context) {
	id := c.Param("id")
	req := pb.DeleteDishRequest{
		Id: id,
	}
	res, err := h.OrderClient.DeleteDish(c, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("DeleteDish request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	if !res.Status {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "DeleteDish request error",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Dish Deleted successfully"})
	}
}

// @Summary Get dish by id
// @Description Get dish by id
// @Tags Dishes
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Get dish by id"
// @Success 200 {object} orders.Dish
// @Failure 400 {object} models.Error
// @Router /api/v1/dishes/{id} [get]
func (h *Handler) GetDishById(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Requested ID:", id)

	req := pb.GetDishRequest{
		DishId: id,
	}

	res, err := h.OrderClient.GetDish(c, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("GetDishById request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}
