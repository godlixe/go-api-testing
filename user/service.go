package user

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	GetAll(ctx context.Context) ([]User, error)
	Get(ctx context.Context, id uuid.UUID) (User, error)
	Create(ctx context.Context, user User) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) *service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll(ctx context.Context) ([]User, error) {
	res, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *service) Get(ctx context.Context, strID string) (User, error) {
	id, err := uuid.Parse(strID)
	if err != nil {
		return User{}, err
	}

	res, err := s.repository.Get(ctx, id)
	if err != nil {
		return User{}, err
	}

	return res, nil
}

func (s *service) Create(ctx context.Context, user User) error {
	err := s.repository.Create(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
