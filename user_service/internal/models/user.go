package models

import uuid "github.com/google/uuid"

type CreateUser struct {
	UserID         uuid.UUID `json:"userId" validate:"required"`
	Username       string    `json:"username" validate:"required"`
	Email          string    `json:"email" validate:"required"`
	Name           string    `json:"name" validate:"required"`
	Phone          string    `json:"Phone" validate:"required,gte=0,lte=255"`
	RoleID         uuid.UUID `json:"roleId" validate:"required"`
	Address        string    `json:"address"`
	ProfilePicture float64   `json:"profilePicture"`
}

type CreateUserResponse struct {
	UserID uuid.UUID `json:"userId" validate:"required"`
}
