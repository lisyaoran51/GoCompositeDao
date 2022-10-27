package dao

import "fmt"

type QueryModel[T comparable] interface {
	AddCondition(field T, condition ...interface{}) QueryModel[T]
	AddContidionToModel(field T, condition ...interface{}) bool

	// Set
	//
	// if you have QueryModel in another type, you could absorb that
	Set(QueryModel[T]) QueryModel[T]

	GetConditions() map[T][]interface{}
}

type QueryModelImpl[T comparable] struct {
	// QueryModel contains itself
	QueryModel[T]
	Conditions map[T][]interface{}
}

func (q *QueryModelImpl[T]) AddCondition(field T, condition ...interface{}) QueryModel[T] {
	if len(condition) == 0 {
		fmt.Printf("AddCondition: no condition added\n")
		return q.QueryModel
	}

	if q.AddContidionToModel(field, condition...) {
		q.Conditions[field] = condition
	}

	return q.QueryModel
}

func (q *QueryModelImpl[T]) AddContidionToModel(field T, condition ...interface{}) bool {

	return q.QueryModel.AddContidionToModel(field, condition...)
}

func (q *QueryModelImpl[T]) Set(originalQueryModel QueryModel[T]) QueryModel[T] {

	for k, v := range originalQueryModel.GetConditions() {

		q.AddCondition(k, v...)
	}

	return q.QueryModel
}
