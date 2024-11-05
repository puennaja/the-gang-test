package mongorepo

import (
	"context"
	"daveslist/internal/core/domain/constant"
	"daveslist/internal/core/domain/dto"
	"daveslist/internal/core/domain/model"
	"daveslist/internal/core/port"
	errors "daveslist/pkg/go-errors"
	logger "daveslist/pkg/go-logger"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type listingRepository struct {
	repository[model.Listing]
}

func NewListingRepository(mc *mongo.Client, db string) port.ListingRepository {
	var repo = &listingRepository{repository[model.Listing]{
		mc:             mc,
		db:             mc.Database(db),
		col:            "listing",
		errInterceptor: listingErrorInterceptor,
	}}

	if err := repo.createIndex(); err != nil {
		logger.Fatal(err.Error())
	}

	return repo
}

func listingErrorInterceptor(err error) error {
	if err == nil {
		return nil
	}

	defultErr := errors.ErrMongo
	defultErr.SetError(err)
	logger.Error(err)
	return defultErr
}

func (repo *listingRepository) createIndex() error {
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
				"category_id": desc,
			},
			Options: options.Index().SetUnique(false),
		},
	})
	return err
}

func (repo *listingRepository) Insert(ctx context.Context, data *model.Listing) (*dto.ListingResponse, error) {
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

func (repo *listingRepository) filter(query *dto.ListingQuery) bson.M {
	out := bson.M{}
	if query.CategoryID != "" {
		out["category_id"] = query.CategoryID
	}
	if query.Level == constant.AuthLevel0 {
		out["is_private"] = false
		out["is_deleted"] = false
		out["is_hide"] = false
	} else if query.Level == constant.AuthLevel1 {
		out["is_hide"] = false
		out["is_deleted"] = false
	} else if query.Level == constant.AuthLevel2 {
		out["is_deleted"] = false
	}
	return out
}

func (repo *listingRepository) FindByQuery(ctx context.Context, query *dto.ListingQuery) (dto.ListingListResponse, error) {
	filter := repo.filter(query)
	opts := repo.paginationFindOptions(query.PaginationQuery)
	res, err := repo.find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	return model.ListingList(res).ToDTO(), nil
}

func (repo *listingRepository) CountByQuery(ctx context.Context, query *dto.ListingQuery) (int64, error) {
	filter := repo.filter(query)
	count, err := repo.countDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (repo *listingRepository) UpdateByQuery(ctx context.Context, query *dto.ListingQuery, update *model.UpdateListing) (int64, error) {
	filter := repo.filter(query)
	res, err := repo.updateMany(ctx, filter, update)
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (repo *listingRepository) FindOneByID(ctx context.Context, id string) (*dto.ListingResponse, error) {
	query := repo.buildQueryByID(id)
	res, err := repo.findOne(ctx, query, options.FindOne())
	if err != nil {
		return nil, err
	}
	return res.ToDTO(), nil
}

func (repo *listingRepository) UpdateOneByID(ctx context.Context, id string, update *model.UpdateListing) (*dto.ListingResponse, error) {
	query := repo.buildQueryByID(id)
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	res, err := repo.updateOne(ctx, query, update, opts)
	if err != nil {
		return nil, err
	}
	return res.ToDTO(), nil
}
