package main

import (
	"context"
	"fmt"
	"log"

	proto "github.com/amogower/shippy/consignment-service/proto/consignment"
	vesselProto "github.com/amogower/shippy/vessel-service/proto/vessel"
	"github.com/gofrs/uuid"
	mgo "gopkg.in/mgo.v2"
)

type service struct {
	session      *mgo.Session
	vesselClient vesselProto.VesselServiceClient
}

func (s *service) GetRepo() Repository {
	return &ConsignmentRepository{s.session.Clone()}
}

func (s *service) Create(ctx context.Context, req *proto.Consignment, res *proto.Response) error {
	log.Println("Creating consignment: ", req)

	repo := s.GetRepo()
	defer repo.Close()

	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	if err != nil {
		return fmt.Errorf("error finding vessel: %v", err)
	}
	log.Printf("Found vessel: %s \n", vesselResponse.Vessel.Name)

	req.VesselId = vesselResponse.Vessel.Id

	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}

	req.Id = uuid.String()

	err = repo.Create(req)
	if err != nil {
		return fmt.Errorf("error creating consignment: %v", err)
	}

	res.Created = true
	res.Consignment = req
	return nil
}

func (s *service) Get(ctx context.Context, req *proto.Request, res *proto.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	consignments, err := repo.GetAll()
	if err != nil {
		return fmt.Errorf("error getting consignments: %v", err)
	}

	res.Consignments = consignments
	log.Printf("Fetched consignments: %v", consignments)
	return nil
}
