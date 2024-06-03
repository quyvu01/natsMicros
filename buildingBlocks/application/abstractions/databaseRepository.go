package abstractions

import "natsMicros/buildingBlocks/application/responses"

type IDatabaseRepository[TModel any] interface {
	GetFirstByCondition(predict func(TModel) bool) (*TModel, error)
	ExistByCondition(predict func(TModel) bool) (bool, error)
	GetManyByCondition(predict func(TModel) bool) ([]*TModel, error)
	GetPaginationByCondition(predict func(TModel) bool) (responses.PaginationResponse[*TModel], error)
	CountByCondition(predict func(TModel) bool) (int64, error)
	CreateOne(entity *TModel) (*TModel, error)
	CreateMany(entities []*TModel) ([]*TModel, error)
	RemoveOne(predict func(TModel) bool) error
	RemoveMany(predict func(TModel) bool) error
}
