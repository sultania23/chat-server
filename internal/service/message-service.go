package service

import (
	"context"
	. "github.com/google/uuid"
	"github.com/tuxoo/idler/internal/model/entity"
	"github.com/tuxoo/idler/internal/repository/mongo-repository"
)

type MessageService struct {
	repository mongo_repository.Messages
}

func NewMessageService(repository mongo_repository.Messages) *MessageService {
	return &MessageService{
		repository: repository,
	}
}

func (s *MessageService) Create(ctx context.Context, message entity.Message) error {
	return s.repository.Save(ctx, message)
}

func (s *MessageService) CreateAll(ctx context.Context, messages []entity.Message) error {
	return s.repository.SaveAll(ctx, messages)
}

func (s *MessageService) GetByConversationId(ctx context.Context, id UUID) (entity.Message, error) {
	return s.repository.FindByConversationId(ctx, id)
}

func (s *MessageService) GetAllConversationId(ctx context.Context, id UUID) ([]entity.Message, error) {
	return s.repository.FindAllByConversationId(ctx, id)
}
