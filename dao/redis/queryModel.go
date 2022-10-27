package redis

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/lisyaoran51/GoCompositeDao/dao"
)

type QueryModel[T comparable] interface {
	dao.QueryModel[T]

	GetTagFromField(field T) string
	GetFieldFromTag(tag string) T

	// HasRequiredQuery
	//
	// if require query not set, return false
	HasRequiredQuery() bool

	// GetMaskedQueryModel
	//
	// get a model some of whiches query is masked
	GetMaskedQueryModel() QueryModel[T]
}

type QueryModelImpl[T comparable] struct {
	/// RequireQuery
	///
	/// this dao must use query with this field, other return error
	RequireQuery mapset.Set[T]
	/// MaskQuery
	///
	/// when set into dao, this field will be masked
	MaskQuery mapset.Set[T]
	dao.QueryModelImpl[T]
}

func (q *QueryModelImpl[T]) GetKey() string {

	key := ""

	v := reflect.ValueOf(q.QueryModel)
	t := reflect.TypeOf(q.QueryModel)
	for i := 0; i < v.Elem().NumField(); i++ {

		if t.Elem().Field(i).Tag == "" {
			continue
		}

		tag, ok := t.Elem().Field(i).Tag.Lookup("redis")
		if !ok {
			fmt.Printf("no redis tag in [%s]\n", t.Elem().Field(i).Tag)
		}
		if v.Elem().Field(i).IsNil() {
			continue
		}

		if q.MaskQuery.Contains(q.GetFieldFromTag(tag)) {
			fmt.Printf("reflect mask tag %s %s %#v\n", tag, t.Elem().Field(i).Type.Kind(), v.Elem().Field(i).Interface())
			continue
		}

		value := v.Elem().Field(i).Interface()

		if t, ok := value.(time.Time); ok {
			value = t.Unix()
		}
		if t, ok := value.(*time.Time); ok {
			value = t.Unix()
		}

		if key != "" {
			key += ":"
		}

		valueMarshalled, _ := json.Marshal(value)
		key += tag + ":" + string(valueMarshalled)
	}
	return key
}

func (q *QueryModelImpl[T]) GetQueryFieldFromKey(field T, key string) string {
	keys := strings.Split(key, ":")
	keys = keys[1:]

	if len(keys)%2 == 1 {
		keys = keys[:len(keys)-1]
	}

	tag := q.GetTagFromField(field)
	for i := 0; i < len(keys); i += 2 {
		if keys[i] == tag {
			return keys[i+1]
		}
	}

	fmt.Printf("GetQueryFieldFromKey: no such field in key [%s] %s\n", keys, key)

	return ""
}

func (q *QueryModelImpl[T]) GetTagFromField(field T) string {
	qm, ok := q.QueryModel.(QueryModel[T])
	if !ok {
		panic("GetTagFromField: failed to cast to child.")
	}
	return qm.GetTagFromField(field)
}

func (q *QueryModelImpl[T]) GetFieldFromTag(tag string) T {
	qm, ok := q.QueryModel.(QueryModel[T])
	if !ok {
		panic("GetFieldFromTag: failed to cast to child.")
	}
	return qm.GetFieldFromTag(tag)
}
