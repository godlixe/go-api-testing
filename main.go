package main

import (
	"api-testing/db"
	"api-testing/user"

	"github.com/gin-gonic/gin"
)

func main() {
	db := db.New()
	repository := user.NewRepository(db)
	service := user.NewService(repository)
	handler := user.NewHandler(service)

	router := gin.Default()

	router.GET("/", handler.GetAll)
	router.GET("/:id", handler.Get)
	router.POST("/", handler.Create)

	router.Run(":8080")
}
