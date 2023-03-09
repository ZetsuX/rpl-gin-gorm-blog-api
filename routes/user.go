package routes

import (
	"go-blogrpl/controller"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, userC controller.UserController) {
	userAuthlessRoutes := router.Group("/users")
	{
		userAuthlessRoutes.GET("/", userC.GetAll)
		// userRoutes.GET("/:id", userController.Get)
		userAuthlessRoutes.POST("/signup", userC.SignUp)
		// userRoutes.PUT("/:id", userController.Update)
		// userRoutes.DELETE("/:id", userController.Delete)
	}

	// userAuthenticatedRoutes := router.Group("/users", middleware.Authenticate(service.NewJWTService(), "admin"))
	// {
	// 	// userRoutes.GET("/:id", userController.Get)
	// 	userAuthenticatedRoutes.POST("/test", userC.SignUp)
	// 	// userRoutes.PUT("/:id", userController.Update)
	// 	// userRoutes.DELETE("/:id", userController.Delete)
	// }

	// userAuthorizedRoutes := router.Group("/users", middleware.Authenticate(service.NewJWTService(), "admin"))
	// {
	// 	userAuthorizedRoutes.GET("/", userC.GetAll)
	// 	// userRoutes.GET("/:id", userController.Get)
	// 	// userAuthorizedRoutes.POST("/aaa", userC.SignUp)
	// 	// userRoutes.PUT("/:id", userController.Update)
	// 	// userRoutes.DELETE("/:id", userController.Delete)
	// }
}
