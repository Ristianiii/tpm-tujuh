package handler

import (
	"fmt"
	"net/http"
	"time"
	"tpm-tujuh/model"
	"tpm-tujuh/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService *service.UserService
}

func (u *UserHandler) Create(ctx *gin.Context) {
	userCreate := model.User{}
	if err := ctx.Bind(&userCreate); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.UserResponse{
			Status: "failed",
			Data:   nil,
		})
	}

	hashedPass, err := service.HashPassword(userCreate.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.UserResponse{
			Status: "failed",
			Data:   nil,
		})
	}

	UserId, err := u.UserService.Create(&model.User{
		Email:    userCreate.Email,
		Password: hashedPass,
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.UserResponse{
			Status: "failed",
			Data:   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, model.UserResponse{
		Status: "Success",
		Data: map[string]interface{}{
			"Email":  userCreate.Email,
			"UserId": UserId,
		},
	})
}

func (u *UserHandler) Login(ctx *gin.Context) {
	userCreate := model.User{}
	if err := ctx.Bind(&userCreate); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.UserResponse{
			Status: "failed",
			Data:   nil,
		})
	}

	unhashedPassword := userCreate.Password

	User, err := u.UserService.Get(&model.User{
		Email: userCreate.Email,
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.UserResponse{
			Status: "failed",
			Data:   "Invalid Email/Password",
		})
		return
	}

	isPasswordMatch := service.CheckPasswordHash(unhashedPassword, User.Password)

	if !isPasswordMatch {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.UserResponse{
			Status: "failed",
			Data:   "Invalid Email/Password",
		})
		return
	}

	uid_string := fmt.Sprintf("%v", User.UserId)

	authToken, _ := service.GenerateUserJWT(uid_string, 2*time.Hour)

	ctx.JSON(http.StatusCreated, model.UserResponse{
		Status: "Success",
		Data: map[string]interface{}{
			"Token": authToken,
		},
	})
}
