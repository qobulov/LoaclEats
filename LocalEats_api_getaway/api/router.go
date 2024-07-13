package api

import (
	_ "api_getaway/api/docs"
	"api_getaway/api/handler"
	"api_getaway/api/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Auth Service API
// @version 1.0
// @description This is a sample server for Auth Service.
// @host localhost:4444
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @schemes http
// @BasePath /
func NewRouter() *gin.Engine {
	router := gin.Default()

	// Swagger endpointini sozlash
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	h := handler.NewHandlerRepo()

	// Middleware sozlash
	router.Use(middleware.JWTMiddleware())

	// UserService endpointlari
	users := router.Group("/api/v1/users")
	{
		users.GET("/profile/:id", h.GetProfileById)
		users.PUT("/profile/update", h.UpdateProfile)
		users.DELETE("/profile/:id", h.DeleteProfile)
	}

	// KitchenService endpointlari
	// kitchens := router.Group("/api/v1/kitchens")
	// {
	//     kitchens.POST("/", h.CreateKitchen)
	//     kitchens.PUT("/:id", h.UpdateKitchen)
	//     kitchens.GET("/:id", h.GetKitchen)
	//     kitchens.GET("/", h.ListKitchens)
	//     kitchens.GET("/search", h.SearchKitchens)
	// }

	return router
}
