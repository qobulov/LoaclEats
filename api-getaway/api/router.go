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
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @schemes http
// @BasePath /
func NewRouter() *gin.Engine {
	router := gin.Default()

	// Swagger endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	h := handler.NewHandlerRepo()

	// Middleware
	router.Use(middleware.JWTMiddleware())

	// UserService endpoints
	users := router.Group("/api/v1/users")
	{
		users.GET("/profile/:id", h.GetProfileById)
		users.PUT("/profile/update", h.UpdateProfile)
		users.DELETE("/profile/:id", h.DeleteProfile)
	}

	// Kitchen endpoints
	kitchens := router.Group("/api/v1/kitchens")
	{
		kitchens.POST("/", h.CreateKitchen)
		kitchens.PUT("/:id", h.UpdateKitchen)
		kitchens.GET("/:id", h.GetKitchenById)
		kitchens.GET("/", h.ListKitchens)
		kitchens.DELETE("/:id", h.DeleteKitchen)
		kitchens.GET("/search", h.SearchKitchens)
	}

	// Order endpoints
	orders := router.Group("/api/v1/orders")
	{
		orders.POST("/", h.CreateOrder)
		orders.GET("/:id", h.GetOrderByID)
		orders.GET("/user/:id", h.GetUserOrders)
		orders.PUT("/status", h.UpdateOrderStatus)
		orders.GET("/", h.ListUserOrders)
		orders.GET("/kitchen/:id", h.ListKitchenOrders)
	}

	// Dish endpoints
	dishes := router.Group("/api/v1/dishes")
	{
		dishes.POST("/", h.CreateDishes)
		dishes.GET("/:id", h.GetDishById)
		dishes.PUT("/update", h.UpdateDish)
		dishes.DELETE("/:id", h.DeleteDish)
		dishes.GET("/kitchen/:id", h.ListDishesByKitchen)
	}

	//Payment endpoints
	payments := router.Group("/api/v1/payments")
	{
		payments.POST("/", h.CreatePayment)
		payments.GET("/:id", h.GetPaymentById)
	}

	//Reviews endpoints
	reviews := router.Group("/api/v1/reviews")
	{
		reviews.POST("/", h.CreateReview)
		reviews.GET("/:id", h.GetKitchenReviewsById)
		reviews.PUT("/update", h.UpdateReview)
		reviews.DELETE("/:id", h.DeleteReview)
	}

	//Extra endpoints
	extra := router.Group("/api/v1/extra")
	{
		extra.GET("/kitchens/:id/statistics", h.GetKitchenStatistics)
		extra.GET("/users/:id/activity", h.GetUserActivity)
		extra.POST("/kitchens/working-hours", h.CreateKitchenWorkingHours)
		extra.PUT("/kitchens/update/working-hours", h.UpdateWorkingHours)
	}
	return router
}
