package routes

import (
	"go-blogrpl/controller"
	"go-blogrpl/middleware"
	"go-blogrpl/service"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, userC controller.UserController) {
	userAuthlessRoutes := router.Group("/users")
	{
		userAuthlessRoutes.GET("/", middleware.Authenticate(service.NewJWTService(), "admin"), userC.GetAllUsers)
		// userRoutes.GET("/:id", userController.Get)
		userAuthlessRoutes.POST("/signup", userC.SignUp)
		// userRoutes.PUT("/:id", userController.Update)
		// userRoutes.DELETE("/:id", userController.Delete)
	}
}
