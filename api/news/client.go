package news

import (
	"context"
	"go-grpc-server/api/news/pb"
	"google.golang.org/grpc"
)

type Client struct {
	conn    *grpc.ClientConn
	service pb.NewsServiceClient
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := pb.NewNewsServiceClient(conn)
	return &Client{conn, c}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) GetNews(ctx context.Context, id string) (*News, error) {
	r, err := c.service.GetNews(
		ctx,
		&pb.GetNewsRequest{Id: id},
	)

	if err != nil {
		return nil, err
	}

	return &News{
		ID:          r.News.Id,
		Title:       r.News.Title,
		Description: r.News.Description,
		H1:          r.News.H1,
		Text:        r.News.Text,
		Published:   r.News.Published,
		UserID:      r.News.UserId,
	}, nil
}
