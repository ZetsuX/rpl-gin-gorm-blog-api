package routes

import (
	"go-blogrpl/controller"
	"go-blogrpl/middleware"
	"go-blogrpl/service"

	"github.com/gin-gonic/gin"
)

func BlogRoutes(router *gin.Engine, blogC controller.BlogController) {
	blogRoutes := router.Group("/blogs")
	{
		// //, middleware.Authenticate(service.NewJWTService(), "admin")
		blogRoutes.GET("/", blogC.GetAllBlogs)
		blogRoutes.GET("/:slug", blogC.GetBlogBySlug)
		blogRoutes.POST("/", middleware.Authenticate(service.NewJWTService(), "user"), blogC.PostBlog)
		// blogRoutes.GET("/:blogname", middleware.Authenticate(service.NewJWTService(), "blog"), blogC.GetblogByblogname)
		// blogRoutes.PUT("/self/name", middleware.Authenticate(service.NewJWTService(), "blog"), blogC.UpdateSelfName)
		// blogRoutes.DELETE("/self", middleware.Authenticate(service.NewJWTService(), "blog"), blogC.DeleteSelfblog)
		// blogRoutes.POST("/signup", blogC.SignUp)
		// blogRoutes.POST("/signin", blogC.SignIn)
	}
}
