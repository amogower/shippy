package main

import (
	proto "github.com/amogower/shippy/vessel-service/proto/vessel"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	dbName           = "shippy"
	vesselCollection = "vessels"
)

type Repository interface {
	FindAvailable(*proto.Specification) (*proto.Vessel, error)
	Create(*proto.Vessel) error
	Close()
}

type VesselRepository struct {
	session *mgo.Session
}

func (repo *VesselRepository) FindAvailable(spec *proto.Specification) (*proto.Vessel, error) {
	var vessel *proto.Vessel

	err := repo.collection().Find(bson.M{"capacity": bson.M{"$gte": spec.Capacity}, "maxweight": bson.M{"$gte": spec.MaxWeight}}).One(&vessel)
	if err != nil {
		return nil, err
	}

	return vessel, nil
}

func (repo *VesselRepository) Create(vessel *proto.Vessel) error {
	return repo.collection().Insert(vessel)
}

func (repo *VesselRepository) Close() {
	repo.session.Close()
}

func (repo *VesselRepository) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(vesselCollection)
}
