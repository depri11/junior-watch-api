package dto

import uuid "github.com/satori/go.uuid"

type CreateUserDto struct {
	ProductID   uuid.UUID `json:"productId" validate:"required"`
	Name        string    `json:"name" validate:"required,gte=0,lte=255"`
	Description string    `json:"description" validate:"required,gte=0,lte=5000"`
	Price       float64   `json:"price" validate:"required,gte=0"`
}

type CreateUserResponseDto struct {
	ProductID uuid.UUID `json:"productId" validate:"required"`
}
