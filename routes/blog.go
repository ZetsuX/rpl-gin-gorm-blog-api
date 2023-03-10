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
	}
}
