package routes

import (
	"go-blogrpl/controller"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, userC controller.UserController) {
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/", userC.GetAllUsers)
		userRoutes.GET("/:username", userC.GetUserByUsername)
		userRoutes.POST("/signup", userC.SignUp)
		// userRoutes.PUT("/:id", userController.Update)
		// userRoutes.DELETE("/:id", userController.Delete)
	}
}
