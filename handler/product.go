package handler

import (
	"hari-ketiga/tpm-keenam/model"
	"hari-ketiga/tpm-keenam/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	ProductService *service.ProductService
}

func (u *ProductHandler) Get(ctx *gin.Context) {

	products, err := u.ProductService.Get()
	if err != nil {
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
			Data:   nil,
		})
	}

	productId, err := u.ProductService.Create(&model.Product{
		Name:  productCreate.Name,
		Price: productCreate.Price,
	})

	productCreate.Id = productId

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ProductResponse{
			Status: "failed",
			Data:   nil,
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
	}
	id, _ := strconv.Atoi(idStr)
	err := u.ProductService.Delete(int(id))
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
	err := u.ProductService.Update(int(id), &productUpdate)
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
