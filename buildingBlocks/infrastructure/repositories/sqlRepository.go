package repositories

import (
	"gorm.io/gorm"
	"natsMicros/buildingBlocks/application/responses"
)

type SqlRepository[TModel any] struct {
	Db *gorm.DB
}

func NewSqlRepository[TModel any](db *gorm.DB) *SqlRepository[TModel] {
	return &SqlRepository[TModel]{Db: db}
}

func (repo *SqlRepository[TModel]) GetFirstByCondition(predict func(TModel) bool) (*TModel, error) {
	var model TModel
	repo.Db.Where(predict).First(&model)
	return &model, nil
}
func (repo *SqlRepository[TModel]) ExistByCondition(predict func(TModel) bool) (bool, error) {
	var model TModel
	res := repo.Db.Where(predict).First(&model)
	if res.Error != nil {
		return false, res.Error
	}
	return true, nil
}
func (repo *SqlRepository[TModel]) GetManyByCondition(predict func(TModel) bool) ([]*TModel, error) {
	return nil, nil
}
func (repo *SqlRepository[TModel]) GetPaginationByCondition(predict func(TModel) bool) (responses.PaginationResponse[*TModel], error) {
	return responses.PaginationResponse[*TModel]{}, nil
}
func (repo *SqlRepository[TModel]) CountByCondition(predict func(TModel) bool) (int64, error) {
	return 0, nil
}
func (repo *SqlRepository[TModel]) CreateOne(entity *TModel) (*TModel, error) {
	return nil, nil
}
func (repo *SqlRepository[TModel]) CreateMany(entities []*TModel) ([]*TModel, error) {
	return nil, nil
}
func (repo *SqlRepository[TModel]) RemoveOne(predict func(TModel) bool) error {
	return nil
}
func (repo *SqlRepository[TModel]) RemoveMany(predict func(TModel) bool) error {
	return nil
}
