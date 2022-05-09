package server

import (
    "context"
    userpb "github.com/jha929/lego-grpc/protos/user"
    "github.com/jha929/lego-grpc/mock"
)

type UserPbServer struct {
    userpb.UserServer
}

func (s *UserPbServer) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	userID := req.UserId

	var userMessage *userpb.UserMessage
	for _, u := range mock.Users {
		if u.UserId != userID {
			continue
		}
		userMessage = u
		break
	}

	return &userpb.GetUserResponse{UserMessage: userMessage}, nil
}

func (s *UserPbServer) ListUsers(ctx context.Context, in *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {
    userMessages := make([]*userpb.UserMessage, len(mock.Users))
	for i, u := range mock.Users {
		userMessages[i] = u
	}

	return &userpb.ListUsersResponse{
		UserMessages: userMessages,
	}, nil
}
