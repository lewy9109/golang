package main

import (
	"edu/pkg/controller/rest/userController"
	"edu/pkg/user"
	"log"
	"net"

	grpcUser "edu/pkg/controller/grpc/userController"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
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

	userServer := userController.DefalutUserServer(userService)

	startHttpServer(userServer)

}

func startHttpServer(userServer userController.UserServerInterface) {

	server := gin.Default()

	group := server.Group("/user/", userServer.Authorize)
	{
		group.GET("/", userServer.GetInfo)
	}

	server.POST("/users", userServer.CreateUser)
	server.POST("/login", userServer.LoginUser)

	err := server.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

func startGRPCServer() {
	listener, err := net.Listen("tcp", ":8081")

	if err != nil {
		log.Fatalln("Failed to start server")
	}

	grpcServer := grpc.NewServer()

	userServer := grpcUser.DefaultGrpcUserService()

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalln("Failed to start grpc server.")
	}
}
