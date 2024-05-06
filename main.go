package main

import (
	"backend-crowdfunding/auth"
	"backend-crowdfunding/handler"
	"backend-crowdfunding/user"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	//dsn := "root:biworksmanjarodb@tcp(127.0.0.1:3306)/bwa-crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:databasebily@tcp(127.0.0.1:3306)/bwa-crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailibility)
	api.POST("/avatars", userHandler.UploadAvatar)

	router.Run()
}
