package service

import (
	"context"
	v1 "kratos-realworld/api/realworld/v1"
	"kratos-realworld/internal/errors"
)

func (s *RealWorldService) Login(ctx context.Context, req *v1.LoginRequest) (reply *v1.UserReply, err error) {
	if len(req.User.Email) == 0 {
		return nil, errors.NewHTTPError(422, "email is required", "email is required")
	}
	return &v1.UserReply{
		User: &v1.UserReply_User{
			Username: "Username",
		},
	}, nil
}
func (s *RealWorldService) Register(ctx context.Context, req *v1.RegisterRequest) (reply *v1.UserReply, err error) {
	u, err := s.uc.Register(ctx, req.User.Username, req.User.Email, req.User.Password)
	if err != nil {
		return nil, err
	}
	return &v1.UserReply{
		User: &v1.UserReply_User{
			Email:    u.Email,
			Username: u.Username,
			Token:    u.Token,
		},
	}, nil
}
func (s *RealWorldService) GetCurrentUser(ctx context.Context, req *v1.GetCurrentUserRequest) (reply *v1.UserReply, err error) {
	return &v1.UserReply{
		User: &v1.UserReply_User{
			Username: "boom",
			Email:    "this is a test email",
			Token:    "this is a test token",
		},
	}, nil
}
func (s *RealWorldService) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (rep *v1.UserReply, err error) {
	return &v1.UserReply{}, nil
}
