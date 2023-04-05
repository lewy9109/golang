package main

import (
	"edu/pkg/helper"
	"edu/pkg/user"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
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

	infraUser := user.DefaultUserInfraStructure(db)
	InsertUserToDB(infraUser)
	// UpdateUser(infraUser, uint(1))
	// GetByEmail(infraUser, "mail@glob.com")
	// TokenHelper()
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
