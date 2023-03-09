package routes

import (
	"go-blogrpl/controller"
	"go-blogrpl/middleware"
	"go-blogrpl/service"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, userC controller.UserController) {
	userRoutes := router.Group("/users")
	{
		//, middleware.Authenticate(service.NewJWTService(), "admin")
		userRoutes.GET("/", middleware.Authenticate(service.NewJWTService(), "admin"), userC.GetAllUsers)
		userRoutes.GET("/:username", middleware.Authenticate(service.NewJWTService(), "user"), userC.GetUserByUsername)
		userRoutes.POST("/signup", userC.SignUp)
		userRoutes.POST("/signin", userC.SignIn)
		// userRoutes.PUT("/:id", userController.Update)
		// userRoutes.DELETE("/:id", userController.Delete)
	}
}
