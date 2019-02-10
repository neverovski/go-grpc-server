package news

import (
	"context"
	"time"

	"github.com/segmentio/ksuid"
)

type Service interface {
	GetNews(ctx context.Context, id string) (*News, error)
	GetNewsForUser(ctx context.Context, userID string) ([]News, error)
	PostNews(ctx context.Context, title, description, h1, text string, published bool) (*News, error)
	UpdateNews(ctx context.Context, title, description, h1, text string, published bool) (*News, error)
	DeleteNews(ctx context.Context, id string) error
}

type News struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	H1          string    `json:"h1"`
	Text        string    `json:"text"`
	UserID      string    `json:"user_id"`
	Published   bool      `json:"published"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type newsService struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &newsService{r}
}

func (s *newsService) GetNews(ctx context.Context, id string) (*News, error) {
	return s.repository.GetNews(ctx, id)
}

func (s *newsService) GetNewsForUser(ctx context.Context, userID string) ([]News, error) {
	return s.repository.GetNewsForUser(ctx, userID)
}

func (s *newsService) PostNews(ctx context.Context, title, description, h1, text string, published bool) (*News, error) {
	//TODO: Realize to user ID
	n := &News{
		ID:          ksuid.New().String(),
		Title:       title,
		Description: description,
		H1:          h1,
		Text:        text,
		UserID:      "Test us",
		Published:   published,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}

	if err := s.repository.PostNews(ctx, *n); err != nil {
		return nil, err
	}

	return n, nil
}

func (s *newsService) UpdateNews(ctx context.Context, title, description, h1, text string, published bool) (*News, error) {
	n := &News{
		ID:          ksuid.New().String(),
		Title:       title,
		Description: description,
		H1:          h1,
		Text:        text,
		UserID:      "Test us",
		Published:   published,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}

	if err := s.repository.PutNews(ctx, *n); err != nil {
		return nil, err
	}

	return n, nil
}

func (s *newsService) DeleteNews(ctx context.Context, id string) error {
	return s.repository.DeleteNews(ctx, id)
}
