package dao

import (
	"context"

	"github.com/lisyaoran51/GoCompositeDao/dao/specification"
)

type Dao[TModel any, TQueryField, TUpdateField comparable] interface {
	GetSourceType() string
	GetModelType() string
	GetDataName() string
	Register(ctx context.Context) error
	CreateDao(ctx context.Context) Dao[TModel, TQueryField, TUpdateField]
	Transaction(ctx context.Context, dataSource interface{}, txFunc func(dataSource interface{}) error) error

	New(ctx context.Context, dataSource interface{}, model *TModel) (interface{}, error)
	Count(ctx context.Context, dataSource interface{}, query QueryModel[TQueryField]) (int, error)
	Get(ctx context.Context, dataSource interface{}, query QueryModel[TQueryField]) (*TModel, error)
	Gets(ctx context.Context, dataSource interface{}, query QueryModel[TQueryField]) ([]*TModel, error)
	GetsWithPagination(ctx context.Context, dataSource interface{}, query QueryModel[TQueryField], paginate *specification.PaginationStruct) ([]*TModel, int, error)
	Modify(ctx context.Context, dataSource interface{}, model *TModel, fields []TUpdateField) error
	Delete(ctx context.Context, dataSource interface{}, query QueryModel[TQueryField]) error
	GetQueryModel() QueryModel[TQueryField]
}
