package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	proto "github.com/amogower/shippy/user-service/proto/user"
	"github.com/micro/go-micro/broker"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repo         Repository
	tokenService Authable
	PubSub       broker.Broker
}

func (srv *service) Get(ctx context.Context, req *proto.User, res *proto.Response) error {
	user, err := srv.repo.Get(req.Id)
	if err != nil {
		return err
	}

	res.User = user
	return nil
}

func (srv *service) GetAll(ctx context.Context, req *proto.Request, res *proto.Response) error {
	users, err := srv.repo.GetAll()
	if err != nil {
		return err
	}

	res.Users = users
	return nil
}

func (srv *service) Auth(ctx context.Context, req *proto.User, res *proto.Token) error {
	log.Println("Logging in with:", req.Email, req.Password)

	user, err := srv.repo.GetByEmail(req.Email)
	if err != nil {
		return err
	}
	log.Println(user)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return err
	}

	token, err := srv.tokenService.Encode(user)
	if err != nil {
		return err
	}

	res.Token = token
	return nil
}

func (srv *service) Create(ctx context.Context, req *proto.User, res *proto.Response) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	req.Password = string(hashedPass)
	if err := srv.repo.Create(req); err != nil {
		return err
	}

	res.User = req

	if err := srv.publishEvent("user.created", req); err != nil {
		return err
	}

	return nil
}

func (srv *service) ValidateToken(ctx context.Context, req *proto.Token, res *proto.Token) error {
	claims, err := srv.tokenService.Decode(req.Token)
	if err != nil {
		return err
	}

	log.Println(claims)

	if claims.User.Id == "" {
		return errors.New("invalid user")
	}

	res.Valid = true

	return nil
}

func (srv *service) publishEvent(topic string, user *proto.User) error {
	body, err := json.Marshal(user)
	if err != nil {
		return err
	}

	msg := &broker.Message{
		Header: map[string]string{
			"id": user.Id,
		},
		Body: body,
	}

	if err := srv.PubSub.Publish(topic, msg); err != nil {
		log.Printf("[pub] failed: %v", err)
	}

	return nil
}
