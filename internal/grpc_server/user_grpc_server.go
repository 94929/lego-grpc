package grpcserver

import (
	"context"
	"strconv"

	"github.com/jha929/lego-grpc/internal/repository"
	userpb "github.com/jha929/lego-grpc/protos/user"
)

type UserGrpcServer struct {
	userpb.UserServer
	UserRepository repository.UserRepository
}

func (ugs UserGrpcServer) CreateUser(
	ctx context.Context, req *userpb.CreateUserRequest,
) (*userpb.GetUserResponse, error) {
	var userMessage *userpb.UserMessage
	user := repository.User{Nickname: req.Nickname}
	ugs.UserRepository.CreateUser(&user)
	userIdStr := strconv.FormatUint(uint64(user.ID), 10)
	userMessage = &userpb.UserMessage{
		UserId:   userIdStr,
		Nickname: user.Nickname,
	}
	return &userpb.GetUserResponse{UserMessage: userMessage}, nil
}

func (ugs UserGrpcServer) GetUser(
	ctx context.Context, req *userpb.GetUserRequest,
) (*userpb.GetUserResponse, error) {
	var userMessage *userpb.UserMessage
	user, _ := ugs.UserRepository.GetUser(req.UserId)
	userIdStr := strconv.FormatUint(uint64(user.ID), 10)
	userMessage = &userpb.UserMessage{
		UserId:   userIdStr,
		Nickname: user.Nickname,
	}
	return &userpb.GetUserResponse{UserMessage: userMessage}, nil
}

func (ugs UserGrpcServer) ListUsers(
	ctx context.Context, in *userpb.ListUsersRequest,
) (*userpb.ListUsersResponse, error) {
	var userMessages []*userpb.UserMessage
	users, _ := ugs.UserRepository.ListUsers()
	for _, user := range users {
		userIdStr := strconv.FormatUint(uint64(user.ID), 10)
		userMessages = append(userMessages, &userpb.UserMessage{
			UserId:   userIdStr,
			Nickname: user.Nickname,
		})
	}
	return &userpb.ListUsersResponse{
		UserMessages: userMessages,
	}, nil
}
