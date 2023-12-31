package v1

import (
	"context"
	"net/http"

	"github.com/depri11/junior-watch-api/api_gateway/config"
	"github.com/depri11/junior-watch-api/api_gateway/internal/dto"
	"github.com/depri11/junior-watch-api/api_gateway/internal/user/commands"
	"github.com/depri11/junior-watch-api/api_gateway/internal/user/service"
	"github.com/depri11/junior-watch-api/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	uuid "github.com/satori/go.uuid"
)

type userHandlers struct {
	group *gin.RouterGroup
	log   logger.Logger
	cfg   *config.Config
	ps    *service.UserService
	v     *validator.Validate
}

func NewUserHandlers(
	group *gin.RouterGroup,
	log logger.Logger,
	cfg *config.Config,
	ps *service.UserService,
	v *validator.Validate,
) *userHandlers {
	return &userHandlers{group: group, log: log, cfg: cfg, ps: ps, v: v}
}

func (h *userHandlers) CreateUser(c *gin.Context) {
	var payload dto.CreateUserDto
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	payload.UserID = uuid.NewV4()
	payload.RoleID = uuid.NewV4()

	ctx := context.Background()

	if err := h.v.StructCtx(ctx, payload); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.ps.Commands.CreateUser.Handle(ctx, commands.NewCreateUserCommand(&payload)); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.CreateUserResponseDto{UserID: payload.UserID})
}
