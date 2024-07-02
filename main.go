package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"tpm-tujuh/handler"
	"tpm-tujuh/middleware"
	"tpm-tujuh/model"
	"tpm-tujuh/repository"
	"tpm-tujuh/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

func main() {
	ge := gin.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	ge.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK,
			map[string]any{
				"status": "OK!",
			})
	})

	pgHost := os.Getenv("PG_HOST")
	pgUser := os.Getenv("PG_USER")
	pgPassword := os.Getenv("PG_PASSWORD")
	pgDB := os.Getenv("PG_DB")
	pgPort := os.Getenv("PG_PORT")

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Jakarta",
		pgHost,
		pgUser,
		pgPassword,
		pgDB,
		pgPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Host: %v \n", pgHost)
		panic(err)
	}

	err = db.AutoMigrate(&model.Product{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}

	productPgRepo := &repository.ProductPgRepo{DB: db}
	productService := &service.ProductService{ProductPgRepo: productPgRepo}
	productHandler := &handler.ProductHandler{ProductService: productService}

	userPgRepo := &repository.UserPgRepo{DB: db}
	userService := &service.UserService{UserPgRepo: userPgRepo}
	userHandler := &handler.UserHandler{UserService: userService}

	userGroup := ge.Group("/auth")
	userGroup.POST("/register", userHandler.Create)
	userGroup.POST("/login", userHandler.Login)

	productGroup := ge.Group("/products")
	productGroup.Use(middleware.BearerAuthorization())
	productGroup.GET("", productHandler.Get)
	productGroup.POST("", productHandler.Create)
	productGroup.PUT("/:id", productHandler.Update)
	productGroup.DELETE("/:id", productHandler.Delete)

	if err := ge.Run(":8080"); err != nil {
		panic(err)
	}
}
