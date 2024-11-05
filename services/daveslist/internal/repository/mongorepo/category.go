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

type categoryRepository struct {
	repository[model.Category]
}

func NewCategoryRepository(mc *mongo.Client, db string) port.CategoryRepository {
	var repo = &categoryRepository{repository[model.Category]{
		mc:             mc,
		db:             mc.Database(db),
		col:            "category",
		errInterceptor: categoryErrorInterceptor,
	}}

	if err := repo.createIndex(); err != nil {
		logger.Fatal(err.Error())
	}

	return repo
}

func categoryErrorInterceptor(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, mongo.ErrNoDocuments) {
		return errors.ErrCategoryNotFound
	}

	if mongo.IsDuplicateKeyError(err) {
		return errors.ErrCategoryAlreadyExists
	}

	defultErr := errors.ErrMongo
	defultErr.SetError(err)
	return defultErr
}

func (repo *categoryRepository) createIndex() error {
	col := repo.collection()
	_, err := col.Indexes().CreateMany(context.Background(), []mongo.IndexModel{
		{
			Keys: bson.M{
				"id": desc,
			},
			Options: options.Index().SetUnique(true),
		},
	})
	return err
}

func (repo *categoryRepository) Insert(ctx context.Context, data *model.Category) (*dto.CategoryResponse, error) {
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

func (repo *categoryRepository) filter(query *dto.CategoryQuery) bson.M {
	out := bson.M{}
	if query.Level == constant.AuthLevel0 {
		out["is_private"] = false
		out["is_deleted"] = false
	} else if query.Level == constant.AuthLevel1 {
		out["is_deleted"] = false
	}
	return out
}

func (repo *categoryRepository) FindByQuery(ctx context.Context, query *dto.CategoryQuery) (dto.CategoryListResponse, error) {
	filter := repo.filter(query)
	opts := repo.paginationFindOptions(query.PaginationQuery)
	res, err := repo.find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	return model.CategoryList(res).ToDTO(), nil
}

func (repo *categoryRepository) CountByQuery(ctx context.Context, query *dto.CategoryQuery) (int64, error) {
	filter := repo.filter(query)
	count, err := repo.countDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (repo *categoryRepository) UpdateOneByID(ctx context.Context, id string, update *model.UpdateCategory) (*dto.CategoryResponse, error) {
	query := repo.buildQueryByID(id)
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	res, err := repo.updateOne(ctx, query, update, opts)
	if err != nil {
		return nil, err
	}
	return res.ToDTO(), nil
}
