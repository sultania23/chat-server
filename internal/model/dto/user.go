package dto

import (
	. "github.com/google/uuid"
	"time"
)

type UserDTO struct {
	Id           UUID      `json:"id"`
	Name         string    `json:"name"`
	LoginEmail   string    `json:"email" db:"login_email"`
	RegisteredAt time.Time `json:"registeredAt" db:"registered_at"`
	VisitedAt    time.Time `json:"lastVisitAt" db:"visited_at"`
	Role         string    `json:"role"`
	IsEnabled    bool      `json:"enabled" db:"is_enabled"`
}

type SignInDTO struct {
	Email    string `json:"email" binding:"required,email,max=64" example:"kill-77@mail.ru"`
	Password string `json:"password" binding:"required,min=6,max=64" example:"qwerty"`
}

type VerifyDTO struct {
	Email     string `json:"email" binding:"required,email,max=64" example:"kill-77@mail.ru"`
	CheckCode string `json:"checkCode" binding:"" example:"e0b3073a05c0ed4920787a4f1574ff0066f7521e"`
}

type SignUpDTO struct {
	Name     string `json:"name" binding:"required,min=2,max=64" example:"alex"`
	Email    string `json:"email" binding:"required,email,max=64" example:"kill-77@mail.ru"`
	Password string `json:"password" binding:"required,min=6,max=64" example:"qwerty"`
}
