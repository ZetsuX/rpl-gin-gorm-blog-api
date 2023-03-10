package routes

import (
	"go-blogrpl/controller"
	"go-blogrpl/middleware"
	"go-blogrpl/service"

	"github.com/gin-gonic/gin"
)

func LikeRoutes(router *gin.Engine, likeC controller.LikeController) {
	blogLikeRoutes := router.Group("likes/blog")
	{
		blogLikeRoutes.GET("/", middleware.Authenticate(service.NewJWTService(), "admin"), likeC.GetAllBlogLikes)
		// blogLikeRoutes.POST("/:blogid", middleware.Authenticate(service.NewJWTService(), "user"), likeC.AddLikeForBlog)
	}

	commentLikeRoutes := router.Group("/likes/comment")
	{
		commentLikeRoutes.GET("/", middleware.Authenticate(service.NewJWTService(), "admin"), likeC.GetAllCommentLikes)
		// commentLikeRoutes.POST("/:commentid", middleware.Authenticate(service.NewJWTService(), "user"), likeC.AddLikeForComment)
	}
}
