package abstractions

import "natsMicros/buildingBlocks/application/responses"

type IDatabaseRepository[TModel any] interface {
	GetFirstByCondition(expression func(TModel) bool) (*TModel, error)
	ExistByCondition(expression func(TModel) bool) (bool, error)
	GetManyByCondition(expression func(TModel) bool) ([]*TModel, error)
	GetPaginationByCondition(expression func(TModel) bool) (responses.PaginationResponse[*TModel], error)
	CountByCondition(expression func(TModel) bool) (int64, error)
	CreateOne(entity *TModel) (*TModel, error)
	CreateMany(entities []*TModel) ([]*TModel, error)
	RemoveOne(expression func(TModel) bool) error
	RemoveMany(expression func(TModel) bool) error
}
