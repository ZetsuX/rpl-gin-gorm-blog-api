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
		userRoutes.PUT("/self/name", middleware.Authenticate(service.NewJWTService(), "user"), userC.UpdateSelfName)
		userRoutes.DELETE("/self", middleware.Authenticate(service.NewJWTService(), "user"), userC.DeleteSelfUser)
		userRoutes.POST("/signup", userC.SignUp)
		userRoutes.POST("/signin", userC.SignIn)
	}
}
