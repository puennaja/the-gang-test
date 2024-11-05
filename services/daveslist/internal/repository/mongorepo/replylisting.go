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

type replyListingRepository struct {
	repository[model.ReplyListing]
}

func NewReplyListingRepository(mc *mongo.Client, db string) port.ReplyListingRepository {
	var repo = &replyListingRepository{repository[model.ReplyListing]{
		mc:             mc,
		db:             mc.Database(db),
		col:            "replyListing",
		errInterceptor: replyListingErrorInterceptor,
	}}

	if err := repo.createIndex(); err != nil {
		logger.Fatal(err.Error())
	}

	return repo
}

func replyListingErrorInterceptor(err error) error {
	if err == nil {
		return nil
	}

	defultErr := errors.ErrMongo
	defultErr.SetError(err)
	return defultErr
}

func (repo *replyListingRepository) createIndex() error {
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
				"listing_id": desc,
			},
			Options: options.Index().SetUnique(false),
		},
	})
	return err
}

func (repo *replyListingRepository) Insert(ctx context.Context, data *model.ReplyListing) (*dto.ReplyListingResponse, error) {
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

func (repo *replyListingRepository) filter(query *dto.ReplyListingQuery) bson.M {
	out := bson.M{}
	if query.ListingID != "" {
		out["listing_id"] = query.ListingID
	}
	return out
}

func (repo *replyListingRepository) FindByQuery(ctx context.Context, query *dto.ReplyListingQuery) (dto.ReplyListingListResponse, error) {
	filter := repo.filter(query)
	opts := repo.paginationFindOptions(query.PaginationQuery)
	res, err := repo.find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	return model.ReplyListingList(res).ToDTO(), nil
}
func (repo *replyListingRepository) CountByQuery(ctx context.Context, query *dto.ReplyListingQuery) (int64, error) {
	filter := repo.filter(query)
	count, err := repo.countDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}
	return count, nil
}
