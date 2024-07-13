package handler

import (
	pb "api_getaway/genproto/proto"
	"fmt"
	"net/http"

	_ "github.com/swaggo/swag"

	"github.com/gin-gonic/gin"
)


// @Summary Get Profile
// @Description Get Profile
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Get Profile"
// @Success 200 {object} proto.User
// @Failure 400 {object} models.Error
// @Router /api/v1/users/profile/{id} [get]
func (h *Handler) GetProfileById(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Requested ID:", id)

	req := pb.UserId{
		Id: id,
	}
	fmt.Println("Request Object:", req)

	res, err := h.UserClient.GetProfile(c, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("GetProfile request error: %v", err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println("Response Object:", res)
	c.JSON(http.StatusOK, res.User)
}


// @Summary update profile
// @Description update profile
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Update Profile"
// @Param user body proto.UpdateProfileRequest true "Update Profile"
// @Success 200 {object} proto.GetProfileResponse
// @Failure 400 {object} models.Error
// @Router /api/v1/users/updateProfile/{id} [put]
func (h *Handler) UpdateProfile(c *gin.Context) {
	fmt.Println("fd")
	id := c.Param("id")
	req := &pb.UpdateProfileRequest{
		UserId: id,
	}
	res, err := h.UserClient.UpdateProfile(c, req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("GetProfile request error: %v", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}
