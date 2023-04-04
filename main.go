package main

import (
	"edu/model/user"
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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

	user, err := infraUser.GetByEmail("mail@glob.com")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user)

	err = infraUser.UpdateUserAccessToken(user.ID, "123qwerty")
	if err != nil {
		log.Fatal(err)
	}

	user, err = infraUser.GetByEmail("mail@glob.com")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user)
	// user := addUser()
	// err = infraUser.CreateUser(user)
	// if err != nil {
	// 	log.Fatal(err)
	// }

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
