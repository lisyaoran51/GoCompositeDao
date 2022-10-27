package initializer

import (
	"context"

	gormWalletRecord "github.com/lisyaoran51/GoCompositeDao/dao/walletRecord/gorm"
	redisWalletRecord "github.com/lisyaoran51/GoCompositeDao/dao/walletRecord/redis"
)

func Initialize(ctx context.Context) {
	redisWalletRecord.NewDao().Register(ctx)
	gormWalletRecord.NewDao().Register(ctx)
}
