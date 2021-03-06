package main

import (
	"context"
	"fmt"

	proto "github.com/amogower/shippy/vessel-service/proto/vessel"
	"gopkg.in/mgo.v2"
)

type service struct {
	session *mgo.Session
}

func (s *service) GetRepo() Repository {
	return &VesselRepository{s.session.Clone()}
}

func (s *service) FindAvailable(ctx context.Context, req *proto.Specification, res *proto.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	vessel, err := repo.FindAvailable(req)
	if err != nil {
		return fmt.Errorf("error finding available vessel: %v", err)
	}

	res.Vessel = vessel
	return nil
}

func (s *service) Create(ctx context.Context, req *proto.Vessel, res *proto.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	if err := repo.Create(req); err != nil {
		return fmt.Errorf("error creating vessel: %v", err)
	}

	res.Vessel = req
	res.Created = true
	return nil
}
