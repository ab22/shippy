package main

import (
	"context"

	pb "github.com/ab22/shippy/user-service/proto/user"
)

type handler struct {
	repo         repository
	tokenService TokenEncoder
}

func (s *handler) Get(ctx context.Context, req *pb.User, res *pb.Response) error {
	user, err := s.repo.Get(req.Id)

	if err != nil {
		return err
	}

	res.User = user
	return nil
}

func (s *handler) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
	users, err := s.repo.GetAll()

	if err != nil {
		return err
	}

	res.Users = users
	return nil
}

func (s *handler) Create(ctx context.Context, req *pb.User, res *pb.Response) error {
	if err := s.repo.Create(req); err != nil {
		return err
	}

	res.User = req
	return nil
}

func (s *handler) ValidateToken(ctx context.Context, req *pb.User, res *pb.Response) error {
	return nil
}
