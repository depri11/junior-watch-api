package dto

import uuid "github.com/satori/go.uuid"

type CreateUserDto struct {
	UserID         uuid.UUID `json:"userId" validate:"required"`
	Username       uuid.UUID `json:"username" validate:"required"`
	Email          uuid.UUID `json:"email" validate:"required"`
	Name           uuid.UUID `json:"name" validate:"required"`
	Phone          string    `json:"Phone" validate:"required,gte=0,lte=255"`
	RoleID         uuid.UUID `json:"roleId" validate:"required"`
	Address        string    `json:"address"`
	ProfilePicture float64   `json:"profilePicture"`
}

type CreateUserResponseDto struct {
	UserID uuid.UUID `json:"userId" validate:"required"`
}
