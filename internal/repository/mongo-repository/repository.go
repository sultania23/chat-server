package mongo_repository

import (
	"context"
	. "github.com/google/uuid"
	"github.com/tuxoo/idler/internal/model/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	messageCollection = "message"
)

type Messages interface {
	Save(ctx context.Context, message entity.Message) error
	SaveAll(ctx context.Context, messages []entity.Message) error
	FindByConversationId(ctx context.Context, conversationId UUID) (entity.Message, error)
	FindAllByConversationId(ctx context.Context, conversationId UUID) ([]entity.Message, error)
}

type Repositories struct {
	Messages Messages
}

func NewRepositories(db *mongo.Database) *Repositories {
	return &Repositories{
		Messages: NewMessageRepository(db),
	}
}
