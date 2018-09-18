package main

import (
	"fmt"
	"time"

	proto "github.com/amogower/shippy/user-service/proto/user"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro"
)

func main() {
	var db *gorm.DB
	for attempts := 0; attempts < 10; attempts++ {
		var err error
		db, err = CreateConnection()
		if err == nil {
			break
		}
		time.Sleep(time.Second)
	}

	defer db.Close()

	db.AutoMigrate(&proto.User{})

	repo := &UserRepository{db}

	tokenService := &TokenService{repo}

	srv := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)

	srv.Init()

	pubsub := srv.Server().Options().Broker

	proto.RegisterUserServiceHandler(srv.Server(), &service{repo, tokenService, pubsub})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
