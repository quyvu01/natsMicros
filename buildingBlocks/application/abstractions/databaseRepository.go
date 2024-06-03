package abstractions

type IDatabaseRepository[TModel any] interface {
	GetFirstByCondition(expression func(TModel) bool) (*TModel, error)
	ExistByCondition(expression func(TModel) bool) (bool, error)
	GetManyByCondition(expression func(TModel) bool) (*TModel[], error)
	CountByCondition(expression func(TModel) bool) (int64, error)
	CreateOne(entity *TModel) (*TModel, error)
	CreateMany(entities *TModel[]) (*TModel[], error)
	RemoveOne(entity *TModel) (*TModel, error)
	RemoveMany(entities *TModel[]) (*TModel[], error)
}
