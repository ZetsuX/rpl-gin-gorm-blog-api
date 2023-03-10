package controller

import (
	"go-blogrpl/dto"
	"go-blogrpl/service"
	"go-blogrpl/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type commentController struct {
	commentService service.CommentService
	jwtService     service.JWTService
}

type CommentController interface {
	// // Comments
	// GetAllComments(ctx *gin.Context)

	// BlogComments
	PostCommentForBlog(ctx *gin.Context)
}

func NewCommentController(commentS service.CommentService, jwtS service.JWTService) CommentController {
	return &commentController{
		commentService: commentS,
		jwtService:     jwtS,
	}
}

func (commentC *commentController) PostCommentForBlog(ctx *gin.Context) {
	var bcDTO dto.CommentRequest
	err := ctx.ShouldBind(&bcDTO)
	if err != nil {
		resp := utils.CreateFailResponse("Failed to process comment request", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	blogID, err := strconv.ParseUint(ctx.Param("blogid"), 10, 64)
	if err != nil {
		resp := utils.CreateFailResponse("Failed to process id of comment for blog request", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	userID := ctx.GetUint64("ID")
	newBlogComment, err := commentC.commentService.CreateNewBlogComment(ctx, bcDTO, blogID, userID)
	if err != nil {
		resp := utils.CreateFailResponse(err.Error(), http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	resp := utils.CreateSuccessResponse("comment posted successfully", http.StatusCreated, newBlogComment)
	ctx.JSON(http.StatusCreated, resp)
}
