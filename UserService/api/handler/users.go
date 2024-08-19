package handler

import (
	"AuthService/api/token"
	pb "AuthService/genproto/users"
	"AuthService/models"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// @Summary Register User
// @Description to register user in the SitAndEat app
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body users.RegisterRequest true "user"
// @Success 200 {object} users.Status
// @Failure 400 {object} users.Status
// @Router /api/v1/auth/register [post]
func (h *Handler) Register(ctx *gin.Context) {
	req := pb.RegisterRequest{}
	err := json.NewDecoder(ctx.Request.Body).Decode(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		h.logger.Error(err.Error())
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		h.logger.Error(err.Error())
		return
	}
	req.Password = string(hashedPassword)

	status, err := h.UserRepo.Register(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		h.logger.Error(err.Error())
		return
	}

	if !status.Status {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "Failed to register user")
		h.logger.Error(errors.New("failed to register user").Error())
		return
	}

	
	ctx.JSON(http.StatusAccepted, gin.H{"message": "User registered successfully"})
}

// @Summary Login User
// @Description to login user in the SitAndEat app
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body users.LoginRequest true "user"
// @Success 200 {object} users.Status
// @Failure 400 {object} models.Error
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
		FullName:    user.FullName,
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
// @Param user body users.LogoutRequest true "user"
// @Success 200 {object} users.Status
// @Failure 400 {object} users.Status
// @Router /api/v1/auth/logout [post]
func (h *Handler) Logout(ctx *gin.Context) {
	req := pb.LogoutRequest{}
	if err := json.NewDecoder(ctx.Request.Body).Decode(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		h.logger.Error(err.Error())
		return
	}
	msg, err := h.UserRepo.DeleteRefreshToken(req.Email)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, msg)
		h.logger.Error(err.Error())
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{"message": msg})
}

// @Summary Reset Password
// @Description to reset password in the SitAndEat app
// @Tags Auth
// @Accept json
// @Produce json
// @Param email path string true "email"
// @Param code path string true "enter code from email"
// @Param password path string true "new password"
// @Success 200 {object} users.Status
// @Failure 400 {object} users.Status
// @Router /api/v1/auth/reset-password/{email}/{code}/{password} [post]
func (h *Handler) ResetPassword(ctx *gin.Context) {
	email := ctx.Param("email")
	cod := ctx.Param("code")
	newPassword := ctx.Param("password")

	if email == "" || cod == "" || newPassword == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "Email, code, and new password are required")
		h.logger.Error(errors.New("email, code, and new password are required").Error())
		return
	}
	code , _ := strconv.Atoi(cod)
	req := pb.ResetPasswordRequest{
		Email:       email,
		NewPassword: newPassword,
		Verification: int32(code),
	}

	status, err := h.UserRepo.ResetPassword(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "Failed to reset password")
		h.logger.Error(err.Error())
		return
	}

	if !status.Status {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message":"Invalid Password!"})
		h.logger.Error(errors.New("invalid password").Error())
		return
	}

	ctx.JSON(http.StatusOK, "Password reset successfully")
}

// @Summary Refresh Token
// @Description to refresh token in the LocalEats app
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body users.LoginRequest true "user"
// @Success 200 {object} users.Token
// @Failure 400 {object} users.Status
// @Router /api/v1/auth/refresh-token [post]
func (h *Handler) RefreshToken(ctx *gin.Context) {
	req := pb.LoginRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		h.logger.Error(err.Error())
		return
	}
	reftoken, err := h.UserRepo.RefreshToken(req.Email)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		h.logger.Error(err.Error())
		return
	}

	token := token.RefreshToken(reftoken)

	ctx.JSON(http.StatusAccepted, gin.H{"access_token": token.RefreshToken, "expires_at": time.Now().Add(time.Hour * 24), "message": "Token refreshed successfully"})
}

// @Summary Forgot Password
// @Description to forgot password in the SitAndEat app
// @Tags Auth
// @Accept json
// @Produce json
// @Param email path string true "email"
// @Success 200 {object} users.Status
// @Failure 400 {object} users.Status
// @Router /api/v1/auth/forgot-password/{email} [post]
func (h *Handler) ForgotPassword(ctx *gin.Context) {
	email := ctx.Param("email")
	req := &pb.ResetPasswordRequest{Email: email}
	if email == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "Email is required")
		h.logger.Error(errors.New("email is required").Error())
		return
	}
	status, err := h.UserRepo.ForgotPassword(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "Your Email is not registered")
		h.logger.Error(err.Error())	
		return
	}
	if !status.Status {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "Failed to send reset password email 1")
		h.logger.Error(errors.New("failed to send reset password email").Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Password reset email sent successfully"})
}