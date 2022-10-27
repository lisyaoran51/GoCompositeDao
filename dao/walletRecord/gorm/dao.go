package gorm

import (
	"context"
	"database/sql"
	"fmt"

	gormLib "github.com/jinzhu/gorm"
	goCompositeDao "github.com/lisyaoran51/GoCompositeDao"
	"github.com/lisyaoran51/GoCompositeDao/dao"
	"github.com/lisyaoran51/GoCompositeDao/dao/gorm"
	"github.com/lisyaoran51/GoCompositeDao/dao/specification"
	"github.com/lisyaoran51/GoCompositeDao/dao/walletRecord"
	models "github.com/lisyaoran51/GoCompositeDao/models/databaseModels"
)

type Dao struct {
	walletRecord.DaoImpl
	gorm.Dao
}

func NewDao() *Dao {
	newDao := &Dao{}
	newDao.DaoImpl.Dao = newDao
	return newDao
}

func (c *Dao) GetTable(db *gormLib.DB) (*gormLib.DB, error) {
	return db.Table(c.GetDataName()), nil
}

func (c *Dao) new(ctx context.Context, db *gormLib.DB, model *models.WalletRecordModel) (interface{}, error) {

	db, _ = c.GetTable(db)

	newModel := &models.WalletRecordNewModel{
		ID:           model.ID,
		RecordID:     model.RecordID,
		Account:      model.Account,
		Action:       model.Action,
		Amount:       model.Amount,
		WalletAmount: model.WalletAmount,
		WalletType:   model.WalletType,
		Ip:           model.Ip,
		Description:  model.Description,
		Memo:         model.Memo,
		CreatedAt:    model.CreatedAt,
		UpdatedAt:    model.UpdatedAt,
		CreatedBy:    model.CreatedBy,
		UpdatedBy:    model.UpdatedBy,
	}

	err := db.Create(newModel).Error
	if err != nil {
		return nil, err
	}

	model.ID = newModel.ID
	model.RecordID = newModel.RecordID
	model.Account = newModel.Account
	model.Action = newModel.Action
	model.Amount = newModel.Amount
	model.WalletAmount = newModel.WalletAmount
	model.WalletType = newModel.WalletType
	model.Ip = newModel.Ip
	model.Description = newModel.Description
	model.Memo = newModel.Memo
	model.CreatedAt = newModel.CreatedAt
	model.UpdatedAt = newModel.UpdatedAt
	model.CreatedBy = newModel.CreatedBy
	model.UpdatedBy = newModel.UpdatedBy

	return model, nil
}

func (c *Dao) count(ctx context.Context, db *gormLib.DB, query *QueryModel) (int, error) {

	db, _ = c.GetTable(db)

	var count int32 = 0

	err := db.
		Scopes(query.GetCountChain()).
		Count(&count).Error

	if gormLib.IsRecordNotFoundError(err) {
		return 0, nil
	}

	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func (c *Dao) get(ctx context.Context, db *gormLib.DB, query *QueryModel) (*models.WalletRecordModel, error) {

	db, _ = c.GetTable(db)

	var model models.WalletRecordModel

	err := db.
		Select(c.GetDataName() + ".*, wallet.currency").
		Joins("JOIN wallet ON wallet.account = " + c.GetDataName() + ".account").
		Scopes(query.GetQueryChain()).
		Scan(&model).Error

	if gormLib.IsRecordNotFoundError(err) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}
	return &model, nil
}

func (c *Dao) gets(ctx context.Context, db *gormLib.DB, query *QueryModel) ([]*models.WalletRecordModel, error) {
	db, _ = c.GetTable(db)

	var records []*models.WalletRecordModel

	err := db.
		Select(c.GetDataName() + ".*, wallet.currency").
		Joins("JOIN wallet ON wallet.account = " + c.GetDataName() + ".account").
		Scopes(query.GetQueryChain()).
		Scan(&records).Error

	if gormLib.IsRecordNotFoundError(err) {
		return []*models.WalletRecordModel{}, nil
	}

	if err != nil {
		return nil, err
	}
	return records, nil
}

func (c *Dao) getsWithPagination(ctx context.Context, db *gormLib.DB, query *QueryModel, paginate *specification.PaginationStruct) (_ []*models.WalletRecordModel, _ int, err error) {
	db, _ = c.GetTable(db)
	db.LogMode(true)
	var records []*models.WalletRecordModel
	var count int32 = 0

	paginate.TableName = c.GetDataName()

	err = db.
		Select(c.GetDataName() + ".*, wallet.currency").
		Scopes(query.GetQueryChain()).
		Count(&count).
		Joins("JOIN wallet ON wallet.account = wallet_record.account AND wallet.wallet_type = wallet_record.wallet_type").
		Scopes(specification.NewPaginationSpecification(paginate)).
		Scan(&records).Error

	if gormLib.IsRecordNotFoundError(err) {
		return []*models.WalletRecordModel{}, 0, nil
	}

	if err != nil {
		return nil, 0, err
	}
	return records, int(count), nil
}

func (c *Dao) modify(ctx context.Context, db *gormLib.DB, model *models.WalletRecordModel, fields []walletRecord.UpdateField) error {

	db, _ = c.GetTable(db)

	if len(fields) == 0 {
		newModel := &models.WalletRecordNewModel{
			ID:           model.ID,
			RecordID:     model.RecordID,
			Account:      model.Account,
			Action:       model.Action,
			Amount:       model.Amount,
			WalletAmount: model.WalletAmount,
			WalletType:   model.WalletType,
			Ip:           model.Ip,
			Description:  model.Description,
			Memo:         model.Memo,
			CreatedAt:    model.CreatedAt,
			UpdatedAt:    model.UpdatedAt,
			CreatedBy:    model.CreatedBy,
			UpdatedBy:    model.UpdatedBy,
		}
		err := db.Save(newModel).Error
		return err
	}

	attrs := map[string]interface{}{}

	for _, f := range fields {
		switch f {
		case walletRecord.UpdateField_Account:
			attrs["account"] = model.Account
		}
	}

	err := db.
		Where("exchange_id = ?", model.ID).
		Updates(attrs).Error

	return err
}

func (c *Dao) delete(ctx context.Context, db *gormLib.DB, query *QueryModel) error {

	db, _ = c.GetTable(db)

	err := db.
		Scopes(query.GetDeleteChain()).
		Delete(&models.WalletRecordNewModel{}).Error

	return err
}

/// ======================== REGION ========================
/// below is only for inheritance assertion. no business logic here
/// ========================================================

/// ============= dao.Dao =============

func (c *Dao) CreateDao(ctx context.Context) dao.Dao[models.WalletRecordModel, walletRecord.QueryModelField, walletRecord.UpdateField] {
	return NewDao()
}

/// ============= walletRecord.Dao =============

func (c *Dao) New(ctx context.Context, dataSource interface{}, model *models.WalletRecordModel) (result interface{}, err error) {
	d, ok := dataSource.(*gormLib.DB)
	if !ok {
		return nil, goCompositeDao.ErrInternal
	}

	// bind db with context
	if tx := d.BeginTx(ctx, &sql.TxOptions{}); tx.Error == nil {
		defer func() {
			if err != nil {
				if rollbackErr := tx.Rollback().Error; rollbackErr != nil {
					fmt.Printf("failed to rollback %s\n", rollbackErr)
				}
				return
			}
			err = tx.Commit().Error
		}()

		return c.new(ctx, tx, model)
	}

	// nested transaction happened. unable to set context
	return c.new(ctx, d, model)
}

func (c *Dao) Count(ctx context.Context, dataSource interface{}, query dao.QueryModel[walletRecord.QueryModelField]) (_ int, err error) {
	d, ok := dataSource.(*gormLib.DB)
	if !ok {
		return 0, goCompositeDao.ErrInternal
	}

	q, ok := query.(*QueryModel)
	if !ok {
		return 0, goCompositeDao.ErrInternal
	}

	// bind db with context
	if tx := d.BeginTx(ctx, &sql.TxOptions{}); tx.Error == nil {
		defer func() {
			if err != nil {
				if rollbackErr := tx.Rollback().Error; rollbackErr != nil {
					fmt.Printf("failed to rollback %s\n", rollbackErr)
				}
				return
			}
			err = tx.Commit().Error
		}()

		return c.count(ctx, tx, q)
	}

	// nested transaction happened. unable to set context
	return c.count(ctx, d, q)
}

func (c *Dao) Get(ctx context.Context, dataSource interface{}, query dao.QueryModel[walletRecord.QueryModelField]) (result *models.WalletRecordModel, err error) {
	d, ok := dataSource.(*gormLib.DB)
	if !ok {
		return nil, goCompositeDao.ErrInternal
	}

	q, ok := query.(*QueryModel)
	if !ok {
		return nil, goCompositeDao.ErrInternal
	}

	// bind db with context
	if tx := d.BeginTx(ctx, &sql.TxOptions{}); tx.Error == nil {
		defer func() {
			if err != nil {
				if rollbackErr := tx.Rollback().Error; rollbackErr != nil {
					fmt.Printf("failed to rollback %s\n", rollbackErr)
				}
				return
			}
			err = tx.Commit().Error
		}()

		return c.get(ctx, tx, q)
	}

	// nested transaction happened. unable to set context
	return c.get(ctx, d, q)
}

func (c *Dao) Gets(ctx context.Context, dataSource interface{}, query dao.QueryModel[walletRecord.QueryModelField]) (result []*models.WalletRecordModel, err error) {

	d, ok := dataSource.(*gormLib.DB)
	if !ok {
		fmt.Printf("Dao Gets casting failed.\n")
		return nil, goCompositeDao.ErrInternal
	}

	q, ok := query.(*QueryModel)
	if !ok {
		fmt.Printf("Dao Gets QueryModel casting failed.\n")
		return nil, goCompositeDao.ErrInternal
	}

	// bind db with context
	if tx := d.BeginTx(ctx, &sql.TxOptions{}); tx.Error == nil {
		defer func() {
			if err != nil {
				if rollbackErr := tx.Rollback().Error; rollbackErr != nil {
					fmt.Printf("failed to rollback %s\n", rollbackErr)
				}
				return
			}
			err = tx.Commit().Error

		}()

		return c.gets(ctx, tx, q)
	}

	// nested transaction happened. unable to set context
	return c.gets(ctx, d, q)
}

func (c *Dao) GetsWithPagination(ctx context.Context, dataSource interface{}, query dao.QueryModel[walletRecord.QueryModelField], paginate *specification.PaginationStruct) (_ []*models.WalletRecordModel, _ int, err error) {

	d, ok := dataSource.(*gormLib.DB)
	if !ok {
		fmt.Printf("Dao Gets casting failed.\n")
		return nil, 0, goCompositeDao.ErrInternal
	}

	q, ok := query.(*QueryModel)
	if !ok {
		fmt.Printf("Dao Gets QueryModel casting failed.\n")
		return nil, 0, goCompositeDao.ErrInternal
	}

	// bind db with context
	if tx := d.BeginTx(ctx, &sql.TxOptions{}); tx.Error == nil {
		defer func() {
			if err != nil {
				if rollbackErr := tx.Rollback().Error; rollbackErr != nil {
					fmt.Printf("failed to rollback %s\n", rollbackErr)
				}
				return
			}
			err = tx.Commit().Error

		}()

		return c.getsWithPagination(ctx, tx, q, paginate)
	}

	// nested transaction happened. unable to set context
	return c.getsWithPagination(ctx, d, q, paginate)
}

func (c *Dao) Modify(ctx context.Context, dataSource interface{}, model *models.WalletRecordModel, fields []walletRecord.UpdateField) (err error) {
	d, ok := dataSource.(*gormLib.DB)
	if !ok {
		return goCompositeDao.ErrInternal
	}

	// bind db with context
	if tx := d.BeginTx(ctx, &sql.TxOptions{}); tx.Error == nil {
		defer func() {
			if err != nil {
				if rollbackErr := tx.Rollback().Error; rollbackErr != nil {
					fmt.Printf("failed to rollback %s\n", rollbackErr)
				}
				return
			}
			err = tx.Commit().Error
		}()

		return c.modify(ctx, tx, model, fields)
	}

	// nested transaction happened. unable to set context
	return c.modify(ctx, d, model, fields)
}

func (c *Dao) Delete(ctx context.Context, dataSource interface{}, query dao.QueryModel[walletRecord.QueryModelField]) (err error) {
	d, ok := dataSource.(*gormLib.DB)
	if !ok {
		return goCompositeDao.ErrInternal
	}

	q, ok := query.(*QueryModel)
	if !ok {
		return goCompositeDao.ErrInternal
	}

	// bind db with context
	if tx := d.BeginTx(ctx, &sql.TxOptions{}); tx.Error == nil {
		defer func() {
			if err != nil {
				if rollbackErr := tx.Rollback().Error; rollbackErr != nil {
					fmt.Printf("failed to rollback %s\n", rollbackErr)
				}
				return
			}
			err = tx.Commit().Error
		}()

		return c.delete(ctx, tx, q)
	}

	// nested transaction happened. unable to set context
	return c.delete(ctx, d, q)
}

func (c *Dao) GetQueryModel() dao.QueryModel[walletRecord.QueryModelField] {
	queryModel := NewQueryModel()
	queryModel.DataName = c.GetDataName()
	return queryModel
}
