package api

import (
	"database/sql"
	"AuthService/api/handler"

	_ "AuthService/api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Auth Service
// @version 1.0
// @description This is the Auth service of LocalEats project

// @contact.name Azizbek
// @contact.url http://www.support_me_with_smile
// @contact.email azizbekqobulov05@gmail.com

// @host localhost:7777

func NewRouter(db *sql.DB) *gin.Engine {

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	h := handler.NewHandler(db)    
	auth := r.Group("/api/v1/auth")
    {
        auth.POST("/register", h.Register)
        auth.POST("/login", h.Login)
        auth.POST("/reset-password", h.ResetPassword)
        auth.POST("/logout", h.Logout)
		auth.POST("/refresh-token", h.RefreshToken)
    }

	return r
}
