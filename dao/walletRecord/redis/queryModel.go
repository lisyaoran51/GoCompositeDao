package redis

import (
	"fmt"
	"reflect"

	"github.com/lisyaoran51/GoCompositeDao/dao"
	"github.com/lisyaoran51/GoCompositeDao/dao/redis"
	"github.com/lisyaoran51/GoCompositeDao/dao/walletRecord"
	"github.com/shopspring/decimal"
)

type redisQueryModelImpl = redis.QueryModelImpl[walletRecord.QueryModelField]
type walletRecordQueryModelImpl = walletRecord.QueryModelImpl

type QueryModel struct {
	redisQueryModelImpl
	walletRecordQueryModelImpl
	IP           *string          `redis:"ip"`
	Account      *string          `redis:"account"`
	Downlines    *[]string        `redis:"downlines"`
	OrderBy      *[]string        `redis:"order_by"`
	Limit        *int             `redis:"limit"`
	Offset       *int             `redis:"offset"`
	LockOrNot    *bool            `redis:"lock_or_not"`
	RecordType   *int             `redis:"record_type"`
	Action       *string          `redis:"action"`
	Modifier     *string          `redis:"modifier"`
	DateStart    *string          `redis:"date_start"`
	DateEnd      *string          `redis:"date_end"`
	Actions      *[]string        `redis:"actions"`
	UserAccounts *[]string        `redis:"user_accounts"`
	Amount       *decimal.Decimal `redis:"amount"`
}

var QueryFieldRedisMap map[string]walletRecord.QueryModelField = map[string]walletRecord.QueryModelField{
	"ip":            walletRecord.QueryModelField_IP,
	"account":       walletRecord.QueryModelField_Account,
	"downlines":     walletRecord.QueryModelField_Downlines,
	"order_by":      walletRecord.QueryModelField_OrderBy,
	"limit":         walletRecord.QueryModelField_Limit,
	"offset":        walletRecord.QueryModelField_Offset,
	"lock_or_not":   walletRecord.QueryModelField_LockOrNot,
	"record_type":   walletRecord.QueryModelField_RecordType,
	"action":        walletRecord.QueryModelField_Action,
	"modifier":      walletRecord.QueryModelField_Modifier,
	"date_start":    walletRecord.QueryModelField_DateStart,
	"date_end":      walletRecord.QueryModelField_DateEnd,
	"actions":       walletRecord.QueryModelField_Actions,
	"user_accounts": walletRecord.QueryModelField_UserAccounts,
	"amount":        walletRecord.QueryModelField_Amount,
}

var RedisQueryFieldMap map[walletRecord.QueryModelField]string = map[walletRecord.QueryModelField]string{
	walletRecord.QueryModelField_IP:           "ip",
	walletRecord.QueryModelField_Account:      "account",
	walletRecord.QueryModelField_Downlines:    "downlines",
	walletRecord.QueryModelField_OrderBy:      "order_by",
	walletRecord.QueryModelField_Limit:        "limit",
	walletRecord.QueryModelField_Offset:       "offset",
	walletRecord.QueryModelField_LockOrNot:    "lock_or_not",
	walletRecord.QueryModelField_RecordType:   "record_type",
	walletRecord.QueryModelField_Action:       "action",
	walletRecord.QueryModelField_Modifier:     "modifier",
	walletRecord.QueryModelField_DateStart:    "date_start",
	walletRecord.QueryModelField_DateEnd:      "date_end",
	walletRecord.QueryModelField_Actions:      "actions",
	walletRecord.QueryModelField_UserAccounts: "user_accounts",
	walletRecord.QueryModelField_Amount:       "amount",
}

func NewQueryModel() *QueryModel {
	newModel := &QueryModel{
		redisQueryModelImpl: redisQueryModelImpl{
			QueryModelImpl: dao.QueryModelImpl[walletRecord.QueryModelField]{
				Conditions: map[walletRecord.QueryModelField][]interface{}{},
			},
		},
		walletRecordQueryModelImpl: walletRecordQueryModelImpl{
			QueryModelImpl: dao.QueryModelImpl[walletRecord.QueryModelField]{
				Conditions: map[walletRecord.QueryModelField][]interface{}{},
			},
		},
	}
	newModel.redisQueryModelImpl.QueryModel = newModel
	newModel.walletRecordQueryModelImpl.QueryModel = newModel

	return newModel
}

func (q *QueryModel) AddCondition(field walletRecord.QueryModelField, condition ...interface{}) dao.QueryModel[walletRecord.QueryModelField] {
	return q.walletRecordQueryModelImpl.AddCondition(field, condition...)
}

func (q *QueryModel) GetConditions() map[walletRecord.QueryModelField][]interface{} {
	return q.walletRecordQueryModelImpl.Conditions
}

func (c *QueryModel) AddContidionToModel(field walletRecord.QueryModelField, condition ...interface{}) bool {

	switch field {
	case walletRecord.QueryModelField_IP:
		if v, ok := condition[0].(string); ok {
			c.IP = &v
		} else {
			fmt.Printf("AddCondition: wrong input type - %s instead of %s\n",
				reflect.TypeOf(condition[0]).Kind(), reflect.TypeOf(c.IP).Elem().Kind())
		}
	case walletRecord.QueryModelField_Account:
		if v, ok := condition[0].(string); ok {
			c.Account = &v
		} else {
			fmt.Printf("AddCondition: wrong input type - %s instead of %s\n",
				reflect.TypeOf(condition[0]).Kind(), reflect.TypeOf(c.Account).Elem().Kind())
		}
	case walletRecord.QueryModelField_Downlines:
		if v, ok := condition[0].([]string); ok {
			c.Downlines = &v
		} else {
			fmt.Printf("AddCondition: wrong input type - %s instead of %s\n",
				reflect.TypeOf(condition[0]).Kind(), reflect.TypeOf(c.Downlines).Elem().Kind())
		}
	case walletRecord.QueryModelField_OrderBy:
		if v, ok := condition[0].([]string); ok {
			c.OrderBy = &v
		} else {
			fmt.Printf("AddCondition: wrong input type - %s instead of %s\n",
				reflect.TypeOf(condition[0]).Kind(), reflect.TypeOf(c.OrderBy).Elem().Kind())
		}
	case walletRecord.QueryModelField_Limit:
		if v, ok := condition[0].(int); ok {
			c.Limit = &v
		} else if v, ok := condition[0].(int32); ok {
			vInt := int(v)
			c.Limit = &vInt
		} else if v, ok := condition[0].(int64); ok {
			vInt := int(v)
			c.Limit = &vInt
		} else {
			fmt.Printf("AddCondition: wrong input type - %s instead of %s\n",
				reflect.TypeOf(condition[0]).Kind(), reflect.TypeOf(c.Limit).Elem().Kind())
		}
	case walletRecord.QueryModelField_Offset:
		if v, ok := condition[0].(int); ok {
			c.Offset = &v
		} else if v, ok := condition[0].(int32); ok {
			vInt := int(v)
			c.Offset = &vInt
		} else if v, ok := condition[0].(int64); ok {
			vInt := int(v)
			c.Offset = &vInt
		} else {
			fmt.Printf("AddCondition: wrong input type - %s instead of %s\n",
				reflect.TypeOf(condition[0]).Kind(), reflect.TypeOf(c.Offset).Elem().Kind())
		}
	case walletRecord.QueryModelField_LockOrNot:
		if v, ok := condition[0].(bool); ok {
			c.LockOrNot = &v
		} else {
			fmt.Printf("AddCondition: wrong input type - %s instead of %s\n",
				reflect.TypeOf(condition[0]).Kind(), reflect.TypeOf(c.LockOrNot).Elem().Kind())
		}
	case walletRecord.QueryModelField_RecordType:
		if v, ok := condition[0].(int); ok {
			c.RecordType = &v
		} else if v, ok := condition[0].(int32); ok {
			vInt := int(v)
			c.RecordType = &vInt
		} else if v, ok := condition[0].(int64); ok {
			vInt := int(v)
			c.RecordType = &vInt
		} else {
			fmt.Printf("AddCondition: wrong input type - %s instead of %s\n",
				reflect.TypeOf(condition[0]).Kind(), reflect.TypeOf(c.RecordType).Elem().Kind())
		}
	case walletRecord.QueryModelField_Action:
		if v, ok := condition[0].(string); ok {
			c.Action = &v
		} else {
			fmt.Printf("AddCondition: wrong input type - %s instead of %s\n",
				reflect.TypeOf(condition[0]).Kind(), reflect.TypeOf(c.Action).Elem().Kind())
		}
	case walletRecord.QueryModelField_Modifier:
		if v, ok := condition[0].(string); ok {
			c.Modifier = &v
		} else {
			fmt.Printf("AddCondition: wrong input type - %s instead of %s\n",
				reflect.TypeOf(condition[0]).Kind(), reflect.TypeOf(c.Modifier).Elem().Kind())
		}
	case walletRecord.QueryModelField_DateStart:
		if v, ok := condition[0].(string); ok {
			c.DateStart = &v
		} else {
			fmt.Printf("AddCondition: wrong input type - %s instead of %s\n",
				reflect.TypeOf(condition[0]).Kind(), reflect.TypeOf(c.DateStart).Elem().Kind())
		}
	case walletRecord.QueryModelField_DateEnd:
		if v, ok := condition[0].(string); ok {
			c.DateEnd = &v
		} else {
			fmt.Printf("AddCondition: wrong input type - %s instead of %s\n",
				reflect.TypeOf(condition[0]).Kind(), reflect.TypeOf(c.DateEnd).Elem().Kind())
		}
	case walletRecord.QueryModelField_Actions:
		if v, ok := condition[0].([]string); ok {
			c.Actions = &v
		} else {
			fmt.Printf("AddCondition: wrong input type - %s instead of %s\n",
				reflect.TypeOf(condition[0]).Kind(), reflect.TypeOf(c.Actions).Elem().Kind())
		}
	case walletRecord.QueryModelField_UserAccounts:
		if v, ok := condition[0].([]string); ok {
			c.UserAccounts = &v
		} else {
			fmt.Printf("AddCondition: wrong input type - %s instead of %s\n",
				reflect.TypeOf(condition[0]).Kind(), reflect.TypeOf(c.UserAccounts).Elem().Kind())
		}
	case walletRecord.QueryModelField_Amount:
		if v, ok := condition[0].(decimal.Decimal); ok {
			c.Amount = &v
		} else {
			fmt.Printf("AddCondition: wrong input type - %s instead of %s\n",
				reflect.TypeOf(condition[0]).Kind(), reflect.TypeOf(c.Amount).Elem().Kind())
		}
	}
	return true
}

func (q *QueryModel) Set(query dao.QueryModel[walletRecord.QueryModelField]) dao.QueryModel[walletRecord.QueryModelField] {
	return q.walletRecordQueryModelImpl.Set(query)
}

func (c *QueryModel) HasRequiredQuery() bool {
	// TODO: use reflect to refactor and put in redis.querymodel
	if c.Account == nil {
		return false
	}
	return true
}

func (c *QueryModel) GetMaskedQueryModel() redis.QueryModel[walletRecord.QueryModelField] {
	// TODO: use reflect to refactor and put in redis.querymodel
	clone := NewQueryModel()
	switch {
	case c.IP != nil:
		clone.AddCondition(walletRecord.QueryModelField_IP, *c.IP)
	case c.Account != nil:
		clone.AddCondition(walletRecord.QueryModelField_Account, *c.Account)
	case c.Downlines != nil:
		clone.AddCondition(walletRecord.QueryModelField_Downlines, *c.Downlines)
	case c.OrderBy != nil:
		clone.AddCondition(walletRecord.QueryModelField_OrderBy, *c.OrderBy)
	case c.Limit != nil:
		clone.AddCondition(walletRecord.QueryModelField_Limit, *c.Limit)
	case c.Offset != nil:
		clone.AddCondition(walletRecord.QueryModelField_Offset, *c.Offset)
	// case c.LockOrNot != nil:	clone.AddCondition(walletRecord.QueryModelField_LockOrNot, *c.LockOrNot)
	case c.RecordType != nil:
		clone.AddCondition(walletRecord.QueryModelField_RecordType, *c.RecordType)
	case c.Action != nil:
		clone.AddCondition(walletRecord.QueryModelField_Action, *c.Action)
	case c.Modifier != nil:
		clone.AddCondition(walletRecord.QueryModelField_Modifier, *c.Modifier)
	// case c.DateStart != nil:	clone.AddCondition(walletRecord.QueryModelField_DateStart, *c.DateStart)
	// case c.DateEnd != nil:	clone.AddCondition(walletRecord.QueryModelField_DateEnd, *c.DateEnd)
	case c.Actions != nil:
		clone.AddCondition(walletRecord.QueryModelField_Actions, *c.Actions)
	case c.UserAccounts != nil:
		clone.AddCondition(walletRecord.QueryModelField_UserAccounts, *c.UserAccounts)
	case c.Amount != nil:
		clone.AddCondition(walletRecord.QueryModelField_Amount, *c.Amount)
	}

	return clone
}

func (q *QueryModel) GetTagFromField(field walletRecord.QueryModelField) string {
	tag, ok := RedisQueryFieldMap[field]
	if !ok {
		message := fmt.Sprintf("GetTagFromField: no such field %s.", reflect.TypeOf(field).Name())
		panic(message)
	}
	return tag
}

func (q *QueryModel) GetFieldFromTag(tag string) walletRecord.QueryModelField {
	field, ok := QueryFieldRedisMap[tag]
	if !ok {
		message := fmt.Sprintf("GetFieldFromTag: no such tag %s.", tag)
		panic(message)
	}
	return field
}

func getTagFromField(field walletRecord.QueryModelField) string {
	tag, ok := RedisQueryFieldMap[field]
	if !ok {
		message := fmt.Sprintf("GetTagFromField: no such field %s.", reflect.TypeOf(field).Name())
		panic(message)
	}
	return tag
}

func getFieldFromTag(tag string) walletRecord.QueryModelField {
	field, ok := QueryFieldRedisMap[tag]
	if !ok {
		message := fmt.Sprintf("GetFieldFromTag: no such tag %s.", tag)
		panic(message)
	}
	return field
}
