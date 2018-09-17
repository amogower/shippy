package main

import (
	"github.com/jinzhu/gorm"

	proto "github.com/amogower/shippy/user-service/proto/user"
)

type Repository interface {
	GetAll() ([]*proto.User, error)
	Get(id string) (*proto.User, error)
	Create(user *proto.User) error
	GetByEmail(email string) (*proto.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func (repo *UserRepository) GetAll() ([]*proto.User, error) {
	var users []*proto.User
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepository) Get(id string) (*proto.User, error) {
	var user *proto.User
	user.Id = id
	if err := repo.db.First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) GetByEmail(email string) (*proto.User, error) {
	user := &proto.User{}
	if err := repo.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) Create(user *proto.User) error {
	if err := repo.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
