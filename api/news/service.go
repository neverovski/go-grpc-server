package news

import (
	"context"
	"github.com/segmentio/ksuid"
	"time"
)

type Service interface {
	PostNews(ctx context.Context, title, description, h1, text, userID string, published bool) (*News, error)
	GetNews(ctx context.Context, id string) (*News, error)
	GetNewsForUser(ctx context.Context, userID string) ([]News, error)
}

type News struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	H1          string    `json:"h1"`
	Text        string    `json:"text"`
	UserID      string    `json:"userID"`
	Published   bool      `json:"published"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (s *newsService) PostNews(ctx context.Context, title, description, h1, text, userID string, published bool) (*News, error) {
	n := &News{
		ID:          ksuid.New().String(),
		Title:       title,
		Description: description,
		H1:          h1,
		Text:        text,
		UserID:      userID,
		Published:   published,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}

	return n, nil
}

func (s *newsService) GetNews(ctx context.Context, id string) (*News, error) {
	return n, nil
}

func (s newsService) GetNewsForUser(ctx context.Context, userID string) ([]News, error) {
	return n, nil
}
