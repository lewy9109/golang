package main

import (
	"edu/pkg/adapthttp"
	"edu/pkg/helper"
	"edu/pkg/user"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("start")
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&user.User{})
	if err != nil {
		log.Fatal(err)
	}

	userInfra := user.DefaultUserInfraStructure(db)
	userService := user.DefalutUserService(userInfra, "secretToken")

	userServer := adapthttp.DefalutUserServer(userService)

	server := gin.Default()

	group := server.Group("/user/", userServer.Authorize)
	{
		group.GET("/", userServer.GetInfo)
	}

	server.POST("/users", userServer.CreateUser)
	server.GET("/login", userServer.LoginUser)

	// user, _ := userServie.GetUserInfo(2)

	// fmt.Println(user)

	// token, _ := userServie.Login("mail2@glob.com", "qwerty123")

	// userID, _ := userServie.Authorize(token)

	// fmt.Println(userID)

	// fmt.Println(token)

}

func TokenHelper() {
	token, err := helper.CreateJWTToken(uint(1), "secretToken")
	if err != nil {
		fmt.Println("KUPA")
		log.Fatal(err)
	}

	fmt.Println(token)

	parsedToken, err := helper.ValidateJWTToken(token, "secretToken")
	if err != nil {
		log.Fatal(err)
	}

	parsedUser := parsedToken.Claims.(jwt.MapClaims)

	fmt.Println(parsedUser)
}

func InsertUserToDB(infraUser user.UserInfrastructure) {

	user := addUser()

	err := infraUser.CreateUser(user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user)

}

func UpdateUser(infraUser user.UserInfrastructure, userID uint) {
	err := infraUser.UpdateUserAccessToken(userID, "123qwerty")
	if err != nil {
		log.Fatal(err)
	}
}

func GetByEmail(infraUser user.UserInfrastructure, email string) {
	user, err := infraUser.GetByEmail(email)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user)
}

func addUser() user.User {
	hashedPassword, err := helper.HashPassword("qwerty123")

	if err != nil {
		log.Fatal(err)
	}

	newUser := user.User{
		FirstName: "Test",
		LastName:  "Fernando",
		Email:     "mail2@glob.com",
		Password:  hashedPassword,
	}

	return newUser
}
