package grpc

import (
	"context"

	"github.com/98y7tbnb97t/GoMicro/proto/userpb"
	"github.com/98y7tbnb97t/users-service/internal/user"
)

type Handler struct {
	svc *user.Service
	userpb.UnimplementedUserServiceServer
}

func NewHandler(svc *user.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	u := &user.User{
		Email: req.Email,
	}

	err := h.svc.CreateUser(u)
	if err != nil {
		return nil, err
	}

	return &userpb.CreateUserResponse{
		User: &userpb.User{
			Id:    uint32(u.ID),
			Email: u.Email,
		},
	}, nil
}

func (h *Handler) GetUser(ctx context.Context, req *userpb.UserRequest) (*userpb.UserResponse, error) {
	u, err := h.svc.GetUserByID(int(req.Id))
	if err != nil {
		return nil, err
	}
	return &userpb.UserResponse{
		User: &userpb.User{
			Id:    uint32(u.ID),
			Email: u.Email,
		},
	}, nil
}

// ListUsers now also returns a top-level User object (the first user in the list, or nil if empty)
func (h *Handler) ListUsers(ctx context.Context, req *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {
	users, err := h.svc.GetAllUsers()
	if err != nil {
		return nil, err
	}
	var pbUsers []*userpb.User
	for _, u := range users {
		pbUsers = append(pbUsers, &userpb.User{
			Id:    uint32(u.ID),
			Email: u.Email,
		})
	}

	var mainUser *userpb.User
	if len(pbUsers) > 0 {
		mainUser = pbUsers[0]
	}

	return &userpb.ListUsersResponse{
		Users: pbUsers,
		User:  mainUser, // new field: top-level user object
	}, nil
}

func (h *Handler) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	updates := &user.User{
		Email: req.Email,
	}
	err := h.svc.UpdateUser(int(req.Id), updates)
	if err != nil {
		return nil, err
	}
	u, err := h.svc.GetUserByID(int(req.Id))
	if err != nil {
		return nil, err
	}
	return &userpb.UpdateUserResponse{
		User: &userpb.User{
			Id:    uint32(u.ID),
			Email: u.Email,
		},
	}, nil
}

func (h *Handler) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	err := h.svc.DeleteUser(int(req.Id))
	if err != nil {
		return nil, err
	}
	return &userpb.DeleteUserResponse{}, nil
}
