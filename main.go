package main

import (
	"backend-crowdfunding/handler"
	"backend-crowdfunding/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	// dsn := "root:biworksmanjarodb@tcp(127.0.0.1:3306)/bwa-crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:databasebily@tcp(127.0.0.1:3306)/bwa-crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	input := user.LoginInput{
		Email:    "bily@gmail.com",
		Password: "anjay",
	}

	user, err := userService.Login(input)
	if err != nil {
		fmt.Println("terjadi kesalahan")
		fmt.Println(err.Error())
	}

	fmt.Println(user.Email)
	fmt.Println(user.Name)

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)

	router.Run()
	//userInput := user.RegisterUserInput{}
	//userInput.Name = "Erlangga"
	//userInput.Email = "erlangga@gmail.com"
	//userInput.Occupation = "actor"
	//userInput.Password = "password"
	//
	//userService.RegisterUser(userInput)
}
