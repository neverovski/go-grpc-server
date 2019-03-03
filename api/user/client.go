package user

import (
	"context"
	"go-grpc-server/api/user/pb"
	"google.golang.org/grpc"
)

type Client struct {
	conn    *grpc.ClientConn
	service pb.UserServiceClient
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := pb.NewUserServiceClient(conn)
	return &Client{conn, c}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) GetUser(ctx context.Context, id string) (*User, error) {
	r, err := c.service.GetUser(
		ctx,
		&pb.GetUserRequest{Id: id},
	)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:       r.User.Id,
		Username: r.User.Username,
		Email:    r.User.Email,
	}, nil
}

func (c *Client) PostUser(ctx context.Context, username, email string) (*User, error) {
	r, err := c.service.PostUser(
		ctx,
		&pb.PostUserRequest{Username: username, Email: email},
	)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:       r.User.Id,
		Username: r.User.Username,
		Email:    r.User.Email,
	}, nil
}

func (c *Client) UpdateUser(ctx context.Context, username, email string) (*User, error) {
	r, err := c.service.UpdateUser(
		ctx,
		&pb.UpdateUserRequest{Username: username, Email: email},
	)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       r.User.Id,
		Username: r.User.Username,
		Email:    r.User.Email,
	}, nil
}

func (c *Client) DeleteUser(ctx context.Context, id string) error {
	_, err := c.service.DeleteUser(
		ctx,
		&pb.DeleteUserRequest{Id: id},
	)

	return err
}
