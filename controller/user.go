package controller

import (
	"go-blogrpl/dto"
	"go-blogrpl/entity"
	"go-blogrpl/service"
	"go-blogrpl/utils"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService service.UserService
	jwtService  service.JWTService
}

type UserController interface {
	SignUp(ctx *gin.Context)
	SignIn(ctx *gin.Context)
	GetAllUsers(ctx *gin.Context)
	GetUserByUsername(ctx *gin.Context)
	UpdateSelfName(ctx *gin.Context)
	DeleteSelfUser(ctx *gin.Context)
}

func NewUserController(userS service.UserService, jwtS service.JWTService) UserController {
	return &userController{
		userService: userS,
		jwtService:  jwtS,
	}
}

func (userC *userController) SignUp(ctx *gin.Context) {
	var userDTO dto.UserSignUpRequest
	err := ctx.ShouldBind(&userDTO)
	if err != nil {
		resp := utils.CreateFailResponse("Failed to process user sign up request", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	newUser, err := userC.userService.CreateNewUser(ctx, userDTO)
	if err != nil {
		resp := utils.CreateFailResponse(err.Error(), http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	resp := utils.CreateSuccessResponse("user signed up successfully", http.StatusCreated, newUser)
	ctx.JSON(http.StatusCreated, resp)
}

func (userC *userController) SignIn(ctx *gin.Context) {
	var userDTO dto.UserSignInRequest
	err := ctx.ShouldBind(&userDTO)
	if err != nil {
		resp := utils.CreateFailResponse("Failed to process user sign in request", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	res := userC.userService.VerifySignIn(ctx.Request.Context(), userDTO.UserIdentifier, userDTO.Password)
	if !res {
		response := utils.CreateFailResponse("Entered credentials invalid", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	user, err := userC.userService.GetUserByIdentifier(ctx.Request.Context(), userDTO.UserIdentifier)
	if err != nil {
		response := utils.CreateFailResponse("Failed to process user sign in request", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	token := userC.jwtService.GenerateToken(user.ID, user.Role)
	authResp := utils.CreateAuthResponse(token, user.Role)
	resp := utils.CreateSuccessResponse("successfully signed in user", http.StatusOK, authResp)
	ctx.JSON(http.StatusOK, resp)
}

func (userC *userController) GetAllUsers(ctx *gin.Context) {
	users, err := userC.userService.GetAllUsers(ctx)
	if err != nil {
		resp := utils.CreateFailResponse("Failed to fetch all users", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	var resp utils.Response
	if len(users) == 0 {
		resp = utils.CreateSuccessResponse("no user found", http.StatusOK, users)
	} else {
		resp = utils.CreateSuccessResponse("successfully fetched all users", http.StatusOK, users)
	}
	ctx.JSON(http.StatusOK, resp)
}

func (userC *userController) GetUserByUsername(ctx *gin.Context) {
	username := ctx.Param("username")
	user, err := userC.userService.GetUserByUsername(ctx, username)
	if err != nil {
		resp := utils.CreateFailResponse(err.Error(), http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	var resp utils.Response
	if reflect.DeepEqual(user, entity.User{}) {
		resp = utils.CreateSuccessResponse("user not found", http.StatusOK, nil)
	} else {
		resp = utils.CreateSuccessResponse("successfully fetched user", http.StatusOK, user)
	}
	ctx.JSON(http.StatusOK, resp)
}

func (userC *userController) UpdateSelfName(ctx *gin.Context) {
	var userDTO dto.UserNameUpdateRequest
	err := ctx.ShouldBind(&userDTO)
	if err != nil {
		resp := utils.CreateFailResponse("Failed to process user name update request", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	id := ctx.GetUint64("ID")
	user, err := userC.userService.UpdateSelfName(ctx, userDTO, id)
	if err != nil {
		resp := utils.CreateFailResponse(err.Error(), http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	var resp utils.Response
	if reflect.DeepEqual(user, entity.User{}) {
		resp = utils.CreateSuccessResponse("user not found", http.StatusOK, nil)
	} else {
		resp = utils.CreateSuccessResponse("successfully updated user", http.StatusOK, user)
	}
	ctx.JSON(http.StatusOK, resp)
}

func (userC *userController) DeleteSelfUser(ctx *gin.Context) {
	id := ctx.GetUint64("ID")
	err := userC.userService.DeleteSelfUser(ctx, id)
	if err != nil {
		resp := utils.CreateFailResponse(err.Error(), http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	resp := utils.CreateSuccessResponse("successfully deleted user", http.StatusOK, nil)
	ctx.JSON(http.StatusOK, resp)
}
