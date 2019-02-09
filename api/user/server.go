package user

import (
	"context"
	"fmt"
	"go-grpc-server/api/user/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	serv := grpc.NewServer()
	pb.RegisterUserServiceServer(serv, &grpcServer{s})
	reflection.Register(serv)
	return serv.Serve(lis)
}

func (s *grpcServer) PostUser(ctx context.Context, r *pb.PostUserRequest) (*pb.PostUserResponse, error) {
	u, err := s.service.PostUser(ctx, r.User)
	if err != nil {
		return nil, err
	}
	return &pb.PostUserResponse{User: &pb.User{
		Id:       u.ID,
		Username: u.Username,
		Email:    u.Email,
	}}, nil
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

func (s *grpcServer) UpdateUser(ctx context.Context, r *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {

}

func (s *grpcServer) DeleteUser(ctx context.Context, r *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {

}
