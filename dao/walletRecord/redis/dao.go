package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
	redisLib "github.com/go-redis/redis/v9"
	goCompositeDao "github.com/lisyaoran51/GoCompositeDao"
	"github.com/lisyaoran51/GoCompositeDao/dao"
	"github.com/lisyaoran51/GoCompositeDao/dao/redis"
	"github.com/lisyaoran51/GoCompositeDao/dao/specification"
	"github.com/lisyaoran51/GoCompositeDao/dao/walletRecord"
	databaseModels "github.com/lisyaoran51/GoCompositeDao/models/databaseModels"
)

var maskQuery []walletRecord.QueryModelField = []walletRecord.QueryModelField{
	// walletRecord.QueryModelField_IP,
	// walletRecord.QueryModelField_Account,
	// walletRecord.QueryModelField_Downlines,
	// walletRecord.QueryModelField_OrderBy,
	walletRecord.QueryModelField_Limit,
	walletRecord.QueryModelField_Offset,
	walletRecord.QueryModelField_LockOrNot,
	// walletRecord.QueryModelField_RecordType,
	// walletRecord.QueryModelField_Action,
	// walletRecord.QueryModelField_Modifier,
	walletRecord.QueryModelField_DateStart,
	walletRecord.QueryModelField_DateEnd,
	// walletRecord.QueryModelField_Actions,
	// walletRecord.QueryModelField_UserAccounts,
	// walletRecord.QueryModelField_Amount,
}

var requireQuery []walletRecord.QueryModelField = []walletRecord.QueryModelField{
	// walletRecord.QueryModelField_IP,
	walletRecord.QueryModelField_Account,
	// walletRecord.QueryModelField_Downlines,
	// walletRecord.QueryModelField_OrderBy,
	// walletRecord.QueryModelField_Limit,
	// walletRecord.QueryModelField_Offset,
	// walletRecord.QueryModelField_LockOrNot,
	// walletRecord.QueryModelField_RecordType,
	// walletRecord.QueryModelField_Action,
	// walletRecord.QueryModelField_Modifier,
	// walletRecord.QueryModelField_DateStart,
	// walletRecord.QueryModelField_DateEnd,
	// walletRecord.QueryModelField_Actions,
	// walletRecord.QueryModelField_UserAccounts,
	// walletRecord.QueryModelField_Amount,
}

type RedisDao = redis.DaoImpl[databaseModels.WalletRecordModel, walletRecord.QueryModelField, walletRecord.UpdateField]
type Dao struct {
	RedisDao
	walletRecord.DaoImpl
}

func NewDao() *Dao {
	newDao := &Dao{}
	newDao.DaoImpl.Dao = newDao
	newDao.RedisDao.Dao = newDao
	return newDao
}

// daoAssertion
//
// this is only for compile assertion. no use in program.
func assertion() {
	newDao := &Dao{}
	var d1 dao.Dao[databaseModels.WalletRecordModel, walletRecord.QueryModelField, walletRecord.UpdateField] = newDao
	var d2 redis.Dao[databaseModels.WalletRecordModel, walletRecord.QueryModelField, walletRecord.UpdateField] = newDao
	var d3 walletRecord.Dao = newDao

	fmt.Println(d1, d2, d3)
}

func (c *Dao) new(ctx context.Context, rdb *redis.CompositeRedis, model *databaseModels.WalletRecordModel) (interface{}, error) {

	result, err := c.DeepNew(ctx, rdb.DB, model)

	if err != nil {
		fmt.Printf("new DeepNew failed %v\n", err)
		return nil, err
	}

	//delete cache

	c.Clean(ctx, &rdb.Redis, func(keys []string) []string {

		selectedKeys := []string{}

		queryModel := c.GetQueryModel().(*QueryModel)

		valueMarshalled, _ := json.Marshal(model.Account)

		for _, k := range keys {

			if queryModel.GetQueryFieldFromKey(walletRecord.QueryModelField_Account, k) == string(valueMarshalled) {
				selectedKeys = append(selectedKeys, k)
			}
		}

		return selectedKeys
	})

	return result, err
}

func (c *Dao) count(ctx context.Context, rdb *redis.CompositeRedis, query *QueryModel) (int, error) {

	return 0, goCompositeDao.ErrNotImplemented
}

func (c *Dao) get(ctx context.Context, rdb *redis.CompositeRedis, query *QueryModel) (*databaseModels.WalletRecordModel, error) {

	if !query.HasRequiredQuery() { /* must have required query field, otherwise it go straight to db.*/
		return c.DeepGet(ctx, rdb.DB, query)
	}

	key := c.GetDataName() + ":" + query.GetKey()

	result, err := rdb.Redis.Get(ctx, key).Result()

	var models []*databaseModels.WalletRecordModel

	if err != nil && err != redisLib.Nil {
		return nil, err
	}

	if err == redisLib.Nil {

		maskedQuery := query.GetMaskedQueryModel()

		models, err = c.DeepGets(ctx, rdb.DB, maskedQuery)
		if err != nil {
			return nil, err
		}

		b, err := json.Marshal(models)
		if err != nil {
			fmt.Printf("gets marshal fail %v\n", err)
			return nil, err
		}

		if err = rdb.Redis.Set(ctx, key, string(b), 0).Err(); err != nil {
			fmt.Printf("gets redis set [%s] fail %v\n", string(b), err)
			return nil, err
		}

		if err = rdb.Redis.SAdd(ctx, c.GetKeysSetKey(), key).Err(); err != nil {
			fmt.Printf("gets redis set key [%s] fail %v\n", c.GetKeysSetKey(), err)
			return nil, err
		}

	} else {
		err = json.Unmarshal([]byte(result), &models)
		if err != nil {
			return nil, err
		}
	}

	query = query.AddCondition(walletRecord.QueryModelField_Limit, 1).(*QueryModel) /* to get the last one row */

	filter := NewFilterer(query)
	models = filter.
		Scopes(filter.filterDateStart()).
		Scopes(filter.filterDateEnd()).
		Scopes(filter.sortOffset()).
		Scopes(filter.sortLimit()).
		Filter(models)

	if len(models) == 0 {
		return nil, nil
	}

	return models[0], nil
}

func (c *Dao) gets(ctx context.Context, rdb *redis.CompositeRedis, query *QueryModel) ([]*databaseModels.WalletRecordModel, error) {

	if !query.HasRequiredQuery() { /* must have required query field, otherwise it go straight to db.*/
		return c.DeepGets(ctx, rdb.DB, query)
	}

	var models []*databaseModels.WalletRecordModel

	key := c.GetDataName() + ":" + query.GetKey()

	result, err := rdb.Redis.Get(ctx, key).Result()

	if err != nil && err != redisLib.Nil {
		fmt.Printf("gets error %v\n", err)
		return nil, err
	}

	if err == redisLib.Nil {

		maskedQuery := query.GetMaskedQueryModel()

		models, err = c.DeepGets(ctx, rdb.DB, maskedQuery)
		if err != nil {
			return nil, err
		}

		b, err := json.Marshal(models)
		if err != nil {
			fmt.Printf("gets marshal fail %v\n", err)
			return nil, err
		}

		if err = rdb.Redis.Set(ctx, key, string(b), 0).Err(); err != nil {
			fmt.Printf("gets redis set [%s] fail %v\n", string(b), err)
			return nil, err
		}

		if err = rdb.Redis.SAdd(ctx, c.GetKeysSetKey(), key).Err(); err != nil {
			fmt.Printf("gets redis set key [%s] fail %v\n", c.GetKeysSetKey(), err)
			return nil, err
		}

	} else {
		err = json.Unmarshal([]byte(result), &models)
		if err != nil {
			return nil, err
		}
	}

	filter := NewFilterer(query)
	models = filter.
		Scopes(filter.filterDateStart()).
		Scopes(filter.filterDateEnd()).
		Scopes(filter.sortOffset()).
		Scopes(filter.sortLimit()).
		Filter(models)

	return models, nil
}

func (c *Dao) getsWithPagination(ctx context.Context, rdb *redis.CompositeRedis, query *QueryModel, paginate *specification.PaginationStruct) (_ []*databaseModels.WalletRecordModel, _ int, err error) {

	if !query.HasRequiredQuery() { /* must have required query field, otherwise it go straight to db.*/
		return c.DeepGetsWithPagination(ctx, rdb.DB, query, paginate)
	}

	var models []*databaseModels.WalletRecordModel

	key := c.GetDataName() + ":" + query.GetKey()

	result, err := rdb.Redis.Get(ctx, key).Result()

	if err != nil && err != redisLib.Nil {
		fmt.Printf("getsWithPagination error %v\n", err)
		return nil, 0, err
	}

	if err == redisLib.Nil {

		maskedQuery := query.GetMaskedQueryModel()

		models, err = c.DeepGets(ctx, rdb.DB, maskedQuery)
		if err != nil {
			fmt.Printf("getsWithPagination error %v\n", err)
			return nil, 0, err
		}

		b, err := json.Marshal(models)
		if err != nil {
			fmt.Printf("getsWithPagination marshal fail %v\n", err)
			return nil, 0, err
		}

		if err = rdb.Redis.Set(ctx, key, string(b), 0).Err(); err != nil {
			fmt.Printf("getsWithPagination redis set [%s] fail %v\n", string(b), err)
			return nil, 0, err
		}

		if err = rdb.Redis.SAdd(ctx, c.GetKeysSetKey(), key).Err(); err != nil {
			fmt.Printf("getsWithPagination redis set key [%s] fail %v\n", c.GetKeysSetKey(), err)
			return nil, 0, err
		}

	} else {
		err = json.Unmarshal([]byte(result), &models)
		if err != nil {
			return nil, 0, err
		}
	}

	filter := NewFilterer(query)
	models = filter.
		Scopes(filter.filterDateStart()).
		Scopes(filter.filterDateEnd()).
		Filter(models)

	count := len(models)

	query.AddCondition(walletRecord.QueryModelField_Offset, int(paginate.Index))
	query.AddCondition(walletRecord.QueryModelField_Limit, int(paginate.PageSize))

	filter = NewFilterer(query)
	models = filter.
		Scopes(filter.filterDateStart()).
		Scopes(filter.filterDateEnd()).
		Filter(models)

	return models, count, nil
}

func (c *Dao) modify(ctx context.Context, rdb *redis.CompositeRedis, model *databaseModels.WalletRecordModel, fields []walletRecord.UpdateField) error {

	/* redis dao do not allow batch update without required field */
	/* it will damage all data in redis */
	if model.Account == "" {
		fmt.Printf("redis dao: no required index\n")
		return errors.New("wallet account not specified")
	}

	err := c.DeepModify(ctx, rdb.DB, model, fields)

	//delete cache
	c.Clean(ctx, &rdb.Redis, func(keys []string) []string {

		selectedKeys := []string{}

		queryModel := c.GetQueryModel().(*QueryModel)

		for _, k := range keys {

			valueMarshalled, _ := json.Marshal(model.Account)

			if queryModel.GetQueryFieldFromKey(walletRecord.QueryModelField_Account, k) == string(valueMarshalled) {
				selectedKeys = append(selectedKeys, k)
			}
		}
		return selectedKeys
	})

	return err
}

func (c *Dao) delete(ctx context.Context, rdb *redis.CompositeRedis, query *QueryModel) error {

	/* redis dao do not allow batch delete without required field */
	/* it will damage all data in redis */
	if !query.HasRequiredQuery() {
		fmt.Printf("redis dao: no required index\n")
		return errors.New("wallet account not specified")
	}

	err := c.DeepDelete(ctx, rdb.DB, query)

	c.Clean(ctx, &rdb.Redis, func(keys []string) []string {

		selectedKeys := []string{}

		queryModel := c.GetQueryModel().(*QueryModel)

		for _, k := range keys {

			valueMarshalled, _ := json.Marshal(*query.Account)

			if queryModel.GetQueryFieldFromKey(walletRecord.QueryModelField_Account, k) == string(valueMarshalled) {
				selectedKeys = append(selectedKeys, k)
			}
		}

		return selectedKeys
	})

	return err
}

/// ======================== REGION ========================
/// below is only for inheritance assertion. no business logic here
/// ========================================================

/// ============= dao.Dao =============

func (c *Dao) CreateDao(ctx context.Context) dao.Dao[databaseModels.WalletRecordModel, walletRecord.QueryModelField, walletRecord.UpdateField] {
	return NewDao()
}

/// ============= walletRecord.Dao =============

func (c *Dao) New(ctx context.Context, dataSource interface{}, model *databaseModels.WalletRecordModel) (result interface{}, err error) {

	if rdb, ok := dataSource.(*redis.CompositeRedis); ok {
		return c.new(ctx, rdb, model)
	}

	return nil, goCompositeDao.ErrInternal
}

func (c *Dao) Count(ctx context.Context, dataSource interface{}, query dao.QueryModel[walletRecord.QueryModelField]) (_ int, err error) {

	q, ok := query.(*QueryModel)
	if !ok {
		return 0, goCompositeDao.ErrInternal
	}

	if rdb, ok := dataSource.(*redis.CompositeRedis); ok {
		return c.count(ctx, rdb, q)
	}
	return 0, goCompositeDao.ErrInternal

}

func (c *Dao) Get(ctx context.Context, dataSource interface{}, query dao.QueryModel[walletRecord.QueryModelField]) (result *databaseModels.WalletRecordModel, err error) {

	q, ok := query.(*QueryModel)
	if !ok {
		return nil, goCompositeDao.ErrInternal
	}

	if rdb, ok := dataSource.(*redis.CompositeRedis); ok {
		return c.get(ctx, rdb, q)
	}

	return nil, goCompositeDao.ErrInternal
}

func (c *Dao) Gets(ctx context.Context, dataSource interface{}, query dao.QueryModel[walletRecord.QueryModelField]) (result []*databaseModels.WalletRecordModel, err error) {

	q, ok := query.(*QueryModel)
	if !ok {
		return nil, goCompositeDao.ErrInternal
	}

	if rdb, ok := dataSource.(*redis.CompositeRedis); ok {
		return c.gets(ctx, rdb, q)
	}

	return nil, goCompositeDao.ErrInternal
}

func (c *Dao) GetsWithPagination(ctx context.Context, dataSource interface{}, query dao.QueryModel[walletRecord.QueryModelField], paginate *specification.PaginationStruct) (_ []*databaseModels.WalletRecordModel, _ int, err error) {

	q, ok := query.(*QueryModel)
	if !ok {
		fmt.Printf("Dao Gets QueryModel casting failed.\n")
		return nil, 0, goCompositeDao.ErrInternal
	}

	if rdb, ok := dataSource.(*redis.CompositeRedis); ok {
		return c.getsWithPagination(ctx, rdb, q, paginate)
	}

	fmt.Printf("Dao Gets casting failed.\n")
	return nil, 0, goCompositeDao.ErrInternal
}

func (c *Dao) Modify(ctx context.Context, dataSource interface{}, model *databaseModels.WalletRecordModel, fields []walletRecord.UpdateField) (err error) {

	if rdb, ok := dataSource.(*redis.CompositeRedis); ok {
		return c.modify(ctx, rdb, model, fields)
	}

	return goCompositeDao.ErrInternal
}

func (c *Dao) Delete(ctx context.Context, dataSource interface{}, query dao.QueryModel[walletRecord.QueryModelField]) (err error) {

	q, ok := query.(*QueryModel)
	if !ok {
		return goCompositeDao.ErrInternal
	}

	if rdb, ok := dataSource.(*redis.CompositeRedis); ok {
		return c.delete(ctx, rdb, q)
	}

	return goCompositeDao.ErrInternal
}

func (c *Dao) GetQueryModel() dao.QueryModel[walletRecord.QueryModelField] {
	q := NewQueryModel()
	q.MaskQuery = c.GetMaskQuery()
	q.RequireQuery = c.GetRequiredQuery()
	return q
}

/// ============= redis.Dao =============

func (c *Dao) GetRequiredQuery() mapset.Set[walletRecord.QueryModelField] {
	return mapset.NewSet(requireQuery...)
}

func (c *Dao) GetMaskQuery() mapset.Set[walletRecord.QueryModelField] {
	return mapset.NewSet(maskQuery...)
}

func (c *Dao) DeepNew(ctx context.Context, dataSource interface{}, model *databaseModels.WalletRecordModel) (interface{}, error) {

	dao := walletRecord.GetDao(ctx, dataSource)
	if dao == nil {
		return nil, nil
	}

	result, err := dao.New(ctx, dataSource, model)

	return result, err
}

func (c *Dao) DeepCount(ctx context.Context, dataSource interface{}, query dao.QueryModel[walletRecord.QueryModelField]) (int, error) {

	dao := walletRecord.GetDao(ctx, dataSource)
	if dao == nil {
		return 0, nil
	}

	newQuery := dao.GetQueryModel().Set(query)

	result, err := dao.Count(ctx, dataSource, newQuery)

	return result, err
}

func (c *Dao) DeepGet(ctx context.Context, dataSource interface{}, query dao.QueryModel[walletRecord.QueryModelField]) (*databaseModels.WalletRecordModel, error) {

	dao := walletRecord.GetDao(ctx, dataSource)
	if dao == nil {
		return nil, nil
	}

	newQuery := dao.GetQueryModel().Set(query)

	result, err := dao.Get(ctx, dataSource, newQuery)

	return result, err
}

func (c *Dao) DeepGets(ctx context.Context, dataSource interface{}, query dao.QueryModel[walletRecord.QueryModelField]) ([]*databaseModels.WalletRecordModel, error) {

	dao := walletRecord.GetDao(ctx, dataSource)
	if dao == nil {
		return nil, nil
	}

	newQuery := dao.GetQueryModel().Set(query)

	result, err := dao.Gets(ctx, dataSource, newQuery)

	return result, err
}

func (c *Dao) DeepGetsWithPagination(ctx context.Context, dataSource interface{}, query dao.QueryModel[walletRecord.QueryModelField], paginate *specification.PaginationStruct) ([]*databaseModels.WalletRecordModel, int, error) {

	dao := walletRecord.GetDao(ctx, dataSource)
	if dao == nil {
		return nil, 0, nil
	}

	newQuery := dao.GetQueryModel().Set(query)

	return dao.GetsWithPagination(ctx, dataSource, newQuery, paginate)

}

func (c *Dao) DeepModify(ctx context.Context, dataSource interface{}, model *databaseModels.WalletRecordModel, fields []walletRecord.UpdateField) error {

	dao := walletRecord.GetDao(ctx, dataSource)
	if dao == nil {
		return nil
	}

	err := dao.Modify(ctx, dataSource, model, fields)

	return err
}

func (c *Dao) DeepDelete(ctx context.Context, dataSource interface{}, query dao.QueryModel[walletRecord.QueryModelField]) error {

	dao := walletRecord.GetDao(ctx, dataSource)
	if dao == nil {
		return nil
	}

	err := dao.Delete(ctx, dataSource, query)

	return err
}
