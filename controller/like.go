package controller

import (
	"go-blogrpl/service"
	"go-blogrpl/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type likeController struct {
	likeService service.LikeService
	jwtService  service.JWTService
}

type LikeController interface {
	// BlogLikes
	GetAllBlogLikes(ctx *gin.Context)

	// CommentLikes
}

func NewLikeController(likeS service.LikeService, jwtS service.JWTService) LikeController {
	return &likeController{
		likeService: likeS,
		jwtService:  jwtS,
	}
}

func (likeC *likeController) GetAllBlogLikes(ctx *gin.Context) {
	likes, err := likeC.likeService.GetAllBlogLikes(ctx)
	if err != nil {
		resp := utils.CreateFailResponse("Failed to fetch all blog likes", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	var resp utils.Response
	if len(likes) == 0 {
		resp = utils.CreateSuccessResponse("no blog like found", http.StatusOK, likes)
	} else {
		resp = utils.CreateSuccessResponse("successfully fetched all blog likes", http.StatusOK, likes)
	}
	ctx.JSON(http.StatusOK, resp)
}
