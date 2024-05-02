package main

import (
	"backend-crowdfunding/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "root:biworksmanjarodb@tcp(127.0.0.1:3306)/bwa-crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userInput := user.RegisterUserInput{}
	userInput.Name = "Erlangga"
	userInput.Email = "erlangga@gmail.com"
	userInput.Occupation = "actor"
	userInput.Password = "password"

	userService.RegisterUser(userInput)
}
