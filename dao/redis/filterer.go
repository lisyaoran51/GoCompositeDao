package redis

type Filterer[TModel any, TQueryModel any] interface {
	Scopes(funcs ...func(models []*TModel) []*TModel) Filterer[TModel, TQueryModel]
	Filter(models []*TModel) []*TModel
	SetQueryModel(*TQueryModel)
	GetQueryModel() *TQueryModel
}

type FiltererImpl[TModel any, TQueryModel any] struct {
	callbacks  []func(models []*TModel) []*TModel
	queryModel *TQueryModel
}

func (f *FiltererImpl[TModel, TQueryModel]) Scopes(funcs ...func(models []*TModel) []*TModel) Filterer[TModel, TQueryModel] {
	f.callbacks = append(f.callbacks, funcs...)
	return f
}

func (f *FiltererImpl[TModel, TQueryModel]) Filter(models []*TModel) []*TModel {
	for _, ff := range f.callbacks {
		models = ff(models)
	}
	return models
}

func (f *FiltererImpl[TModel, TQueryModel]) SetQueryModel(qm *TQueryModel) {
	f.queryModel = qm
}

func (f *FiltererImpl[TModel, TQueryModel]) GetQueryModel() *TQueryModel {
	return f.queryModel
}
