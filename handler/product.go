package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"tpm-tujuh/model"
	"tpm-tujuh/service"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	ProductService *service.ProductService
}

func (u *ProductHandler) Get(ctx *gin.Context) {
	userId := ctx.GetString("user_id")
	uid_int, err := strconv.Atoi(userId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ProductResponse{
			Status: "failed",
			Data:   nil,
		})
	}

	products, err := u.ProductService.Get(uid_int)
	if err != nil {
		fmt.Print(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ProductResponse{
			Status: "failed",
			Data:   nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, model.ProductResponse{
		Status: "success",
		Data:   products,
	})
}

func (u *ProductHandler) Create(ctx *gin.Context) {
	productCreate := model.Product{}
	if err := ctx.Bind(&productCreate); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ProductResponse{
			Status: "failed",
			Data:   err.Error(),
		})
		return
	}

	userId := ctx.GetString("user_id")
	uid_int, err := strconv.Atoi(userId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ProductResponse{
			Status: "failed",
			Data:   err.Error(),
		})
		return
	}

	productId, err := u.ProductService.Create(&model.Product{
		Name:   productCreate.Name,
		Price:  productCreate.Price,
		UserId: uid_int,
	})

	productCreate.Id = productId
	productCreate.UserId = uid_int

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ProductResponse{
			Status: "failed",
			Data:   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, model.ProductResponse{
		Status: "Success",
		Data:   productCreate,
	})
}

func (u *ProductHandler) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	if idStr == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ProductResponse{
			Status: "failed",
			Data:   nil,
		})
		return
	}
	id, _ := strconv.Atoi(idStr)

	userId := ctx.GetString("user_id")
	uid_int, err := strconv.Atoi(userId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ProductResponse{
			Status: "failed",
			Data:   err.Error(),
		})
		return
	}

	err = u.ProductService.Delete(int(id), uid_int)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ProductResponse{
			Status: "failed",
			Data:   nil,
		})
		return
	}
	ctx.JSON(http.StatusCreated, model.ProductResponse{
		Status: "Success",
		Data:   nil,
	})
}

func (u *ProductHandler) Update(ctx *gin.Context) {
	idStr := ctx.Param("id")
	if idStr == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ProductResponse{
			Status: "failed",
			Data:   nil,
		})
	}
	id, _ := strconv.Atoi(idStr)
	productUpdate := model.ProductUpdate{}
	if err := ctx.Bind(&productUpdate); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ProductResponse{
			Status: "failed",
			Data:   nil,
		})
	}

	userId := ctx.GetString("user_id")
	uid_int, err := strconv.Atoi(userId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ProductResponse{
			Status: "failed",
			Data:   err.Error(),
		})
		return
	}

	err = u.ProductService.Update(int(id), uid_int, &productUpdate)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ProductResponse{
			Status: "failed",
			Data:   nil,
		})
		return
	}
	ctx.JSON(http.StatusCreated, model.ProductResponse{
		Status: "Success",
		Data:   productUpdate,
	})
}
