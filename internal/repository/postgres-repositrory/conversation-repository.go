package postgres_repository

import (
	"context"
	"fmt"
	. "github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/tuxoo/idler/internal/model/dto"
	"github.com/tuxoo/idler/internal/model/entity"
)

type ConversationRepository struct {
	db *pgxpool.Pool
}

func NewConversationRepository(db *pgxpool.Pool) *ConversationRepository {
	return &ConversationRepository{db: db}
}

func (r *ConversationRepository) Save(ctx context.Context, conversation entity.Conversation) (*dto.ConversationDTO, error) {
	var newConversation dto.ConversationDTO

	query := fmt.Sprintf("INSERT INTO %s (name, owner) VALUES ($1, $2) RETURNING name, owner", conversationTable)
	row := r.db.QueryRow(ctx, query, conversation.Name, conversation.Owner)
	if err := row.Scan(&newConversation.Name, &newConversation.Owner); err != nil {
		return &newConversation, err
	}

	return &newConversation, nil
}

func (r *ConversationRepository) FindByOwnerId(ctx context.Context, id UUID) ([]dto.ConversationDTO, error) {
	var conversations []dto.ConversationDTO
	//err := r.db.Select(&conversations, query, id)
	//if err != nil {
	//	return conversations, err
	//}

	return conversations, nil
}

func (r *ConversationRepository) FindById(ctx context.Context, id UUID) (*dto.ConversationDTO, error) {
	var conversation dto.ConversationDTO
	query := fmt.Sprintf("SELECT name, owner FROM %s WHERE id=$1", conversationTable)

	row := r.db.QueryRow(ctx, query, id)
	if err := row.Scan(&conversation.Name, &conversation.Owner); err != nil { // TODO: Perhaps need convert id to string
		return &conversation, err
	}

	return &conversation, nil
}

func (r *ConversationRepository) DeleteById(ctx context.Context, id UUID) error {
	query := fmt.Sprintf("DELETE FROM %s where id=$1", conversationTable)
	_, err := r.db.Exec(ctx, query, id)
	return err
}
