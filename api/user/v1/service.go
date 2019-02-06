package user

import (
	"context"
	"time"

	"github.com/segmentio/ksuid"
)

type Service interface {
	PostUser(ctx context.Context, username, email string) (*User, error)
	GetUser(ctx context.Context, id string) (*User, error)
}

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type userService struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &userService{r}
}

func (s *userService) PostUser(ctx context.Context, username, email string) (*User, error) {
	u := &User{
		ID:        ksuid.New().String(),
		Username:  username,
		Email:     email,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	if err := s.repository.PutUser(ctx, *u); err != nil {
		return nil, err
	}

	return u, nil
}

func (s *userService) GetUser(ctx context.Context, id string) (*User, error) {
	return s.repository.GetUser(ctx, id)
}
