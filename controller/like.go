package controller

import (
	"go-blogrpl/service"
	"go-blogrpl/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type likeController struct {
	likeService service.LikeService
	jwtService  service.JWTService
}

type LikeController interface {
	// BlogLikes
	GetAllBlogLikes(ctx *gin.Context)
	ChangeLikeForBlog(ctx *gin.Context)

	// CommentLikes
	GetAllCommentLikes(ctx *gin.Context)
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

func (likeC *likeController) ChangeLikeForBlog(ctx *gin.Context) {
	blogID, err := strconv.ParseUint(ctx.Param("blogid"), 10, 64)
	if err != nil {
		resp := utils.CreateFailResponse("Failed to process id of blog for like request", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	userID := ctx.GetUint64("ID")
	msg, err := likeC.likeService.ChangeLikeForBlog(ctx, blogID, userID)
	if err != nil {
		resp := utils.CreateFailResponse("Failed to process like request", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	resp := utils.CreateSuccessResponse(msg, http.StatusOK, nil)
	ctx.JSON(http.StatusOK, resp)
}

func (likeC *likeController) GetAllCommentLikes(ctx *gin.Context) {
	likes, err := likeC.likeService.GetAllCommentLikes(ctx)
	if err != nil {
		resp := utils.CreateFailResponse("Failed to fetch all comment likes", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	var resp utils.Response
	if len(likes) == 0 {
		resp = utils.CreateSuccessResponse("no comment like found", http.StatusOK, likes)
	} else {
		resp = utils.CreateSuccessResponse("successfully fetched all comment likes", http.StatusOK, likes)
	}
	ctx.JSON(http.StatusOK, resp)
}
