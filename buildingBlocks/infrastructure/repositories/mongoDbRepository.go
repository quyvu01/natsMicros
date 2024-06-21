package repositories

import (
	"context"
	"errors"
	"github.com/ahmetb/go-linq"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"natsMicros/buildingBlocks/application/configurations"
	"natsMicros/buildingBlocks/application/responses"
)

type MongoDbRepository[TModel any] struct {
	Collections *mongo.Collection
}

func NewMongoDbRepository[TModel any](client *mongo.Client,
	setting *configurations.MongoDbSetting,
	collectionSetting *configurations.MongoDbCollectionSetting[TModel]) *MongoDbRepository[TModel] {
	return &MongoDbRepository[TModel]{Collections: client.Database(setting.DatabaseName).Collection(collectionSetting.CollectionName)}
}

func (repo *MongoDbRepository[TModel]) GetFirstByCondition(predict func(TModel) bool, ctx context.Context) (*TModel, error) {
	filter := GetFilter[TModel](predict)
	var result TModel
	err := repo.Collections.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
func (repo *MongoDbRepository[TModel]) ExistByCondition(predict func(TModel) bool, ctx context.Context) (bool, error) {
	filter := GetFilter[TModel](predict)
	var result TModel
	err := repo.Collections.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
func (repo *MongoDbRepository[TModel]) GetManyByCondition(predict func(TModel) bool, ctx context.Context) ([]*TModel, error) {
	var filter = bson.D{}
	if predict == nil {
		filter = bson.D{}
	}
	var result []*TModel
	cursor, err := repo.Collections.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		_ = cursor.Close(ctx)
	}(cursor, ctx)
	for cursor.Next(ctx) {
		var model TModel
		err := cursor.Decode(&model)
		if err != nil {
			return nil, err
		}
		result = append(result, &model)
	}
	return result, nil
}
func (repo *MongoDbRepository[TModel]) GetPaginationByCondition(predict func(TModel) bool, pageSize int64, pageIndex int64, ctx context.Context) (responses.PaginationResponse[*TModel], error) {
	var filter = bson.D{}
	if predict == nil {
		filter = bson.D{}
	}
	var result []*TModel
	cursor, err := repo.Collections.Find(ctx, filter)
	if err != nil {
		return responses.PaginationResponse[*TModel]{}, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		_ = cursor.Close(ctx)
	}(cursor, ctx)
	for cursor.Next(ctx) {
		var model TModel
		err := cursor.Decode(&model)
		if err != nil {
			return responses.PaginationResponse[*TModel]{}, err
		}
		result = append(result, &model)
	}
	totalItem, err := repo.CountByCondition(predict, ctx)
	var pSize int64 = 1
	if pageSize > 0 {
		pSize = pageSize
	}
	totalPage := (totalItem + pageSize - 1) / pSize
	return responses.PaginationResponse[*TModel]{Items: result, TotalRecord: totalItem, CurrentPageIndex: pageIndex, TotalPage: totalPage}, nil
}
func (repo *MongoDbRepository[TModel]) CountByCondition(predict func(TModel) bool, ctx context.Context) (int64, error) {
	filter := GetFilter[TModel](predict)
	counting, err := repo.Collections.
		CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}
	return counting, nil
}
func (repo *MongoDbRepository[TModel]) CreateOne(entity *TModel, ctx context.Context) (*TModel, error) {
	_, err := repo.Collections.InsertOne(ctx, entity)
	if err != nil {
		return nil, err
	}
	return entity, nil
}
func (repo *MongoDbRepository[TModel]) CreateMany(entities []*TModel, ctx context.Context) ([]*TModel, error) {
	models := make([]interface{}, len(entities))
	linq.From(entities).SelectT(func(model TModel) interface{} {
		return model
	}).ToSlice(&models)
	_, err := repo.Collections.InsertMany(ctx, models)
	if err != nil {
		return nil, err
	}
	return entities, nil
}
func (repo *MongoDbRepository[TModel]) RemoveOne(predict func(TModel) bool, ctx context.Context) error {
	filter := GetFilter[TModel](predict)
	_, err := repo.Collections.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
func (repo *MongoDbRepository[TModel]) RemoveMany(predict func(TModel) bool, ctx context.Context) error {
	filter := GetFilter[TModel](predict)
	_, err := repo.Collections.DeleteMany(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
func GetFilter[TModel any](predict func(TModel) bool) bson.D {
	var filter = bson.D{}
	if predict == nil {
		filter = bson.D{}
	}
	return filter
}
