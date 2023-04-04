package main

import (
	"edu/model/user"
	"edu/seciurity"
	"fmt"

	"log"

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
	infraUser := user.DefaultUserInfraStructure(db)
	user, err := infraUser.GetByEmail("mail@glob.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user)

	// err = infraUser.UpdateUserAccessToken(user.ID, "123qwerty")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// user, err = infraUser.GetByEmail("mail@glob.com")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(user)

	token, err := helper.CreateJWTToken(uint(user.ID), "secretToken")
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

func addUser() user.User {

	newUser := user.User{
		FirstName: "Xavi",
		LastName:  "Fernando",
		Email:     "mail@glob.com",
		Password:  "qwerty123",
	}

	return newUser
}
