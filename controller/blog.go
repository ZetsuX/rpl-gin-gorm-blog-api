package controller

import (
	"go-blogrpl/service"
	"go-blogrpl/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type blogController struct {
	blogService service.BlogService
	jwtService  service.JWTService
}

type BlogController interface {
	GetAllBlogs(ctx *gin.Context)
}

func NewBlogController(blogS service.BlogService, jwtS service.JWTService) BlogController {
	return &blogController{
		blogService: blogS,
		jwtService:  jwtS,
	}
}

func (blogC *blogController) GetAllBlogs(ctx *gin.Context) {
	blogs, err := blogC.blogService.GetAllBlogs(ctx)
	if err != nil {
		resp := utils.CreateFailResponse("Failed to fetch all blogs", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	var resp utils.Response
	if len(blogs) == 0 {
		resp = utils.CreateSuccessResponse("no blog found", http.StatusOK, blogs)
	} else {
		resp = utils.CreateSuccessResponse("successfully fetched all blogs", http.StatusOK, blogs)
	}
	ctx.JSON(http.StatusOK, resp)
}
