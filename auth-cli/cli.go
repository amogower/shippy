package main

import (
	"context"
	"log"
	"os"

	proto "github.com/amogower/shippy/auth-service/proto/auth"
	"github.com/micro/go-micro"
	microclient "github.com/micro/go-micro/client"
)

func main() {
	srv := micro.NewService(micro.Name("shippy.auth-cli"))

	srv.Init()

	client := proto.NewAuthServiceClient("shippy.auth", microclient.DefaultClient)

	name := "Andrew Gower"
	email := "amogower@gmail.com"
	password := "rtw987"
	company := "Capability"

	r, err := client.Create(context.TODO(), &proto.User{
		Name:     name,
		Email:    email,
		Password: password,
		Company:  company,
	})
	if err != nil {
		log.Fatalf("Could not create: %v", err)
	}
	log.Printf("Created: %s", r.User.Id)

	getAll, err := client.GetAll(context.Background(), &proto.Request{})
	if err != nil {
		log.Fatalf("Could not list users: %v", err)
	}

	for _, v := range getAll.Users {
		log.Println(v)
	}

	authResponse, err := client.Auth(context.TODO(), &proto.User{
		Email:    email,
		Password: password,
	})
	if err != nil {
		log.Fatalf("Could not authenticate user: %s error: $v\n", email, err)
	}

	log.Printf("Your access token is: %s \n", authResponse.Token)

	os.Exit(0)
}
