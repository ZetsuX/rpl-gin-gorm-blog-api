package routes

import (
	"go-blogrpl/controller"
	"go-blogrpl/middleware"
	"go-blogrpl/service"

	"github.com/gin-gonic/gin"
)

func CommentRoutes(router *gin.Engine, commentC controller.CommentController) {
	commentRoutes := router.Group("/blog/comments")
	{
		commentRoutes.GET("/", middleware.Authenticate(service.NewJWTService(), "admin"), commentC.GetAllComments)
		// commentRoutes.GET("/:blogid", middleware.Authenticate(service.NewJWTService(), "user"), commentC.GetBlogComment)
		commentRoutes.POST("/:blogid", middleware.Authenticate(service.NewJWTService(), "user"), commentC.PostCommentForBlog)
	}
}
