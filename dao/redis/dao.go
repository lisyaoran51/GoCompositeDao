package redis

import (
	"context"
	"fmt"
	"reflect"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/go-redis/redis/v9"
	goCompositeDao "github.com/lisyaoran51/GoCompositeDao"
	"github.com/lisyaoran51/GoCompositeDao/dao"
	"github.com/lisyaoran51/GoCompositeDao/dao/specification"
)

type REDIS_STRATEGY int

const (
	REDIS_STRATEGY_NONE REDIS_STRATEGY = iota
	REDIS_STRATEGY_REDIS_ONLY
	REDIS_STRATEGY_CACHE_ASIDE
	REDIS_STRATEGY_DOUBLE_DELETE
	REDIS_STRATEGY_RW_THROUGH
	REDIS_STRATEGY_WRITE_BEHIND
)

type WalletRecordRedisDao[TModel any, TQueryField, TUpdateField comparable] Dao[TModel, TQueryField, TUpdateField]

type Dao[TModel any, TQueryField, TUpdateField comparable] interface {
	dao.Dao[TModel, TQueryField, TUpdateField]

	// GetKeysStringKey
	//
	// get the key of a string field in redis, which contains all the chache keys in single string, devided by ' '
	//
	// this should be unuseful. no maintainence anymore
	GetKeysStringKey() string

	// GetKeysSetKey
	//
	// get the key of a set in redis, which contains all the chache keys
	GetKeysSetKey() string

	// GetMaskQuery
	//
	// some of the query condition should be mask in redis in order to ultimate utility
	GetMaskQuery() mapset.Set[TQueryField]
	AddKeys(ctx context.Context, rdb *redis.Client, keys []string) error
	GetKeys(ctx context.Context, rdb *redis.Client) ([]string, error)
	CleanKeys(ctx context.Context, rdb *redis.Client, keys []string) error
	Clean(ctx context.Context, rdb *redis.Client, cleanFunc func([]string) []string) error

	DeepNew(ctx context.Context, dataSource interface{}, model *TModel) (interface{}, error)
	DeepCount(ctx context.Context, dataSource interface{}, query dao.QueryModel[TQueryField]) (int, error)
	DeepGet(ctx context.Context, dataSource interface{}, query dao.QueryModel[TQueryField]) (*TModel, error)
	DeepGets(ctx context.Context, dataSource interface{}, query dao.QueryModel[TQueryField]) ([]*TModel, error)
	DeepGetsWithPagination(ctx context.Context, dataSource interface{}, query dao.QueryModel[TQueryField], paginate *specification.PaginationStruct) ([]*TModel, int, error)
	DeepModify(ctx context.Context, dataSource interface{}, model *TModel, fields []TUpdateField) error
	DeepDelete(ctx context.Context, dataSource interface{}, query dao.QueryModel[TQueryField]) error
}

type DaoImpl[TModel any, TQueryField, TUpdateField comparable] struct {
	// Dao contains itself
	dao.Dao[TModel, TQueryField, TUpdateField]
}

func (r *DaoImpl[TModel, TQueryField, TUpdateField]) GetSourceType() string {
	source := &CompositeRedis{}
	return reflect.TypeOf(source).String()
}

func (r *DaoImpl[TModel, TQueryField, TUpdateField]) Transaction(ctx context.Context, dataSource interface{}, txFunc func(dataSource interface{}) error) error {
	if rdb, ok := dataSource.(*CompositeRedis); ok {
		return r.transaction(ctx, rdb, txFunc)
	}
	return goCompositeDao.ErrInternal
}

func (r *DaoImpl[TModel, TQueryField, TUpdateField]) transaction(ctx context.Context, rdb *CompositeRedis, txFunc func(dataSource interface{}) error) error {

	// directly go to deep transaction since redis has no transaction
	// delete all related cache
	// may need to provide a way to clean specific cache

	return nil
}

func (r *DaoImpl[TModel, TQueryField, TUpdateField]) AddKeys(ctx context.Context, rdb *redis.Client, keys []string) error {

	if len(keys) == 0 {
		return nil
	}

	_, err := rdb.Get(ctx, r.GetKeysStringKey()).Result()
	if err != nil && err != redis.Nil {
		return err
	}

	if err == nil {
		_, err := rdb.Del(ctx, r.GetKeysStringKey()).Result()
		if err != nil && err != redis.Nil {
			return err
		}
	}

	keysInInterface := []interface{}{}
	for _, k := range keys {
		keysInInterface = append(keysInInterface, k)
	}

	rdb.SAdd(ctx, r.GetKeysSetKey(), keysInInterface...).Result()

	return nil
}

func (r *DaoImpl[TModel, TQueryField, TUpdateField]) GetKeysStringKey() string {
	return r.GetDataName() + ":" + "KeyString"
}

func (r *DaoImpl[TModel, TQueryField, TUpdateField]) GetKeysSetKey() string {
	return r.GetDataName() + ":" + "KeySet"
}

func (r *DaoImpl[TModel, TQueryField, TUpdateField]) GetKeys(ctx context.Context, rdb *redis.Client) ([]string, error) {
	// keys, err := rdb.Get(ctx, r.GetKeysStringKey()).Result()
	// if err != nil && err != redis.Nil {
	// 	return nil, err
	// }

	// if err == redis.Nil {
	// 	keysSlice := strings.Split(keys, " ")
	// 	return keysSlice, nil
	// }

	keysSlice, err := rdb.SMembers(ctx, r.GetKeysSetKey()).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}

	return keysSlice, err
}

func (r *DaoImpl[TModel, TQueryField, TUpdateField]) CleanKeys(ctx context.Context, rdb *redis.Client, keys []string) error {

	_, err := rdb.Del(ctx, r.GetKeysStringKey()).Result()
	if err != nil && err != redis.Nil {
		return err
	}

	_, err = rdb.Del(ctx, r.GetKeysSetKey()).Result()
	if err != nil && err != redis.Nil {
		return err
	}

	return nil
}

func (r *DaoImpl[TModel, TQueryField, TUpdateField]) Clean(ctx context.Context, rdb *redis.Client, selectFunc func([]string) []string) error {

	keys, err := r.Dao.(Dao[TModel, TQueryField, TUpdateField]).GetKeys(ctx, rdb)
	if err != nil {
		return err
	}

	keys = selectFunc(keys)

	// ??
	// TOSO: delete selected keys only
	// r.CleanKeys(ctx, rdb, keys)

	for _, k := range keys {

		result, err := rdb.Del(ctx, k).Result()
		if err != nil && err != redis.Nil {
			fmt.Printf("Clean: failed to delete redis key [%s], %v\n", k, err)
			return err
		} else if err == redis.Nil {
			fmt.Printf("Clean: failed to delete redis key [%s]\n", k)
		}

		if result < int64(len(keys)) {
			// add some warning..
		}

		_, err = rdb.SRem(ctx, r.GetKeysSetKey(), k).Result()
		if err != nil && err != redis.Nil {
			fmt.Printf("Clean: failed to delete redis key [%s] from set [%s] %v\n", k, r.GetKeysSetKey(), err)
			return err
		} else if err == redis.Nil {
			fmt.Printf("Clean: failed to delete redis key [%s] from set [%s]\n", k, r.GetKeysSetKey())
		}
	}
	return nil
}

/// ======================== REGION ========================
/// below is only for inheritance assertion. no business logic here
/// ========================================================

/// ============= Dao =============
