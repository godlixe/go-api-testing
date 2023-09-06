package user_test

import (
	"api-testing/user"
	mock_user "api-testing/user/mocks"
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestRepositoryGetAll(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userID, _ := uuid.Parse("d571099d-b6c6-4a68-9273-b7f0f496f993")
	u := user.User{
		ID:       userID,
		Username: "username",
	}

	mockDB := mock_user.NewMockDB(mockCtrl)
	userRepository := user.NewRepository(mockDB)

	var tests = []struct {
		expectedOutput []user.User
		expectedError  error
	}{
		{expectedOutput: []user.User{u}, expectedError: nil},
		{expectedOutput: nil, expectedError: errors.New("failed getting users.")},
	}

	for _, tt := range tests {
		mockDB.EXPECT().GetAll().Return(tt.expectedOutput, tt.expectedError)
		res, err := userRepository.GetAll(context.Background())
		assert.Equal(t, tt.expectedOutput, res)
		assert.Equal(t, tt.expectedError, err)
	}
}

func TestRepositoryGet(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userID, _ := uuid.Parse("d571099d-b6c6-4a68-9273-b7f0f496f993")
	u := user.User{
		ID:       userID,
		Username: "username",
	}

	mockDB := mock_user.NewMockDB(mockCtrl)
	userRepository := user.NewRepository(mockDB)

	var tests = []struct {
		input          any
		expectedOutput user.User
		expectedError  error
	}{
		{input: userID, expectedOutput: u, expectedError: nil},
		{input: userID, expectedOutput: user.User{}, expectedError: errors.New("failed getting user.")},
	}

	for _, tt := range tests {
		mockDB.EXPECT().Get(tt.input).Return(tt.expectedOutput, tt.expectedError)
		res, err := userRepository.Get(context.Background(), userID)
		assert.Equal(t, tt.expectedOutput, res)
		assert.Equal(t, tt.expectedError, err)
	}
}

func TestRepositoryCreate(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userID, _ := uuid.Parse("d571099d-b6c6-4a68-9273-b7f0f496f993")
	u := user.User{
		ID:       userID,
		Username: "username",
	}

	mockDB := mock_user.NewMockDB(mockCtrl)
	userRepository := user.NewRepository(mockDB)

	var tests = []struct {
		input         user.User
		expectedError error
	}{
		{input: u, expectedError: nil},
		{input: u, expectedError: errors.New("failed getting users.")},
	}

	for _, tt := range tests {
		mockDB.EXPECT().Create(u).Return(tt.expectedError)
		err := userRepository.Create(context.Background(), tt.input)
		assert.Equal(t, tt.expectedError, err)
	}
}
