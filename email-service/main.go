package main

import (
	"context"
	"log"

	proto "github.com/amogower/shippy/auth-service/proto/auth"
	"github.com/micro/go-micro"
)

const topic = "user.created"

type Subscriber struct{}

func (sub *Subscriber) Process(ctx context.Context, user *proto.User) error {
	log.Println("Picked up a new message")
	log.Println("Sending email to:", user.Name)
	return nil
}

func main() {
	srv := micro.NewService(
		micro.Name("shippy.email"),
		micro.Version("latest"),
	)

	srv.Init()

	micro.RegisterSubscriber(topic, srv.Server(), new(Subscriber))

	if err := srv.Run(); err != nil {
		log.Println(err)
	}
}

func sendEmail(user *proto.User) error {
	log.Println("Sending email to: ", user.Name)
	return nil
}
