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
	GetAllComments(ctx *gin.Context)
	PostCommentForBlog(ctx *gin.Context)
}

func NewCommentController(commentS service.CommentService, jwtS service.JWTService) CommentController {
	return &commentController{
		commentService: commentS,
		jwtService:     jwtS,
	}
}

func (commentC *commentController) GetAllComments(ctx *gin.Context) {
	comments, err := commentC.commentService.GetAllComments(ctx)
	if err != nil {
		resp := utils.CreateFailResponse("Failed to fetch all comments", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	var resp utils.Response
	if len(comments) == 0 {
		resp = utils.CreateSuccessResponse("no comment found", http.StatusOK, comments)
	} else {
		resp = utils.CreateSuccessResponse("successfully fetched all comments", http.StatusOK, comments)
	}
	ctx.JSON(http.StatusOK, resp)
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
