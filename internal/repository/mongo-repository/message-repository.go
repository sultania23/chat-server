package mongo_repository

import (
	"context"
	. "github.com/google/uuid"
	"github.com/tuxoo/idler/internal/model/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MessageRepository struct {
	db *mongo.Collection
}

func NewMessageRepository(db *mongo.Database) *MessageRepository {
	return &MessageRepository{
		db: db.Collection(messageCollection),
	}
}

func (r *MessageRepository) Save(ctx context.Context, message entity.Message) error {
	_, err := r.db.InsertOne(ctx, message)
	return err
}

func (r *MessageRepository) SaveAll(ctx context.Context, messages []entity.Message) error {
	_, err := r.db.InsertMany(ctx, []any{})
	return err
}

func (r *MessageRepository) FindByConversationId(ctx context.Context, conversationId UUID) (entity.Message, error) {
	var message entity.Message
	err := r.db.FindOne(ctx, bson.M{"conversationId": conversationId.String()}).Decode(&message)
	return message, err
}

func (r *MessageRepository) FindAllByConversationId(ctx context.Context, conversationId UUID) ([]entity.Message, error) {
	var messages []entity.Message
	cur, err := r.db.Find(ctx, bson.M{"conversationId": conversationId.String()})
	if err != nil {
		return nil, err
	}

	err = cur.All(ctx, &messages)
	if err != nil {
		return nil, err
	}

	return messages, err
}
