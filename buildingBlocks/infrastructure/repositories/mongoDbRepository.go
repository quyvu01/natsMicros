package repositories

import (
	"context"
	"errors"
	"github.com/thoas/go-funk"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"natsMicros/buildingBlocks/application/configurations"
	"natsMicros/buildingBlocks/application/responses"
)

type MongoDbRepository[TModel any] struct {
	// Define the MongoDb Driver here
	Database    string
	Collection  string
	MongoClient *mongo.Client
}

func NewMongoDbRepository[TModel any](client *mongo.Client, setting *configurations.MongoDbSetting, collectionSetting *configurations.MongoDbCollectionSetting[TModel]) *MongoDbRepository[TModel] {
	return &MongoDbRepository[TModel]{MongoClient: client, Database: setting.DatabaseName, Collection: collectionSetting.CollectionName}
}

func (repo *MongoDbRepository[TModel]) GetFirstByCondition(expression func(TModel) bool) (*TModel, error) {
	filter := GetFilter[TModel](expression)
	var result TModel
	err := repo.MongoClient.Database(repo.Database).Collection(repo.Collection).FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
func (repo *MongoDbRepository[TModel]) ExistByCondition(expression func(TModel) bool) (bool, error) {
	filter := GetFilter[TModel](expression)
	var result TModel
	err := repo.MongoClient.Database(repo.Database).Collection(repo.Collection).FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (repo *MongoDbRepository[TModel]) GetManyByCondition(expression func(TModel) bool) ([]*TModel, error) {
	var filter = bson.D{}
	if expression == nil {
		filter = bson.D{}
	}
	var result []*TModel
	cursor, err := repo.MongoClient.Database(repo.Database).Collection(repo.Collection).Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		_ = cursor.Close(ctx)
	}(cursor, context.TODO())
	for cursor.Next(context.TODO()) {
		var model TModel
		err := cursor.Decode(&model)
		if err != nil {
			return nil, err
		}
		result = append(result, &model)
	}
	return result, nil
}

func (repo *MongoDbRepository[TModel]) GetPaginationByCondition(expression func(TModel) bool) (responses.PaginationResponse[*TModel], error) {
	var filter = bson.D{}
	if expression == nil {
		filter = bson.D{}
	}
	var result []*TModel
	cursor, err := repo.MongoClient.Database(repo.Database).Collection(repo.Collection).Find(context.TODO(), filter)
	if err != nil {
		return responses.PaginationResponse[*TModel]{}, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		_ = cursor.Close(ctx)
	}(cursor, context.TODO())
	for cursor.Next(context.TODO()) {
		var model TModel
		err := cursor.Decode(&model)
		if err != nil {
			return responses.PaginationResponse[*TModel]{}, err
		}
		result = append(result, &model)
	}
	totalItem, err := repo.CountByCondition(expression)
	return responses.PaginationResponse[*TModel]{Items: result, TotalRecord: totalItem}, nil
}

func (repo *MongoDbRepository[TModel]) CountByCondition(expression func(TModel) bool) (int64, error) {
	filter := GetFilter[TModel](expression)
	counting, err := repo.MongoClient.Database(repo.Database).Collection(repo.Collection).CountDocuments(context.TODO(), filter)
	if err != nil {
		return 0, err
	}
	return counting, nil
}

func (repo *MongoDbRepository[TModel]) CreateOne(entity *TModel) (*TModel, error) {
	_, err := repo.MongoClient.Database(repo.Database).Collection(repo.Collection).InsertOne(context.TODO(), entity)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (repo *MongoDbRepository[TModel]) CreateMany(entities []*TModel) ([]*TModel, error) {
	_, err := repo.MongoClient.Database(repo.Database).Collection(repo.Collection).InsertMany(context.TODO(), funk.Map(entities, func(e *TModel) any {
		return e
	}).([]any))
	if err != nil {
		return nil, err
	}
	return entities, nil
}

func (repo *MongoDbRepository[TModel]) RemoveOne(expression func(TModel) bool) error {
	filter := GetFilter[TModel](expression)
	_, err := repo.MongoClient.Database(repo.Database).Collection(repo.Collection).DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (repo *MongoDbRepository[TModel]) RemoveMany(expression func(TModel) bool) error {
	filter := GetFilter[TModel](expression)
	_, err := repo.MongoClient.Database(repo.Database).Collection(repo.Collection).DeleteMany(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

func GetFilter[TModel any](expression func(TModel) bool) bson.D {
	var filter = bson.D{}
	if expression == nil {
		filter = bson.D{}
	}
	return filter
}
