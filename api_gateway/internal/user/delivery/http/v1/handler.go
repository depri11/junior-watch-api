package v1

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/depri11/junior-watch-api/api_gateway/config"
	"github.com/depri11/junior-watch-api/api_gateway/internal/models"
	"github.com/depri11/junior-watch-api/api_gateway/internal/user/service"
	httpErrors "github.com/depri11/junior-watch-api/pkg/http_errors"
	"github.com/depri11/junior-watch-api/pkg/logger"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type userHandlers struct {
	group       *mux.Router
	log         logger.Logger
	cfg         *config.Config
	v           *validator.Validate
	userService *service.UserService
}

func NewUserHandlers(
	group *mux.Router,
	log logger.Logger,
	cfg *config.Config,
	v *validator.Validate,
	userService *service.UserService,
) *userHandlers {
	return &userHandlers{group: group, log: log, cfg: cfg, v: v, userService: userService}
}

func (h *userHandlers) CreateUser(w http.ResponseWriter, r *http.Request) {
	var payload models.CreateUser
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		httpErrors.ErrorCtxResponse(w, r, err)
		return
	}

	payload.UserID = uuid.New()
	payload.RoleID = uuid.New()

	ctx := context.Background()

	if err := h.v.StructCtx(ctx, payload); err != nil {
		httpErrors.ErrorCtxResponse(w, r, err)
		return
	}

	res, err := h.userService.CreateUser(ctx, payload)
	if err != nil {
		httpErrors.ErrorCtxResponse(w, r, err)
		return
	}

	resBytes, err := json.Marshal(res)
	if err != nil {
		httpErrors.ErrorCtxResponse(w, r, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(resBytes)
}
