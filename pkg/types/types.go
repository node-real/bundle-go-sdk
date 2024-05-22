package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Hashes []common.Hash

type BundleStatus uint8

const (
	BundleStatusPending BundleStatus = iota
	BundleStatusConfirmed
	BundleStatusFailed
)

type Bundle struct {
	Hash                 common.Hash
	Txs                  Hashes         `gorm:"column:txs;type:blob"`
	MaxBlockNumber       uint64         `gorm:"column:max_block_number"`
	MaxTimestamp         uint64         `gorm:"column:max_timestamp"`
	Status               BundleStatus   `gorm:"column:status;default:0"`
	GasFee               *hexutil.Big   `gorm:"column:gas_fee;type:VARBINARY(32)"`
	Builder              common.Address `gorm:"column:builder;type:BINARY(20)"`
	ConfirmedBlockNumber uint64         `gorm:"column:confirmed_block_number"`
	ConfirmedDate        uint64         `gorm:"column:confirmed_date"`
}
