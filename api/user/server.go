package user

import (
	"context"
	"fmt"
	"go-grpc-server/api/user/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type grpcServer struct {
	service Service
}

func ListenGRPC(s Service, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		return err
	}

	newServer := grpc.NewServer()
	pb.RegisterUserServiceServer(newServer, &grpcServer{s})
	reflection.Register(newServer)
	return newServer.Serve(lis)
}

func (s *grpcServer) GetUser(ctx context.Context, r *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	u, err := s.service.GetUser(ctx, r.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserResponse{
		User: &pb.User{
			Id:       u.ID,
			Username: u.Username,
			Email:    u.Email,
		},
	}, nil
}

func (s *grpcServer) PostUser(ctx context.Context, r *pb.PostUserRequest) (*pb.PostUserResponse, error) {
	u, err := s.service.PostUser(ctx, r.Username, r.Email)

	if err != nil {
		return nil, err
	}

	return &pb.PostUserResponse{
		User: &pb.User{
			Id:       u.ID,
			Username: u.Username,
			Email:    u.Email,
		},
	}, nil
}

func (s *grpcServer) UpdateUser(ctx context.Context, r *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	u, err := s.service.UpdateUser(ctx, r.Username, r.Email)

	if err != nil {
		return nil, err
	}

	return &pb.UpdateUserResponse{
		User: &pb.User{
			Id:       u.ID,
			Username: u.Username,
			Email:    u.Email,
		},
	}, nil
}

func (s *grpcServer) DeleteUser(ctx context.Context, r *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	err := s.service.DeleteUser(ctx, r.Id)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &pb.DeleteUserResponse{}, nil
}
