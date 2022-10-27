package walletRecord

import (
	"context"
	"fmt"
	"reflect"

	goCompositeDao "github.com/lisyaoran51/GoCompositeDao"
	"github.com/lisyaoran51/GoCompositeDao/dao"
	models "github.com/lisyaoran51/GoCompositeDao/models/databaseModels"
)

type UpdateField int

const (
	UpdateField_None UpdateField = iota
	UpdateField_ID
	UpdateField_RecordID
	UpdateField_Account
	UpdateField_Action
	UpdateField_Amount
	UpdateField_WalletAmount
	UpdateField_WalletType
	UpdateField_Ip
	UpdateField_Description
	UpdateField_Memo
	UpdateField_CreatedAt
	UpdateField_UpdatedAt
	UpdateField_CreatedBy
	UpdateField_UpdatedBy
)

var daos = map[string]func() Dao{}

type Dao interface {
	dao.Dao[models.WalletRecordModel, QueryModelField, UpdateField]
}

type DaoImpl struct {
	dao.Dao[models.WalletRecordModel, QueryModelField, UpdateField]
}

func GetDao(ctx context.Context, source interface{}) Dao {

	sourceType := reflect.TypeOf(source).String()

	if _, ok := daos[sourceType]; !ok {
		fmt.Printf("GetDao: no such dao. [%s]\n", sourceType)
		return nil
	}
	return daos[sourceType]()
}

func (c *DaoImpl) Register(ctx context.Context) error {
	if _, ok := c.Dao.CreateDao(ctx).(Dao); !ok {
		return goCompositeDao.ErrInternal
	}

	daos[c.Dao.GetSourceType()] = func() Dao {
		d := c.Dao.CreateDao(ctx)
		dao := d.(Dao)
		return dao
	}

	for k, v := range daos {
		fmt.Printf("daos [%s] [*%s] initialized.\n", k, reflect.TypeOf(v()).Elem().Name())
	}
	return nil
}

func (c *DaoImpl) GetModelType() string {
	model := &models.WalletRecordModel{}
	return reflect.TypeOf(model).String()
}

func (c *DaoImpl) GetDataName() string {
	return "wallet_record"
}
