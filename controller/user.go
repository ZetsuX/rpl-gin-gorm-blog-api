package controller

import (
	"go-blogrpl/dto"
	"go-blogrpl/service"
	"go-blogrpl/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService service.UserService
}

type UserController interface {
	SignUp(ctx *gin.Context)
	GetAll(ctx *gin.Context)
}

func NewUserController(userS service.UserService) UserController {
	return &userController{userService: userS}
}

func (userC *userController) SignUp(ctx *gin.Context) {
	var userDTO dto.UserSignUpRequest
	err := ctx.ShouldBind(&userDTO)
	if err != nil {
		resp := utils.CreateResponse("Failed to process user sign up request", http.StatusBadRequest, nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	newUser, err := userC.userService.CreateNewUser(ctx, userDTO)
	if err != nil {
		resp := utils.CreateResponse(err.Error(), http.StatusBadRequest, nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	resp := utils.CreateResponse("user signed up successfully", http.StatusCreated, newUser)
	ctx.JSON(http.StatusCreated, resp)
}

func (userC *userController) GetAll(ctx *gin.Context) {

	users, err := userC.userService.GetAllUsers(ctx)
	if err != nil {
		resp := utils.CreateResponse("Failed to fetch all users", http.StatusBadRequest, nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}
	resp := utils.CreateResponse("user signed up successfully", http.StatusCreated, users)
	ctx.JSON(http.StatusCreated, resp)
}
