package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/brianvoe/gofakeit"
	"github.com/vashkevichjr/auth/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const grpcPort = "50051"

type server struct {
	user_v1.UnimplementedUserServiceServer
}

func (s *server) CreateUser(ctx context.Context, req *user_v1.CreateUserRequest) (*user_v1.CreateUserResponse, error) {
	return &user_v1.CreateUserResponse{Id: int64(gofakeit.Uint64())}, nil
}

func (s *server) GetUser(ctx context.Context, req *user_v1.GetUserRequest) (*user_v1.GetUserResponse, error) {

	return &user_v1.GetUserResponse{
		Id:        req.Id,
		Name:      gofakeit.Name(),
		Email:     gofakeit.Email(),
		Role:      0,
		CreatedAt: timestamppb.New(gofakeit.Date()),
		UpdatedAt: timestamppb.New(gofakeit.Date()),
	}, nil
}

func (s *server) UpdateUser(ctx context.Context, req *user_v1.UpdateUserRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *server) DeleteUser(ctx context.Context, req *user_v1.DeleteUserRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	user_v1.RegisterUserServiceServer(s, &server{})

	log.Printf("gRPC server listening on port %s", grpcPort)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
