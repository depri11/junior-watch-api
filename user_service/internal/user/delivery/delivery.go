package delivery

import (
	"context"

	"github.com/depri11/junior-watch-api/pkg/logger"
	go_proto "github.com/depri11/junior-watch-api/pkg/proto"
	"github.com/depri11/junior-watch-api/user_service/internal/models"
	"github.com/depri11/junior-watch-api/user_service/internal/user/interfaces"
	"github.com/google/uuid"
)

type UserDelivery struct {
	userService interfaces.UserService
	logger      logger.Logger
}

func NewUserDelivery(userService interfaces.UserService, logger logger.Logger) *UserDelivery {
	return &UserDelivery{userService: userService, logger: logger}
}

func (u *UserDelivery) CreateUser(ctx context.Context, r *go_proto.CreateUserRequest) (*go_proto.CreateUserResponse, error) {

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
	return &go_proto.CreateUserResponse{UserID: res.UserID}, nil
}
