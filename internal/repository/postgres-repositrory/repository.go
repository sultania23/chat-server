package postgres_repository

import (
	"context"
	. "github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/tuxoo/idler/internal/model/dto"
	"github.com/tuxoo/idler/internal/model/entity"
)

const (
	userTable         = "\"user\""
	conversationTable = "conversation"
)

type Users interface {
	Save(ctx context.Context, user entity.User) (*dto.UserDTO, error)
	UpdateById(ctx context.Context, id UUID) error
	FindByCredentials(ctx context.Context, email, password string) (*dto.UserDTO, error)
	FindById(ctx context.Context, id UUID) (*dto.UserDTO, error)
	FindAll(ctx context.Context) ([]dto.UserDTO, error)
	FindByEmail(ctx context.Context, email string, isEnabled bool) (*dto.UserDTO, error)
}

type Conversations interface {
	Save(ctx context.Context, conversation entity.Conversation) (*dto.ConversationDTO, error)
	FindByOwnerId(ctx context.Context, id UUID) ([]dto.ConversationDTO, error)
	FindById(ctx context.Context, id UUID) (*dto.ConversationDTO, error)
	DeleteById(ctx context.Context, id UUID) error
}

type Repositories struct {
	Users         Users
	Conversations Conversations
}

func NewRepositories(db *pgxpool.Pool) *Repositories {
	//func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		Users:         NewUserRepository(db),
		Conversations: NewConversationRepository(db),
	}
}
