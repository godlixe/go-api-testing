package user

import (
	"api-testing/response"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=./handler.go -destination=./mocks/handler_mock.go

type Service interface {
	GetAll(ctx context.Context) ([]User, error)
	Get(ctx context.Context, id string) (User, error)
	Create(ctx context.Context, user User) error
}

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) GetAll(ctx *gin.Context) {
	res, err := h.service.GetAll(
		ctx.Request.Context(),
	)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.Response{
				Message: err.Error(),
				Data:    nil,
			},
		)

		return
	}

	ctx.JSON(
		http.StatusOK,

		response.Response{
			Message: "fetched users successfully.",
			Data:    res,
		},
	)
}

func (h *handler) Get(ctx *gin.Context) {
	userID := ctx.Param("id")

	res, err := h.service.Get(
		ctx.Request.Context(),
		userID,
	)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.Response{
				Message: err.Error(),
				Data:    nil,
			},
		)

		return
	}

	ctx.JSON(
		http.StatusOK,

		response.Response{
			Message: "fetched user successfully.",
			Data:    res,
		},
	)
}

func (h *handler) Create(ctx *gin.Context) {
	var userDTO UserRegister
	var err error

	err = ctx.ShouldBind(&userDTO)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			err,
		)

		return
	}

	fmt.Println(userDTO)

	err = h.service.Create(
		ctx.Request.Context(),
		User{
			Username: userDTO.Username,
			Password: userDTO.Password,
		},
	)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.Response{
				Message: err.Error(),
				Data:    nil,
			},
		)

		return
	}

	ctx.JSON(
		http.StatusOK,
		response.Response{
			Message: "created user successfully.",
			Data:    nil,
		},
	)
}
