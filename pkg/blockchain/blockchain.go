package blockchain

import (
	"github.com/shopspring/decimal"
	"mouse/pkg/blockchain/chain"
	"mouse/pkg/blockchain/model"
)

var ChainMap = map[string]ChainInterface{
	BtcChain:     chain.NewBtcChain(),
	TronChain:    chain.NewTronChain(),
	EthChain:     chain.NewEvmChain(),
	BscChain:     chain.NewEvmChain(),
	PolygonChain: chain.NewEvmChain(),
	ArbChain:     chain.NewEvmChain(),
	AvaxChain:    chain.NewEvmChain(),
	BaseChain:    chain.NewEvmChain(),
	SuiChain:     chain.NewSuiChain(),
	InjChain:     chain.NewInjChain(),
	SolChain:     chain.NewSolChain(),
	TonChain:     chain.NewTonChain(),
}

type ChainInterface interface {
	GenAddr() (string, string, error)
	GenHdAddr() (string, string, error)
	GetAddrBalance(string, model.CurrencyContract) (decimal.Decimal, error)
}
