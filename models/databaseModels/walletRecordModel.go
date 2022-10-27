package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type WalletRecordNewModel struct {
	ID           uint64          `gorm:"column:id; primary_key"`
	RecordID     string          `gorm:"column:record_id"`
	Account      string          `gorm:"column:account"`
	Action       string          `gorm:"column:action"`
	Amount       decimal.Decimal `gorm:"column:amount"`
	WalletAmount decimal.Decimal `gorm:"column:wallet_amount"`
	WalletType   int32           `gorm:"column:wallet_type; default:3"`
	Ip           string          `gorm:"column:ip"`
	Description  string          `gorm:"column:description"`
	Memo         string          `gorm:"column:memo"`
	CreatedAt    time.Time       `gorm:"column:created_at"`
	UpdatedAt    time.Time       `gorm:"column:updated_at"`
	CreatedBy    string          `gorm:"column:created_by"`
	UpdatedBy    string          `gorm:"column:updated_by"`
}

// WalletRecordModel is raw-data-form exactly as db data
type WalletRecordModel struct {
	ID           uint64          `gorm:"column:id; primary_key"`
	RecordID     string          `gorm:"column:record_id"`
	Account      string          `gorm:"column:account"`
	Action       string          `gorm:"column:action"`
	Amount       decimal.Decimal `gorm:"column:amount"`
	WalletAmount decimal.Decimal `gorm:"column:wallet_amount"`
	WalletType   int32           `gorm:"column:wallet_type; default:3"`
	Ip           string          `gorm:"column:ip"`
	Description  string          `gorm:"column:description"`
	Memo         string          `gorm:"column:memo"`
	CreatedAt    time.Time       `gorm:"column:created_at"`
	UpdatedAt    time.Time       `gorm:"column:updated_at"`
	CreatedBy    string          `gorm:"column:created_by"`
	Currency     string          `gorm:"column:currency"`
	UpdatedBy    string          `gorm:"column:updated_by"`
}
