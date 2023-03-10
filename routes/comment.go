package routes

import (
	"go-blogrpl/controller"
	"go-blogrpl/middleware"
	"go-blogrpl/service"

	"github.com/gin-gonic/gin"
)

func CommentRoutes(router *gin.Engine, commentC controller.CommentController) {
	blogCommentRoutes := router.Group("/blog/comments")
	{
		blogCommentRoutes.POST("/:blogid", middleware.Authenticate(service.NewJWTService(), "user"), commentC.PostCommentForBlog)
	}
}
