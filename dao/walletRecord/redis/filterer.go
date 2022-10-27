package redis

import (
	"time"

	"github.com/lisyaoran51/GoCompositeDao/dao/redis"
	databaseModels "github.com/lisyaoran51/GoCompositeDao/models/databaseModels"
)

type Filterer struct {
	redis.FiltererImpl[databaseModels.WalletRecordModel, QueryModel]
}

func NewFilterer(q *QueryModel) *Filterer {
	f := &Filterer{}
	f.SetQueryModel(q)
	return f
}

func (f *Filterer) filterDateStart() func(models []*databaseModels.WalletRecordModel) []*databaseModels.WalletRecordModel {
	if f.GetQueryModel().DateStart == nil {
		return func(models []*databaseModels.WalletRecordModel) []*databaseModels.WalletRecordModel {
			return models
		}
	}

	return func(models []*databaseModels.WalletRecordModel) []*databaseModels.WalletRecordModel {
		newModels := make([]*databaseModels.WalletRecordModel, 0)

		dateStart, err := time.Parse("2006-01-02 15:04:05", *f.GetQueryModel().DateStart)
		if err != nil {
			panic(err)
		}

		for _, m := range models {
			if !m.CreatedAt.Before(dateStart) {
				newModels = append(newModels, m)
			}
		}

		return newModels
	}
}

func (f *Filterer) filterDateEnd() func(models []*databaseModels.WalletRecordModel) []*databaseModels.WalletRecordModel {
	if f.GetQueryModel().DateEnd == nil {
		return func(models []*databaseModels.WalletRecordModel) []*databaseModels.WalletRecordModel {
			return models
		}
	}

	return func(models []*databaseModels.WalletRecordModel) []*databaseModels.WalletRecordModel {
		newModels := make([]*databaseModels.WalletRecordModel, 0)

		dateEnd, err := time.Parse("2006-01-02 15:04:05", *f.GetQueryModel().DateEnd)
		if err != nil {
			panic(err)
		}

		for _, m := range models {
			if !m.CreatedAt.After(dateEnd) {
				newModels = append(newModels, m)
			}
		}

		return newModels
	}
}

func (f *Filterer) sortOffset() func(models []*databaseModels.WalletRecordModel) []*databaseModels.WalletRecordModel {
	if f.GetQueryModel().Offset == nil {
		return func(models []*databaseModels.WalletRecordModel) []*databaseModels.WalletRecordModel {
			return models
		}
	}

	return func(models []*databaseModels.WalletRecordModel) []*databaseModels.WalletRecordModel {
		if len(models) == 0 {
			return models
		}

		if len(models)-1 < *f.GetQueryModel().Offset {
			return make([]*databaseModels.WalletRecordModel, 0)
		}
		return models[*f.GetQueryModel().Offset:]
	}
}

func (f *Filterer) sortLimit() func(models []*databaseModels.WalletRecordModel) []*databaseModels.WalletRecordModel {
	if f.GetQueryModel().Limit == nil {
		return func(models []*databaseModels.WalletRecordModel) []*databaseModels.WalletRecordModel {
			return models
		}
	}

	return func(models []*databaseModels.WalletRecordModel) []*databaseModels.WalletRecordModel {
		if len(models) <= *f.GetQueryModel().Limit {
			return models
		}

		return models[:*f.GetQueryModel().Limit]
	}
}
