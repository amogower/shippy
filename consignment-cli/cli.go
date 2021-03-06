package main

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"

	proto "github.com/amogower/shippy/consignment-service/proto/consignment"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/metadata"
)

const (
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*proto.Consignment, error) {
	var consignment *proto.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err
}

func main() {
	cmd.Init()

	client := proto.NewConsignmentServiceClient("shippy.consignment", microclient.DefaultClient)

	file := defaultFilename
	var token string
	log.Println(os.Args)

	if len(os.Args) < 3 {
		log.Fatal(errors.New("Not enough arguments, expecting file and token"))
	}

	file = os.Args[1]
	token = os.Args[2]

	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	ctx := metadata.NewContext(context.Background(), map[string]string{
		"token": token,
	})

	r, err := client.Create(ctx, consignment)
	if err != nil {
		log.Fatalf("Could not create consignment: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	getAll, err := client.Get(ctx, &proto.Request{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}

	for _, v := range getAll.Consignments {
		log.Println(v)
	}
}
