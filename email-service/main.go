package main

import (
	"encoding/json"
	"log"

	proto "github.com/amogower/shippy/user-service/proto/user"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
)

func main() {
	srv := micro.NewService(
		micro.Name("go.micro.srv.email"),
		micro.Version("latest"),
	)

	srv.Init()

	pubsub := srv.Server().Options().Broker
	if err := pubsub.Connect(); err != nil {
		log.Fatal(err)
	}

	_, err := pubsub.Subscribe("user.created", func(p broker.Publication) error {
		var user *proto.User
		if err := json.Unmarshal(p.Message().Body, &user); err != nil {
			return err
		}

		log.Println(user)
		go sendEmail(user)

		return nil
	})
	if err != nil {
		log.Println(err)
	}

	if err := srv.Run(); err != nil {
		log.Println(err)
	}
}

func sendEmail(user *proto.User) error {
	log.Println("Sending email to: ", user.Name)
	return nil
}
