package blockchain

import (
	"github.com/luzen23141/mouse/pkg/blockchain/_const"
	"github.com/luzen23141/mouse/pkg/blockchain/chain"
	"github.com/luzen23141/mouse/pkg/blockchain/model"
	"github.com/shopspring/decimal"
)

var ChainMap = map[string]ChainInterface{
	_const.BtcChain:     chain.NewBtcChain(BtcCfg),
	_const.TronChain:    chain.NewTronChain(),
	_const.EthChain:     chain.NewEvmChain(),
	_const.BscChain:     chain.NewEvmChain(),
	_const.PolygonChain: chain.NewEvmChain(),
	_const.ArbChain:     chain.NewEvmChain(),
	_const.AvaxChain:    chain.NewEvmChain(),
	_const.BaseChain:    chain.NewEvmChain(),
	_const.SuiChain:     chain.NewSuiChain(),
	_const.InjChain:     chain.NewInjChain(), // 差取餘額
	_const.SolChain:     chain.NewSolChain(),
	_const.TonChain:     chain.NewTonChain(), // 差取餘額
}

type ChainInterface interface {
	GenAddr() (string, string, error)
	GenHdAddr() (string, string, error)
	GetAddrBalance(string, model.CurrencyContract) (decimal.Decimal, error)
}
