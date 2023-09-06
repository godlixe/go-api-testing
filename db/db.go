package db

import (
	"api-testing/user"
	"errors"

	"github.com/google/uuid"
)

type db struct {
	data map[uuid.UUID]user.User
}

func New() *db {
	mp := make(map[uuid.UUID]user.User)

	return &db{
		data: mp,
	}
}

func (d *db) GetAll() ([]user.User, error) {
	var res []user.User

	for _, data := range d.data {
		res = append(res, data)
	}

	return res, nil
}

func (d *db) Get(id uuid.UUID) (user.User, error) {
	if v, ok := d.data[id]; ok {
		return v, nil
	}

	return user.User{}, errors.New("user not found.")
}

func (d *db) Create(user user.User) error {
	id := uuid.New()

	user.ID = id

	d.data[id] = user

	return nil
}
