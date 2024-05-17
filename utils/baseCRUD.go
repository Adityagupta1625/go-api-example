package base

import (
	"github.com/jinzhu/gorm"
)

type BaseCRUD[T any] interface {
	add(entity *T) error
	getById(id string) (*T, error)
	update(entity *T) error
	delete(entity *T) error
	get(condition string, args any) ([]any, error)
	getPaginated(condition string, args any, orderField string, page int, limit int) ([]any, error)
}

type GormRepository struct {
	DB *gorm.DB
}

func (gormRepository *GormRepository) get(condition string, args any) ([]any, error) {
	var data []any
	err := gormRepository.DB.Where(condition, args).Find(&data).Error

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (gormRepository *GormRepository) getPaginated(condition string, args any, orderField string, page int, limit int) ([]any, error) {

	var data []any
	err := gormRepository.DB.Order(orderField).Where(condition, args).Limit(limit).Offset((page - 1) * limit).Find(&data).Error

	if err != nil {
		return nil, err
	}

	return data, nil
}

func NewBaseRespositry(database *gorm.DB) *GormRepository {
	return &GormRepository{
		DB: database,
	}
}
