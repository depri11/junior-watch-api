package models

type CreateUser struct {
	UserID         string  `json:"userId" validate:"required"`
	Username       string  `json:"username" validate:"required"`
	Email          string  `json:"email" validate:"required"`
	Name           string  `json:"name" validate:"required"`
	Phone          string  `json:"Phone" validate:"required,gte=0,lte=15"`
	Role           int     `json:"roleId"`
	Address        string  `json:"address"`
	ProfilePicture float64 `json:"profilePicture"`
}

type CreateUserResponse struct {
	UserID string `json:"userId" validate:"required"`
}
