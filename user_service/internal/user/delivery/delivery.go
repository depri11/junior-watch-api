package delivery

import (
	"context"

	"github.com/depri11/junior-watch-api/pkg/logger"
	"github.com/depri11/junior-watch-api/user_service/internal/models"
	"github.com/depri11/junior-watch-api/user_service/internal/user/interfaces"
	userService "github.com/depri11/junior-watch-api/user_service/proto"
	"github.com/google/uuid"
)

type UserDelivery struct {
	userService interfaces.UserService
	logger      logger.Logger
}

func NewUserDelivery(userService interfaces.UserService, logger logger.Logger) *UserDelivery {
	return &UserDelivery{userService: userService, logger: logger}
}

func (u *UserDelivery) CreateUser(ctx context.Context, r *userService.CreateUserRequest) (*userService.CreateUserResponse, error) {

	user := &models.CreateUser{
		UserID:   uuid.NewString(),
		Name:     r.Name,
		Email:    r.Email,
		Username: r.Username,
		Address:  r.Address,
		Phone:    r.Phone,
		Role:     r.Role.Descriptor().Index(),
	}

	res, err := u.userService.Register(ctx, user)
	if err != nil {
		return nil, err
	}
	return &userService.CreateUserResponse{UserID: res.UserID}, nil
}
