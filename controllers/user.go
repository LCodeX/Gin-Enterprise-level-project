package controllers

import (
	"net/http"
	"yky-gin/models/dto"
	"yky-gin/services"
	"yky-gin/utils/resp"
	"yky-gin/validator"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{UserService: userService}
}

// 用户注册
// CreateUser godoc
// @Summary      Create User
// @Description  Add user
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        name     body      string  true  "UserName"
// @Param        password    body      string  true  "Email"
// @Param       phone     body      string  true  "Phone"
// @Success      200      {object}  map[string]interface{}
// @Failure      400      {object}  map[string]interface{}
// @Router       /user/create [post]
func (uc *UserController) Register(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required,min=3,max=10,starts_with_letter,alphanum"`
		Password string `json:"password" binding:"required,min=6,max=20"`
		Phone    string `json:"phone" binding:"required,zh_phone_number"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validator.TranslateError(err)})
		return
	}
	user, token, err := uc.UserService.Register(req.Username, req.Password, req.Phone)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	responseData := map[string]interface{}{
		"user":  user,
		"token": token,
	}
	resp.RespHelper.OK(c, responseData)
}

func (uc *UserController) UpdateUserPassword(c *gin.Context) {
	var req dto.UpdatePasswordRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		resp.RespHelper.Fail(c, resp.Error.Code, err.Error())
		return
	}
	user_id := c.MustGet("userID").(uint64)
	err := uc.UserService.UpdateUserPassword(req, user_id)
	if err != nil {
		resp.RespHelper.Fail(c, resp.Error.Code, err.Error())
		return
	}
	resp.RespHelper.OK(c, nil)
}

func (uc *UserController) ForgetPassword(c *gin.Context) {
	var req dto.ForgetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.RespHelper.Fail(c, resp.Error.Code, err.Error())
		return
	}
	err := uc.UserService.ForgotPassword(req)
	if err != nil {
		resp.RespHelper.Fail(c, resp.Error.Code, err.Error())
		return
	}
	resp.RespHelper.OK(c, nil)
}

func (uc *UserController) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		resp.RespHelper.Fail(c, resp.RequestDataError.Code, resp.RequestDataError.Message)
		return
	}

	authData, err := uc.UserService.Login(req.Username, req.Password)
	if err != nil {
		resp.RespHelper.Fail(c, resp.InvalidCredentials.Code, resp.InvalidCredentials.Message)
		return
	}
	resp.RespHelper.OK(c, authData)
}
