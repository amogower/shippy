package main

import (
	"context"
	"log"

	proto "github.com/amogower/shippy/user-service/proto/user"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repo         Repository
	tokenService Authable
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
	return nil
}

func (srv *service) ValidateToken(ctx context.Context, req *proto.Token, res *proto.Token) error {
	return nil
}
