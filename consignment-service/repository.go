package main

import (
	proto "github.com/amogower/shippy/consignment-service/proto/consignment"
	mgo "gopkg.in/mgo.v2"
)

const (
	dbName                = "shippy"
	consignmentCollection = "consignments"
)

type Repository interface {
	Create(*proto.Consignment) error
	GetAll() ([]*proto.Consignment, error)
	Close()
}

type ConsignmentRepository struct {
	session *mgo.Session
}

func (repo *ConsignmentRepository) Create(consignment *proto.Consignment) error {
	return repo.collection().Insert(consignment)
}

func (repo *ConsignmentRepository) GetAll() ([]*proto.Consignment, error) {
	var consignments []*proto.Consignment

	err := repo.collection().Find(nil).All(&consignments)
	return consignments, err
}

func (repo *ConsignmentRepository) Close() {
	repo.session.Close()
}

func (repo *ConsignmentRepository) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(consignmentCollection)
}
