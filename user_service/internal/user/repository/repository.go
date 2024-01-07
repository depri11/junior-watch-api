package repository

import (
	"context"

	"github.com/depri11/junior-watch-api/pkg/logger"
	"github.com/depri11/junior-watch-api/user_service/internal/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	log *logger.Logger
	db  *sqlx.DB
}

func NewUserRepository(log *logger.Logger, db *sqlx.DB) *UserRepository {
	return &UserRepository{log, db}
}

func (r *UserRepository) SaveUser(ctx context.Context, user *models.CreateUser) (string, error) {
	query := `INSERT INTO users (id, username, email, role, name, address) VALUES ($1, $2, $3, $4, $5, $6) returning id`
	var id string
	err := r.db.QueryRowContext(ctx, query, uuid.NewString(), user.Username, user.Email, user.Role, user.Name, user.Address).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (r *UserRepository) FindByID(ctx context.Context, id string) (*models.User, error) {
	query := `
	SELECT
		id,
		username,
		email,
		role,
		name,
		address,
		phone,
		profile_picture,
		created_at,
		created_by,
		updated_at,
		updated_by,
		deleted,by
	FROM
		users
	WHERE
		id = $1
	`
	var user models.User
	err := r.db.QueryRowContext(ctx, query, id).Scan(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
