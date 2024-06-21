package abstractions

import (
	"context"
	"natsMicros/buildingBlocks/application/responses"
)

type IDatabaseRepository[TModel any] interface {
	GetFirstByCondition(predict func(TModel) bool, ctx context.Context) (*TModel, error)
	ExistByCondition(predict func(TModel) bool, ctx context.Context) (bool, error)
	GetManyByCondition(predict func(TModel) bool, ctx context.Context) ([]*TModel, error)
	GetPaginationByCondition(predict func(TModel) bool, pageSize int64, pageIndex int64, ctx context.Context) (responses.PaginationResponse[*TModel], error)
	CountByCondition(predict func(TModel) bool, ctx context.Context) (int64, error)
	CreateOne(entity *TModel, ctx context.Context) (*TModel, error)
	CreateMany(entities []*TModel, ctx context.Context) ([]*TModel, error)
	RemoveOne(predict func(TModel) bool, ctx context.Context) error
	RemoveMany(predict func(TModel) bool, ctx context.Context) error
}
