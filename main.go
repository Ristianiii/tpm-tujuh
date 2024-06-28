package main

import (
	"hari-ketiga/tpm-keenam/handler"
	"hari-ketiga/tpm-keenam/model"
	"hari-ketiga/tpm-keenam/repository"
	"hari-ketiga/tpm-keenam/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	ge := gin.New()

	ge.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK,
			map[string]any{
				"status": "OK!",
			})
	})

	dsn := "host=localhost user=postgres password=123456 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&model.Product{})
	if err != nil {
		panic(err)
	}
	productPgRepo := &repository.ProductPgRepo{DB: db}
	productService := &service.ProductService{ProductPgRepo: productPgRepo}
	productHandler := &handler.ProductHandler{ProductService: productService}

	productGroup := ge.Group("/products")
	productGroup.GET("", productHandler.Get)
	productGroup.POST("", productHandler.Create)
	productGroup.PUT("/:id", productHandler.Update)
	productGroup.DELETE("/:id", productHandler.Delete)
	if err := ge.Run(":8080"); err != nil {
		panic(err)
	}
}
