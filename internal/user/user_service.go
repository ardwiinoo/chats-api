package user

import (
	"chats-api/util"
	"context"
	"strconv"
	"time"
)

type service struct {
	Repository
	timeout time.Duration
}

// CreateUser implements Service.
func (s *service) CreateUser(ctx context.Context, req *CreateUserReq) (*CreateUserRes, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	u := &User{
		Username: req.Username,
		Email: req.Email,
		Password: hashedPassword,
	}

	r, err := s.Repository.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}

	res := &CreateUserRes{
		ID: strconv.Itoa(int(r.ID)),
		Username: r.Username,
		Email: r.Email,
	}

	return res, nil
}

func NewService(repository Repository) Service {
	return &service{
		repository,
		time.Duration(2) * time.Second,
	}
}