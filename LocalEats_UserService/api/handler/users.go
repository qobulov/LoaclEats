package handler

import (
	"AuthService/api/token"
	pb "AuthService/genproto/proto"
	"AuthService/models"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// @Summary Register User
// @Description to register user in the SitAndEat app
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body proto.RegisterRequest true "user"
// @Success 200 {object} proto.Status
// @Failure 400 {object} proto.Status
// @Router /api/v1/auth/register [post]
func (h *Handler) Register(ctx *gin.Context) {
	req := pb.RegisterRequest{}
	err := json.NewDecoder(ctx.Request.Body).Decode(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		h.logger.Error(err.Error())
		return
	}

	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		h.logger.Error(err.Error())
		return
	}
	req.Password = string(hashedpassword)

	status, err := h.UserRepo.Register(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		h.logger.Error(err.Error())
		return
	}

	if !status.Status {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		h.logger.Error(err.Error())
		return
	}

	ctx.JSON(http.StatusAccepted, nil)
}

// @Summary Login User
// @Description to login user in the SitAndEat app
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body proto.LoginRequest true "user"
// @Success 200 {object} proto.Status
// @Failure 400 {object} proto.Status
// @Router /api/v1/auth/login [post]
func (h *Handler) Login(ctx *gin.Context) {
	req := pb.LoginRequest{}

	if err := json.NewDecoder(ctx.Request.Body).Decode(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		h.logger.Error(err.Error())
		return
	}

	user, err := h.UserRepo.GetUserByEmail(req.Email)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		h.logger.Error(err.Error())
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		h.logger.Error(err.Error())
		return
	}

	token := token.GenerateJWT(&pb.User{
		Id:          user.Id,
		Username:    user.Username,
		Email:       req.Email,
		PhoneNumber: user.PhoneNumber,
		UserType:    user.UserType,
		Address:     user.Address,
		FullName:    user.FullName,
		IsVerified:  user.IsVerified,
		Bio:         user.Bio,
		Specialties: user.Specialties,
		YearsOfExperience: user.YearsOfExperience,
	})

	err = h.UserRepo.StoreRefreshToken(&models.RefreshToken{
		UserId:    user.Id,
		Email:     req.Email,
		Token:     token.RefreshToken,
		CreatedAt: time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		h.logger.Error(err.Error())
		return
	}

	ctx.JSON(http.StatusAccepted, token)
}

// @Summary Logout User
// @Description to logout user in the SitAndEat app
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body proto.LogoutRequest true "user"
// @Success 200 {object} proto.Status
// @Failure 400 {object} proto.Status
// @Router /api/v1/auth/logout [post]
func (h *Handler) Logout(ctx *gin.Context) {
	req := pb.LogoutRequest{}
	if err := json.NewDecoder(ctx.Request.Body).Decode(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		h.logger.Error(err.Error())
		return
	}
	msg,err := h.UserRepo.DeleteRefreshToken(req.Email)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, msg)
		h.logger.Error(err.Error())
		return
	}
	ctx.JSON(http.StatusAccepted, msg)
}

// @Summary Reset Password
// @Description to reset password in the SitAndEat app
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body proto.ResetPasswordRequest true "user"
// @Success 200 {object} proto.Status
// @Failure 400 {object} proto.Status
// @Router /api/v1/auth/reset-password [post]
func (h *Handler) ResetPassword(ctx *gin.Context) {
	req := pb.ResetPasswordRequest{}
	if err := json.NewDecoder(ctx.Request.Body).Decode(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		h.logger.Error(err.Error())
		return
	}
	status, err := h.UserRepo.ResetPassword(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		h.logger.Error(err.Error())
		return
	}

	if !status.Status {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "Password reset failed")
		h.logger.Error(err.Error())
		return
	}
	ctx.JSON(http.StatusAccepted, "Password reset successfully")
}

// @Summary Refresh Token
// @Description to refresh token in the LocalEats app
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body proto.Token true "user"
// @Success 200 {object} proto.Token
// @Failure 400 {object} proto.Status
// @Router /api/v1/auth/refresh-token [post]
func (h *Handler) RefreshToken(ctx *gin.Context) {
	req := pb.Token{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		h.logger.Error(err.Error())
		return
	}
	err := h.UserRepo.RefreshToken(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		h.logger.Error(err.Error())
		return
	}

	token := token.RefreshToken(req.RefreshToken)

	ctx.JSON(http.StatusAccepted, gin.H{"refresh_token": token.RefreshToken,"expires_at": time.Now().Add(time.Hour * 24),"message":"Token refreshed successfully"})
}