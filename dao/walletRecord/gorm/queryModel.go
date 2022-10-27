package gorm

import (
	"fmt"
	"reflect"

	"github.com/jinzhu/gorm"
	gormLib "github.com/jinzhu/gorm"
	"github.com/lisyaoran51/GoCompositeDao/dao"
	"github.com/lisyaoran51/GoCompositeDao/dao/walletRecord"
	"github.com/shopspring/decimal"
)

type QueryModel struct {
	walletRecord.QueryModelImpl
	DataName     string
	IP           *string
	Account      *string
	Downlines    *[]string
	OrderBy      *[]string
	Limit        *int
	Offset       *int
	LockOrNot    *bool
	RecordType   *int
	Action       *string
	Modifier     *string
	DateStart    *string
	DateEnd      *string
	Actions      *[]string
	UserAccounts *[]string
	Amount       *decimal.Decimal
}

func NewQueryModel() *QueryModel {
	newModel := &QueryModel{
		QueryModelImpl: walletRecord.QueryModelImpl{
			QueryModelImpl: dao.QueryModelImpl[walletRecord.QueryModelField]{
				Conditions: map[walletRecord.QueryModelField][]interface{}{},
			},
		},
	}
	newModel.QueryModelImpl.QueryModel = newModel
	return newModel
}

func (c *QueryModel) AddContidionToModel(field walletRecord.QueryModelField, condition ...interface{}) bool {

	switch field {
	case walletRecord.QueryModelField_IP:
		if v, ok := condition[0].(string); ok {
			c.IP = &v
		} else {
			fmt.Printf("AddCondition: wrong input type - %s instead of %s\n",
				reflect.TypeOf(condition[0]).Name(), reflect.TypeOf(c.IP).Elem().Kind())
		}
	case walletRecord.QueryModelField_Account:
		if v, ok := condition[0].(string); ok {
			c.Account = &v
		} else {
			fmt.Printf("AddCondition: wrong input type - %s instead of %s\n",
				reflect.TypeOf(condition[0]).Name(), reflect.TypeOf(c.Account).Elem().Kind())
		}
	case walletRecord.QueryModelField_Downlines:
		if v, ok := condition[0].([]string); ok {
			c.Downlines = &v
		} else {
			fmt.Printf("AddCondition: wrong input type - %s instead of %s\n",
				reflect.TypeOf(condition[0]).Name(), reflect.TypeOf(c.Downlines).Elem().Kind())
		}
	case walletRecord.QueryModelField_OrderBy:
		if v, ok := condition[0].([]string); ok {
			c.OrderBy = &v
		} else {
			fmt.Printf("AddCondition: wrong input type - %s instead of %s\n",
				reflect.TypeOf(condition[0]).Name(), reflect.TypeOf(c.OrderBy).Elem().Kind())
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
				reflect.TypeOf(condition[0]).Name(), reflect.TypeOf(c.Limit).Elem().Kind())
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
				reflect.TypeOf(condition[0]).Name(), reflect.TypeOf(c.Offset).Elem().Kind())
		}
	case walletRecord.QueryModelField_LockOrNot:
		if v, ok := condition[0].(bool); ok {
			c.LockOrNot = &v
		} else {
			fmt.Printf("AddCondition: wrong input type - %s instead of %s\n",
				reflect.TypeOf(condition[0]).Name(), reflect.TypeOf(c.LockOrNot).Elem().Kind())
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
				reflect.TypeOf(condition[0]).Name(), reflect.TypeOf(c.RecordType).Elem().Kind())
		}
	case walletRecord.QueryModelField_Action:
		if v, ok := condition[0].(string); ok {
			c.Action = &v
		} else {
			fmt.Printf("AddCondition: wrong input type - %s instead of %s\n",
				reflect.TypeOf(condition[0]).Name(), reflect.TypeOf(c.Action).Elem().Kind())
		}
	case walletRecord.QueryModelField_Modifier:
		if v, ok := condition[0].(string); ok {
			c.Modifier = &v
		} else {
			fmt.Printf("AddCondition: wrong input type - %s instead of %s\n",
				reflect.TypeOf(condition[0]).Name(), reflect.TypeOf(c.Modifier).Elem().Kind())
		}
	case walletRecord.QueryModelField_DateStart:
		if v, ok := condition[0].(string); ok {
			c.DateStart = &v
		} else {
			fmt.Printf("AddCondition: wrong input type - %s instead of %s\n",
				reflect.TypeOf(condition[0]).Name(), reflect.TypeOf(c.DateStart).Elem().Kind())
		}
	case walletRecord.QueryModelField_DateEnd:
		if v, ok := condition[0].(string); ok {
			c.DateEnd = &v
		} else {
			fmt.Printf("AddCondition: wrong input type - %s instead of %s\n",
				reflect.TypeOf(condition[0]).Name(), reflect.TypeOf(c.DateEnd).Elem().Kind())
		}
	case walletRecord.QueryModelField_Actions:
		if v, ok := condition[0].([]string); ok {
			c.Actions = &v
		} else {
			fmt.Printf("AddCondition: wrong input type - %s instead of %s\n",
				reflect.TypeOf(condition[0]).Name(), reflect.TypeOf(c.Actions).Elem().Kind())
		}
	case walletRecord.QueryModelField_UserAccounts:
		if v, ok := condition[0].([]string); ok {
			c.UserAccounts = &v
		} else {
			fmt.Printf("AddCondition: wrong input type - %s instead of %s\n",
				reflect.TypeOf(condition[0]).Name(), reflect.TypeOf(c.UserAccounts).Elem().Kind())
		}
	case walletRecord.QueryModelField_Amount:
		if v, ok := condition[0].(decimal.Decimal); ok {
			c.Amount = &v
		} else {
			fmt.Printf("AddCondition: wrong input type - %s instead of %s\n",
				reflect.TypeOf(condition[0]).Name(), reflect.TypeOf(c.Amount).Elem().Kind())
		}
	}
	return true
}

func (c *QueryModel) GetCountChain() func(db *gormLib.DB) *gormLib.DB {
	return func(db *gormLib.DB) *gormLib.DB {
		return db.
			Scopes(c.ipEqualScope()).
			Scopes(c.accountEqualScope()).
			Scopes(c.downlinesEqualScope()).
			Scopes(c.accountsInScope()).
			Scopes(c.actionEqualScope()).
			Scopes(c.recordTypeEqualScope()).
			Scopes(c.lockScope()).
			Scopes(c.createdAtBetweenScope()).
			Scopes(c.actionArrayINScope()).
			Scopes(c.modifierEqualScope()).
			Scopes(c.amountEqualScope())
	}
}

func (c *QueryModel) GetQueryChain() func(db *gormLib.DB) *gormLib.DB {
	return func(db *gormLib.DB) *gormLib.DB {
		return db.
			Scopes(c.ipEqualScope()).
			Scopes(c.accountEqualScope()).
			Scopes(c.downlinesEqualScope()).
			Scopes(c.accountsInScope()).
			Scopes(c.actionEqualScope()).
			Scopes(c.recordTypeEqualScope()).
			Scopes(c.orderByScope()).
			Scopes(c.lockScope()).
			Scopes(c.createdAtBetweenScope()).
			Scopes(c.actionArrayINScope()).
			Scopes(c.limitOffsetScope()).
			Scopes(c.modifierEqualScope()).
			Scopes(c.amountEqualScope())
	}
}

func (c *QueryModel) GetDeleteChain() func(db *gormLib.DB) *gormLib.DB {
	return func(db *gormLib.DB) *gormLib.DB {
		return db.
			Scopes(c.ipEqualScope()).
			Scopes(c.accountEqualScope()).
			Scopes(c.downlinesEqualScope()).
			Scopes(c.accountsInScope()).
			Scopes(c.actionEqualScope()).
			Scopes(c.recordTypeEqualScope()).
			Scopes(c.orderByScope()).
			Scopes(c.lockScope()).
			Scopes(c.createdAtBetweenScope()).
			Scopes(c.actionArrayINScope()).
			Scopes(c.limitOffsetScope()).
			Scopes(c.modifierEqualScope()).
			Scopes(c.amountEqualScope())
	}
}

func (c *QueryModel) ipEqualScope() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if c.IP == nil {
			return db
		}

		return db.Where(c.DataName+".ip = ?", *c.IP)
	}
}

func (c *QueryModel) actionEqualScope() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if c.Action == nil {
			return db
		}

		return db.Where(c.DataName+".action = ?", *c.Action)
	}
}

func (c *QueryModel) recordTypeEqualScope() func(db *gorm.DB) *gorm.DB {
	action := []string{"deposit", "manual", "withdraw"}
	return func(db *gorm.DB) *gorm.DB {
		if c.RecordType == nil {
			return db
		}

		switch *c.RecordType {
		case 1:
			return db.Where(c.DataName+".amount >= 0 AND "+c.DataName+".action IN (?)", action)
		case 2:
			return db.Where(c.DataName+".amount < 0 AND "+c.DataName+".action IN (?)", action)
		case 3:
			return db.Where(c.DataName+".action IN (?)", action)
		case 4:
			return db.Where(c.DataName+".action NOT IN (?)", action)
		}
		return db
	}
}

func (c *QueryModel) downlinesEqualScope() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if c.Downlines == nil {
			return db
		}
		if len(*c.Downlines) > 0 {
			return db.Where(c.DataName+".account IN (?)", *c.Downlines)
		} else {
			return db
		}
	}
}

func (c *QueryModel) accountEqualScope() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if c.Account == nil {
			return db
		}

		return db.Where(c.DataName+".account = ?", *c.Account)
	}
}

func (c *QueryModel) modifierEqualScope() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if c.Modifier == nil {
			return db
		}

		return db.Where(c.DataName+".updated_by = ? OR "+c.DataName+".created_by = ?", *c.Modifier, *c.Modifier)
	}
}

func (c *QueryModel) accountsInScope() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if c.UserAccounts == nil {
			return db
		}

		if len(*c.UserAccounts) > 0 {
			return db.Where(c.DataName+".account IN (?)", *c.UserAccounts)
		}
		return db
	}
}

func (c *QueryModel) actionArrayINScope() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if c.Actions == nil {
			return db
		}
		if len(*c.Actions) > 0 {
			return db.Where(c.DataName+".action IN (?)", *c.Actions)
		}
		return db
	}
}

func (c *QueryModel) createdAtBetweenScope() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if c.DateStart == nil || c.DateEnd == nil {
			return db
		}
		return db.Where(c.DataName+".created_at BETWEEN ? AND ?", *c.DateStart, *c.DateEnd)
	}
}

func (c *QueryModel) orderByScope() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		if c.OrderBy == nil {
			return db
		}
		if len(*c.OrderBy) != 0 {
			order := (*c.OrderBy)[0]
			for _, o := range (*c.OrderBy)[1:] {
				order = order + ", " + o
			}
			return db.Order(order)
		}
		return db
	}
}

func (c *QueryModel) limitOffsetScope() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		if c.Limit == nil {
			return db
		}
		if *c.Limit > 0 {
			if c.Offset == nil {
				return db.Limit(*c.Limit)
			}
			return db.Limit(*c.Limit).Offset(*c.Offset)
		}
		return db
	}
}

func (c *QueryModel) lockScope() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if c.LockOrNot == nil {
			return db
		}
		if *c.LockOrNot {
			return db.Set("gorm:query_option", "FOR UPDATE")
		}
		return db
	}
}

func (c *QueryModel) amountEqualScope() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if c.Amount == nil {
			return db
		}
		if !c.Amount.IsZero() {
			return db.Where(c.DataName+".amount = ?", *c.Amount)
		}
		return db
	}
}
