package blockchain

import (
	"github.com/luzen23141/mouse/pkg/blockchain/_const"
	"github.com/luzen23141/mouse/pkg/blockchain/model"
)

var CurMap = map[string]model.Currency{
	_const.BTC: {
		Name:   "Bitcoin",
		Symbol: "BTC",
		Remark: "比特幣",
		Sort:   1,
		Chain: map[string]model.CurrencyContract{
			_const.BtcChain: {
				Chain:   _const.BtcChain,
				Addr:    "",
				Name:    "Bitcoin",
				Decimal: -8,
				IsGov:   true,
			},
		},
	},
	_const.ETH: {
		Name:   "Ethereum",
		Symbol: "ETH",
		Remark: "以太坊",
		Sort:   2,
		Chain: map[string]model.CurrencyContract{
			_const.EthChain: {
				Chain:   _const.EthChain,
				Addr:    "",
				Name:    "Ethereum",
				Decimal: -18,
				IsGov:   true,
			},
		},
	},
	_const.USDT: {
		Name:   "Tether",
		Symbol: "USDT",
		Remark: "USDT",
		Sort:   3,
		Chain: map[string]model.CurrencyContract{
			"erc20": {
				Chain:   _const.EthChain,
				Addr:    "0xdAC17F958D2ee523a2206206994597C13D831ec7",
				Name:    "Tether",
				Decimal: -6,
				IsGov:   false,
			},
			"trc20": {
				Chain:   _const.TronChain,
				Addr:    "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t",
				Name:    "Tether",
				Decimal: -6,
				IsGov:   false,
			},
		},
	},
	_const.TRX: {
		Name:   "Tron",
		Symbol: "TRX",
		Remark: "Tron",
		Sort:   4,
		Chain: map[string]model.CurrencyContract{
			_const.TronChain: {
				Chain:   _const.TronChain,
				Addr:    "",
				Name:    "Tron",
				Decimal: -6,
				IsGov:   true,
			},
		},
	},
}
