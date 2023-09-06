package user_test

import (
	"api-testing/user"
	mock_user "api-testing/user/mocks"
	"context"
	"errors"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestHandlerGet(t *testing.T) {
	router := gin.Default()
	mockCtrl := gomock.NewController(t)

	userID, _ := uuid.Parse("d571099d-b6c6-4a68-9273-b7f0f496f993")
	u := user.User{
		ID:       userID,
		Username: "username",
	}

	mockService := mock_user.NewMockService(mockCtrl)
	userHandler := user.NewHandler(mockService)

	var tests = []struct {
		expectedOutput byte
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
