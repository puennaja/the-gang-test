package mongorepo

import (
	"context"
	"daveslist/internal/core/domain/dto"
	"daveslist/internal/core/domain/model"
	"daveslist/internal/core/port"
	errors "daveslist/pkg/go-errors"
	logger "daveslist/pkg/go-logger"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type messageRepository struct {
	repository[model.Message]
}

func NewMessageRepository(mc *mongo.Client, db string) port.MessageRepository {
	var repo = &messageRepository{repository[model.Message]{
		mc:             mc,
		db:             mc.Database(db),
		col:            "message",
		errInterceptor: messageErrorInterceptor,
	}}

	if err := repo.createIndex(); err != nil {
		logger.Fatal(err.Error())
	}

	return repo
}

func messageErrorInterceptor(err error) error {
	if err == nil {
		return nil
	}

	defultErr := errors.ErrMongo
	defultErr.SetError(err)
	return defultErr
}

func (repo *messageRepository) createIndex() error {
	col := repo.collection()
	_, err := col.Indexes().CreateMany(context.Background(), []mongo.IndexModel{
		{
			Keys: bson.M{
				"id": desc,
			},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bson.M{
				"sender_id": desc,
			},
			Options: options.Index().SetUnique(false),
		},
	})
	return err
}

func (repo *messageRepository) Insert(ctx context.Context, data *model.Message) (*dto.MessageResponse, error) {
	oid, err := repo.insertOne(ctx, data)
	if err != nil {
		return nil, err
	}

	query := repo.buildQueryByID(oid.Hex())
	res, err := repo.findOne(ctx, query, options.FindOne())
	if err != nil {
		return nil, err
	}
	return res.ToDTO(), nil
}

func (repo *messageRepository) filter(query *dto.MessageQuery) bson.M {
	out := bson.M{}
	if query.SenderID != "" {
		out["sender_id"] = query.SenderID
	}
	return out
}

func (repo *messageRepository) FindByQuery(ctx context.Context, query *dto.MessageQuery) (dto.MessageListResponse, error) {
	filter := repo.filter(query)
	opts := repo.paginationFindOptions(query.PaginationQuery)
	res, err := repo.find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	return model.MessageList(res).ToDTO(), nil
}
func (repo *messageRepository) CountByQuery(ctx context.Context, query *dto.MessageQuery) (int64, error) {
	filter := repo.filter(query)
	count, err := repo.countDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}
	return count, nil
}
