package main

import (
	"edu/model/user"
	"fmt"
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

	// var id uint
	// id = 1

	// user, _ := user.UserInfrastructure.GetUser(id)

	// fmt.Println(user)
}
