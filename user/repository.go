package user

import (
	"context"

	"github.com/google/uuid"
)

//go:generate mockgen -source=./repository.go -destination=./mocks/repository_test.go

type DB interface {
	GetAll() ([]User, error)
	Get(id uuid.UUID) (User, error)
	Create(user User) error
}

type repository struct {
	db DB
}

func NewRepository(db DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]User, error) {
	return r.db.GetAll()
}

func (r *repository) Get(ctx context.Context, id uuid.UUID) (User, error) {
	return r.db.Get(id)
}

func (r *repository) Create(ctx context.Context, user User) error {
	return r.db.Create(user)
}
