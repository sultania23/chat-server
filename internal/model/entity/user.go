package entity

import "time"
import . "github.com/google/uuid"

const (
	UserRole   = "USER"
	AdminNRole = "ADMIN"
)

type User struct {
	Id           UUID      `json:"-" db:"id"`
	Name         string    `json:"name" db:"name" binding:"required"`
	LoginEmail   string    `json:"email" db:"login_email" binding:"required"`
	PasswordHash string    `json:"-" db:"-" binding:"required"`
	RegisteredAt time.Time `json:"registeredAt" db:"registered_at"`
	VisitedAt    time.Time `json:"lastVisitAt" db:"visited_at"`
	Role         string    `json:"role" db:"role"`
	IsEnabled    bool      `json:"-" db:"is_enabled"`
}
