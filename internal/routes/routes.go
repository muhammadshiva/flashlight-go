package routes

import (
	"flashlight-go/internal/handler"
	"flashlight-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

type Router struct {
	userHandler      *handler.UserHandler
	workOrderHandler *handler.WorkOrderHandler
}

func NewRouter(
	userHandler *handler.UserHandler,
	workOrderHandler *handler.WorkOrderHandler,
) *Router {
	return &Router{
		userHandler:      userHandler,
		workOrderHandler: workOrderHandler,
	}
}

func (r *Router) Setup() *gin.Engine {
	router := gin.Default()

	// Apply middleware
	router.Use(middleware.CORSMiddleware())

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// API v1
	v1 := router.Group("/api/v1")
	{
		// Public routes
		auth := v1.Group("/auth")
		{
			auth.POST("/login", r.userHandler.Login)
			auth.POST("/register", r.userHandler.Create)
		}

		// Protected routes
		protected := v1.Group("")
		protected.Use(middleware.AuthMiddleware())
		{
			// Users
			users := protected.Group("/users")
			{
				users.GET("", r.userHandler.GetAll)
				users.GET("/:id", r.userHandler.GetByID)
				users.PUT("/:id", r.userHandler.Update)
				users.DELETE("/:id", r.userHandler.Delete)
			}

			// Work Orders
			workOrders := protected.Group("/work-orders")
			{
				workOrders.POST("", r.workOrderHandler.Create)
				workOrders.GET("", r.workOrderHandler.GetAll)
				workOrders.GET("/:id", r.workOrderHandler.GetByID)
				workOrders.PUT("/:id", r.workOrderHandler.Update)
				workOrders.DELETE("/:id", r.workOrderHandler.Delete)
			}

			// Admin only routes
			admin := protected.Group("/admin")
			admin.Use(middleware.RoleMiddleware("owner", "admin"))
			{
				admin.POST("/users", r.userHandler.Create)
			}
		}
	}

	return router
}
