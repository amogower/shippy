package main

import (
	"fmt"
	"log"

	proto "github.com/amogower/shippy/user-service/proto/user"
	"github.com/micro/go-micro"
)

func main() {
	db, err := CreateConnection()
	defer db.Close()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	db.AutoMigrate(&proto.User{})

	repo := &UserRepository{db}

	tokenService := &TokenService{repo}

	srv := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)

	srv.Init()

	proto.RegisterUserServiceHandler(srv.Server(), &service{repo, tokenService})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
