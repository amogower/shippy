package main

import (
	"fmt"
	"log"
	"os"

	proto "github.com/amogower/shippy/vessel-service/proto/vessel"
	"github.com/micro/go-micro"
)

const (
	defaultHost = "localhost:27017"
)

func createDummyData(repo Repository) {
	defer repo.Close()
	vessels := []*proto.Vessel{{Id: "vessel001", Name: "Kane's Salty Secret", MaxWeight: 200000, Capacity: 500}}
	for _, v := range vessels {
		repo.Create(v)
	}
}

func main() {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = defaultHost
	}

	session, err := CreateSession(host)
	if err != nil {
		log.Fatalf("Error connecting to datastore: %v", err)
	}

	defer session.Close()

	repo := &VesselRepository{session.Copy()}

	createDummyData(repo)

	srv := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)

	srv.Init()

	proto.RegisterVesselServiceHandler(srv.Server(), &service{session})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
