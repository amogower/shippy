package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	authService "github.com/amogower/shippy/auth-service/proto/auth"
	proto "github.com/amogower/shippy/consignment-service/proto/consignment"
	vesselService "github.com/amogower/shippy/vessel-service/proto/vessel"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
)

const (
	defaultHost = "localhost:27017"
)

func main() {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = defaultHost
	}

	session, err := CreateSession(host)
	if err != nil {
		log.Panicf("Could not connect to datastore with host %s - %v", host, err)
	}

	defer session.Close()

	srv := micro.NewService(
		micro.Name("shippy.consignment"),
		micro.WrapHandler(AuthWrapper),
	)

	vesselClient := vesselService.NewVesselServiceClient("shippy.vessel", srv.Client())

	srv.Init()

	proto.RegisterConsignmentServiceHandler(srv.Server(), &service{session, vesselClient})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}

func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		if os.Getenv("DISABLE_AUTH") == "true" {
			return fn(ctx, req, resp)
		}

		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("No auth metadata found in request")
		}

		token := meta["Token"]
		log.Println("Working...")
		log.Println("Authenticating with token: ", token)

		authClient := authService.NewAuthServiceClient("shippy.auth", client.DefaultClient)
		_, err := authClient.ValidateToken(context.Background(), &authService.Token{
			Token: token,
		})
		if err != nil {
			return err
		}

		err = fn(ctx, req, resp)
		return err
	}
}
