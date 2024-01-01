package repository

import (
	"context"
	"log"

	"github.com/depri11/junior-watch-api/pkg/logger"
	userService "github.com/depri11/junior-watch-api/user_service/proto"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	log logger.Logger
	db  *sqlx.DB
}

func NewUserRepository(log logger.Logger, db *sqlx.DB) *UserRepository {
	return &UserRepository{log, db}
}

func (r *UserRepository) SaveUser(ctx context.Context, user *userService.CreateUserRequest) (string, error) {
	query := `INSERT INTO users (username, email, role_id, name, address) VALUES ($1, $2, $3, $4, $5) returning id`

	res, err := r.db.ExecContext(ctx, query, user.Username, user.Email, user.RoleID, user.Name, user.Address)
	if err != nil {
		return "", err
	}
	log.Println(res)
	return "", nil
}
