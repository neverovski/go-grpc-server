package news

import (
	"context"
	"fmt"
	"go-grpc-server/api/news/pb"
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
	serv := grpc.NewServer()
	pb.RegisterNewsServiceServer(serv, &grpcServer{s})
	reflection.Register(serv)
	return serv.Serve(lis)
}

func (s *grpcServer) GetNews(ctx context.Context, r *pb.GetNewsRequest) (*pb.GetNewsResponse, error) {
	n, err := s.service.GetNews(ctx, r.Id)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.GetNewsResponse{
		News: &pb.News{
			Id:          n.ID,
			H1:          n.H1,
			Title:       n.Title,
			Text:        n.Text,
			Description: n.Description,
			UserId:      n.UserID,
			Published:   n.Published,
		},
	}, nil
}

func (s *grpcServer) GetNewsForUser(ctx context.Context, r *pb.GetNewsForUserRequest) (*pb.GetNewsForUserResponse, error) {
	userNews, err := s.service.GetNewsForUser(ctx, r.UserId)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	news := []*pb.News{}

	return &pb.GetNewsForUserResponse{News: news}, nil
}

func (s *grpcServer) PostNews(ctx context.Context, r *pb.PostNewsRequest) (*pb.PostNewsResponse, error) {
	n, err := s.service.PostNews(ctx, r.Title, r.Description, r.H1, r.Text, r.Published)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.PostNewsResponse{
		News: &pb.News{
			Id:          n.ID,
			H1:          n.H1,
			Title:       n.Title,
			Text:        n.Text,
			Description: n.Description,
			UserId:      n.UserID,
			Published:   n.Published,
		},
	}, nil
}

func (s *grpcServer) UpdateNews(ctx context.Context, r *pb.UpdateNewsRequest) (*pb.UpdateNewsResponse, error) {
	n, err := s.service.UpdateNews(ctx, r.Title, r.Description, r.H1, r.Text, r.Published)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.UpdateNewsResponse{
		News: &pb.News{
			Id:          n.ID,
			H1:          n.H1,
			Title:       n.Title,
			Text:        n.Text,
			Description: n.Description,
			UserId:      n.UserID,
			Published:   n.Published,
		},
	}, nil
}

func (s *grpcServer) DeleteNews(ctx context.Context, r *pb.DeleteNewsRequest) (*pb.DeleteNewsResponse, error) {
	err := s.service.DeleteNews(ctx, r.Id)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.DeleteNewsResponse{}, nil
}
