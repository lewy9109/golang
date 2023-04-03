package main

import (
	// "edu/model/user"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func main() {
	fmt.Println("siema")

	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&user.User{})

	if err != nil {
		log.Fatal(err)
	}

	infraUser := user.DefaultUserInfraStructure(db)

	user := addUser()

	err = infraUser.CreateUser(user)

	if err != nil {
		log.Fatal(err)
	}

	// var id uint
	// id = 1

	// user, _ := user.UserInfrastructure.GetUser(id)

	// fmt.Println(user)
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
