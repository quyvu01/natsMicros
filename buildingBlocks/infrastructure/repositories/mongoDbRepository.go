package repositories

import "errors"

type MongoDbRepository[TModel any] struct {
	// Define the MongoDb Driver here
}

func NewMongoDbRepository[TModel any]() *MongoDbRepository[TModel] {
	return &MongoDbRepository[TModel]{}
}

func (repo *MongoDbRepository[TModel]) GetFirstByCondition(expression func(TModel) bool) (*TModel, error) {
	return nil, errors.New("not implemented")
}
func (repo *MongoDbRepository[TModel]) ExistByCondition(expression func(TModel) bool) (bool, error) {
	return false, errors.New("not implemented")
}

func (repo *MongoDbRepository[TModel]) GetManyByCondition(expression func(TModel) bool) (*TModel[], error) {
	return nil, errors.New("not implemented")
}

func (repo *MongoDbRepository[TModel]) CountByCondition(expression func(TModel) bool) (int64, error) {
	return 0, errors.New("not implemented")
}

func (repo *MongoDbRepository[TModel]) CreateOne(entity *TModel) (*TModel, error) {
	return nil, errors.New("not implemented")
}

func (repo *MongoDbRepository[TModel]) CreateMany(entities *TModel[]) (*TModel[], error) {
	return nil, errors.New("not implemented")
}

func (repo *MongoDbRepository[TModel]) RemoveOne(entity *TModel) (*TModel, error) {
	return nil, errors.New("not implemented")
}

func (repo *MongoDbRepository[TModel]) RemoveMany(entities *TModel[]) (*TModel[], error) {
	return nil, errors.New("not implemented")
}