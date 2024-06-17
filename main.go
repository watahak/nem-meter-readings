package main

import (
	"log"
	"nem12/db"
	docs "nem12/docs"
	"nem12/routers"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	db.Init()

	r := gin.Default()
	r.SetTrustedProxies(nil)

	docs.SwaggerInfo.BasePath = "/api/v1"

	routers.SetupRouter(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")
}
