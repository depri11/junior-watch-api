package v1

import (
	"context"
	"net/http"

	"github.com/depri11/junior-watch-api/api_gateway/config"
	"github.com/depri11/junior-watch-api/api_gateway/internal/models"
	"github.com/depri11/junior-watch-api/api_gateway/internal/user/service"
	"github.com/depri11/junior-watch-api/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

type userHandlers struct {
	group       *gin.RouterGroup
	log         logger.Logger
	cfg         *config.Config
	v           *validator.Validate
	userService *service.UserService
}

func NewUserHandlers(
	group *gin.RouterGroup,
	log logger.Logger,
	cfg *config.Config,
	v *validator.Validate,
	userService *service.UserService,
) *userHandlers {
	return &userHandlers{group: group, log: log, cfg: cfg, v: v, userService: userService}
}

func (h *userHandlers) CreateUser(c *gin.Context) {
	var payload models.CreateUser
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	payload.UserID = uuid.New()
	payload.RoleID = uuid.New()

	ctx := context.Background()

	if err := h.v.StructCtx(ctx, payload); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := h.userService.CreateUser(ctx, payload)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}
